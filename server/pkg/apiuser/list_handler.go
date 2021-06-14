package apiuser

import (
	coreuser "github.com/Ekkawin/golang/server/core/user"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) ListHandler(c *gin.Context) {

	cu := coreuser.NewService(h.db)
	users := cu.ListUser()

	c.JSON(200, gin.H{
		"status": "success",
		"data":   users})
}
