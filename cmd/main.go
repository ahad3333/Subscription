package main

import (
	"add/api"
	"add/config"
	"add/pkg/cronjob"
	"add/storage/postgres"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	storage, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer storage.CloseDB()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	api.NewApi(r, storage)
	cronjobService := cronjob.NewCronJobService(storage, log.Logger{})
	go cronjob.CronJob(cronjobService)

	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
