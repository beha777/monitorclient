package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"monitorclient/settings"
	"net/http"
	"os/exec"
)

func getParams(context *gin.Context) {
	appSettings, err := settings.ReadSettings()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, appSettings)
	}
	return
}

func sendCommand(context *gin.Context) {
	text := context.Query("text")
	execResult, err := exec.Command("bash", "-c", text).Output()
	execResultString := string(execResult)
	if err != nil {
		if settings.AppSettings.AppParams.Log {
			log.Println("EXEC error!\nCommand: ", text, "\nError: ", err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"response": execResultString,
		})
	}
}
