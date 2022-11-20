package main

import (
	"final-project-backend/config"
	"final-project-backend/pkg/database/postgres"
	"final-project-backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.InitConfig()
	db := postgres.New(cfg)

	r := gin.Default()
	router.RouterSetUp(r, db)

	r.Run(cfg.Server.Url)
}
