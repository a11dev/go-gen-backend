package handlers

import (
	"github.com/a11dev/go-gen-backend/internal/helpers"
	"github.com/a11dev/go-gen-backend/internal/runtimes"
	"github.com/gin-gonic/gin"
)

func InvokeBackend(c *gin.Context) {
	client := helpers.GetClientInputChannel(c)

	res := make(chan gin.H)

	go func() {
		client <- runtimes.BackendMsg{c.Param("id"), res}
	}()
	ginret := <-res

	c.JSON(200, ginret)

}
