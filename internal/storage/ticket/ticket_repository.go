package ticket

import (
	"ticket/internal/helpers"
	"ticket/internal/model"
	"ticket/pkg/logger"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (p *Repository) BuyTicket(userId uint) (*model.Ticket, error) {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	result := tx.Model(&model.User{}).Where("id = ?", userId).Where("balance >= ?", model.TICKET_COST).Update("balance", gorm.Expr("balance - ?", model.TICKET_COST))
	if err := result.Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		return nil, errNotEnoughMoney
	}

	ticket := model.NewTicket(userId)
	if err := tx.Create(&ticket).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return ticket, nil
}

func (p *Repository) EatTicket(userId, ticketId uint) (int, string, error) {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, "", err
	}

	ticket := new(model.Ticket)
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", ticketId).First(&ticket).Error; err != nil {
		tx.Rollback()
		return 0, "", errIdNotFound
	}

	if ticket.UserID != userId {
		tx.Rollback()
		return 0, "", errPermissionDenied
	}

	if ticket.Status == model.StatusEaten {
		tx.Rollback()
		return 0, "", errTicketAlreadyUsed
	}

	newLuck := helpers.GetLuckFromTicket(ticket.Number)

	user := new(model.User)
	if err := tx.Model(&model.User{}).Where("id = ?", userId).First(&user).Error; err != nil {
		tx.Rollback()
		return 0, "", err
	}

	updatedLuck := helpers.UpdateLuck(user.Luck, newLuck)
	if err := tx.Model(&user).Update("luck", updatedLuck).Error; err != nil {
		tx.Rollback()
		return 0, "", err
	}

	if err := tx.Model(&ticket).Update("status", model.StatusEaten).Error; err != nil {
		tx.Rollback()
		return 0, "", err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, "", err
	}

	return newLuck, ticket.Number, nil
}

func (p *Repository) AllByUserId(userId uint) ([]*model.Ticket, error) {
	var tickets []*model.Ticket

	err := p.db.Where("user_id = ?", userId).Find(&tickets).Error
	if err != nil {
		logger.Errorf("Error getting all tickets by user id: %s", err)
		return nil, err
	}

	return tickets, err
}

func (p *Repository) FindById(id uint) (*model.Ticket, error) {
	ticket := new(model.Ticket)

	err := p.db.Where("id = ?", id).First(&ticket).Error
	if err != nil {
		logger.Errorf("Error finding ticket: %s", err)
		return nil, err
	}

	return ticket, err
}

func (p *Repository) Update(ticket *model.Ticket) (*model.Ticket, error) {
	err := p.db.Save(&ticket).Error

	if err != nil {
		logger.Errorf("Error updating  ticket: %s", err)
		return nil, err
	}

	return ticket, err
}

func (p *Repository) Create(ticket *model.Ticket) (*model.Ticket, error) {
	err := p.db.Create(&ticket).Error

	if err != nil {
		logger.Errorf("Error creating ticket: %s", err)
		return nil, err
	}

	return ticket, err
}

func (p *Repository) Delete(id uint) error {
	err := p.db.Delete(&model.Ticket{ID: id}).Error

	return err
}

func (p *Repository) Migrate() error {
	return p.db.AutoMigrate(&model.Ticket{})
}
