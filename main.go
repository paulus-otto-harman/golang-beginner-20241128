package main

import (
	"20241128/config"
	"20241128/controller"
	"20241128/database"
	"20241128/middleware"
	"20241128/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()

	// Middleware
	r.Use(middleware.Logger())
	//r.Use(middleware.BasicAuth())
	cfg, _ := config.LoadConfig()
	db, _ := database.ConnectDB(cfg)

	router.APIRouter(r, controller.NewController(db))

	r.Run(viper.GetString("SERVER_PORT"))
}
