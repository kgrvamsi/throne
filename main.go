package main

import (
	"fmt"

	"github.com/kgrvamsi/throne/conf"
	"github.com/kgrvamsi/throne/logger"

	"github.com/gin-gonic/gin"
)

var config = conf.GetConf()

func init() {
	log, err := logger.GetLogger(config.Log.LogLevel, config.Log.LogType)
	if err != nil {
		fmt.Println(err)
	}
	if config.Log.LogLevel == "production" {
		gin.SetMode("release")
	}
	log.Info("Server Started @ Port : ", config.Default.Port)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(gin.Recovery())
	r.Run(":" + config.Default.Port) // listen and serve on 0.0.0.0:8080
}
