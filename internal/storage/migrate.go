package storage

import "ticket/pkg/logger"

func MigrateTables(s Store) {
	if err := s.Tickets().Migrate(); err != nil {
		logger.Errorf("failed to migrate tickets: %v", err)
	}

	if err := s.Users().Migrate(); err != nil {
		logger.Errorf("failed to migrate users: %v", err)
	}
}
