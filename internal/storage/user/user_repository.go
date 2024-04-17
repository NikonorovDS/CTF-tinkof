package user

import (
	"ticket/internal/model"
	"ticket/pkg/logger"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (p *Repository) FindById(id uint) (*model.User, error) {
	user := new(model.User)

	err := p.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Errorf("Error finding user: %s", err)
		return nil, err
	}

	return user, err
}

func (p *Repository) FindByUsername(username string) (*model.User, error) {
	user := new(model.User)

	err := p.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Errorf("Error finding user: %s", err)
		return nil, err
	}

	return user, err
}

func (p *Repository) Update(user *model.User) (*model.User, error) {
	err := p.db.Save(&user).Error

	if err != nil {
		logger.Errorf("Error updating  user: %s", err)
		return nil, err
	}

	return user, err
}

func (p *Repository) Create(user *model.User) (*model.User, error) {
	err := p.db.Create(&user).Error

	if err != nil {
		logger.Errorf("Error creating user: %s", err)
		return nil, err
	}

	return user, err
}

func (p *Repository) Delete(id uint) error {
	err := p.db.Delete(&model.User{ID: id}).Error

	return err
}

func (p *Repository) Migrate() error {
	return p.db.AutoMigrate(&model.User{})
}
