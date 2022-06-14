package helpers

import (
	"github.com/a11dev/go-gen-backend/internal/runtimes"
	"github.com/gin-gonic/gin"
)

func GetClientInputChannel(c *gin.Context) chan runtimes.BackendMsg {
	incClient := c.MustGet("retChan").(chan runtimes.BackendMsg)

	return incClient
}
