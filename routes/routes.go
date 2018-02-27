package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"os"
)


func CreateRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/logs", processLogs)

	return r
}

// Binding from JSON
type Log struct {
	DeviceId string `form:"id" json:"id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Log      string `form:"log" json:"log" binding:"required"`
}

func processLogs(c *gin.Context) {
	var json Log
	if err := c.ShouldBindWith(&json, binding.JSON); err == nil {
		if json.Password == "123" {
			err := saveLogs(json)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal error"})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func saveLogs(log Log) error {
	logStream := []byte(log.Log)
	_, err := os.Create("../logs/" + log.DeviceId + ".log")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("../logs/" + log.DeviceId + ".log", logStream, 0644)

	if err != nil {
		return err
	}
	return nil
}

