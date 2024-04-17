package storage

import (
	"log"
	"os"

	"ticket/internal/storage/ticket"
	"ticket/internal/storage/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	db *gorm.DB

	users   UserRepository
	tickets TicketRepository
}

func DBConn(dsn string) (*gorm.DB, error) {
	// Ignore "RecordNotFound" Error
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true,
		},
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return db, err
	}

	return db, err
}

func New(db *gorm.DB) *Database {
	return &Database{
		db:      db,
		users:   user.NewRepository(db),
		tickets: ticket.NewRepository(db),
	}
}

func (db *Database) Users() UserRepository {
	return db.users
}

func (db *Database) Tickets() TicketRepository {
	return db.tickets
}

func (db *Database) Ping() error {
	sqlDB, err := db.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}
