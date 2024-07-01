package handlers

import (
	"fakhriya/app/controllers"

	"github.com/gin-gonic/gin"
)

func Handlers() {
	r := gin.Default()
	r.POST("/sigup",controllers.Register)
}
func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://192.168.43.45:6600")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	}

	c.Next()
}