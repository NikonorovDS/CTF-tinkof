package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"ticket/internal/router"
	"ticket/internal/storage"
	"ticket/pkg/logger"
	"ticket/internal/config"
)

func main() {
	if err := os.Chdir(filepath.Dir(appFilePath())); err != nil {
		logger.Fatalf("os.Chdir failed error: %v", err)
	}

	cfg := config.GetConfig()

	db, err := storage.DBConn(cfg.DB.DSN)
	if err != nil {
		logger.Fatalf("failed to init db: %s", err)
	}

	s := storage.New(db)

	storage.MigrateTables(s)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Web.Port),
		Handler: router.New(s, cfg.Web.SecretKey),
	}

	msg := fmt.Sprintf("Ticket is up and running on ':%d'", cfg.Web.Port)
	logger.Infof(msg)

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("failed to start server: %v", err)
	}
}

func appFilePath() string {
	path, err := os.Executable()
	if err != nil {
		return os.Args[0]
	}
	return path
}
