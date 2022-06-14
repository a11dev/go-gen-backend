package middleware

import (
	"github.com/a11dev/go-gen-backend/internal/runtimes"
	"github.com/gin-gonic/gin"
)

// add the backend client input channel to the context
func Backend1ChannelsMiddleware(inc chan runtimes.BackendMsg) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("retChan", inc)
		// c.Set("sxpClient", sxpclient)
	}
}
