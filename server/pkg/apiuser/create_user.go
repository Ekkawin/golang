package apiuser

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	userName := c.Param("userName")
	password := c.Param("password")
	fmt.Printf("userName: %v; password: v%", userName, password)
	status := c.Writer.Status()
	fmt.Println("status", status)
	c.JSON(200, gin.H{
		"status":   "successssssssss",
		"userName": userName,
		"password": password,
	})

}
