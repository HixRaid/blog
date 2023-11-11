package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hixraid/blog/internal/config"
	"github.com/hixraid/blog/internal/data/repository"
	"github.com/hixraid/blog/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	cfgFile, err := config.LoadConfig("config")
	if err != nil {
		logrus.Fatalf("can't load config: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		logrus.Fatalf("can't parse config: %v", err)
	}

	_, err = repository.NewMySql(cfg.DB)
	if err != nil {
		logrus.Fatalf("error database connection: %v", err)
	}

	srv := server.New(cfg.Server.Addr, nil)

	go func() {
		if err := srv.Run(); err != nil {
			logrus.Fatalf("error occurred while running server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(); err != nil {
		logrus.Errorf("error occurred while shutting down server: %v", err)
	}
}
