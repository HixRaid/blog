package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hixraid/blog/internal/config"
	"github.com/hixraid/blog/internal/handler"
	"github.com/hixraid/blog/internal/server"
	"github.com/hixraid/blog/pkg/data/repository"
	"github.com/hixraid/blog/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	if !debugFlag() {
		gin.SetMode(gin.ReleaseMode)
	}

	cfgFile, err := config.LoadConfig("config")
	if err != nil {
		logrus.Fatalf("can't load config: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		logrus.Fatalf("can't parse config: %v", err)
	}

	db, err := repository.NewMySql(cfg.DB)
	if err != nil {
		logrus.Fatalf("error database connection: %v", err)
	}

	repos := repository.New(db)
	service := service.New(repos)
	handler := handler.New(service)
	router := handler.InitRouter()

	srv := server.New(cfg.Server.Addr, router)

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

func debugFlag() bool {
	var isEnable = flag.Bool("debug", false, "-debug, uses false as default")
	flag.Parse()
	return *isEnable
}
