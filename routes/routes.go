package routes

import (
	"github.com/gin-gonic/gin"
	"monitorclient/settings"
)

func Init() {
	router := gin.Default()
	router.GET("/getParams", getParams)
	router.GET("/sendCommand", sendCommand)
	router.Run(":" + settings.AppSettings.AppParams.PortRun)
}
