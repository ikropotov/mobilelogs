package routes

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)


func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	// Run the other tests
	retCode := m.Run()

	os.Exit(retCode)
}