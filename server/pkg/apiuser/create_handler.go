package apiuser

import (
	"fmt"

	coreuser "github.com/Ekkawin/golang/server/core/user"
	"github.com/Ekkawin/golang/server/datamodel"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) CreateUser(c *gin.Context) {

	var user datamodel.User
	c.BindJSON(&user)
	fmt.Printf("user: %v", user.Email)
	cu := coreuser.NewService(h.db)
	cu.UserCreate(&user)

	c.JSON(200, gin.H{
		"status": "success"})
}
