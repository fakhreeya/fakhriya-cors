package controllers

import (
	"fakhriya/app/structs"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context)  {
	var RegisterPeremenaya structs.RegisterStruct
	c.ShouldBindJSON(&RegisterPeremenaya)
	if RegisterPeremenaya.Name=="" || RegisterPeremenaya.Surname=="" || RegisterPeremenaya.Phone==0{
		c.JSON(200,"SUCCESS")
	}
}