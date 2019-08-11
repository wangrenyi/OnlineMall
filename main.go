package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"onlinemall/config"
	"onlinemall/logging"
	"onlinemall/router"
	"time"
)

func main() {
	serverConfig := config.ServerConfig

	gin.SetMode(serverConfig.Mod)
	engine := gin.Default()
	engine.Use(logging.WinLoggerHandler())

	//配置跨域
	engine.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	router.Router(engine)

	logging.Info("server start..")
	engine.Run(serverConfig.Address)
}
