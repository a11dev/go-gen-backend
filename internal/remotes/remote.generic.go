package remotes

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

func InvokeBck(backendIDX int, requestID string) gin.H {
	return gin.H{
		"response":  rand.Intn(100),
		"requestID": requestID,
		"goruntime": backendIDX,
	}
}
