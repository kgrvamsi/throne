package main

import (
	"fmt"
	"github.com/kgrvamsi/throne/conf"
	"github.com/kgrvamsi/throne/logger"

	"github.com/gin-gonic/gin"
)

var config = conf.GetConf()

func init() {
	log, err := logger.GetLogger(config.Log.LogLevel, "elasticsearch")
	if err != nil {
		fmt.Println(err)
	}
	if config.Log.LogLevel == "production" {
		gin.SetMode("release")
	}
	log.Info("Server Started")
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + config.Default.Port) // listen and serve on 0.0.0.0:8080
}
