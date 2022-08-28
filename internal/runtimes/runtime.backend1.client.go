package runtimes

import (
	"log"
	"time"

	"github.com/a11dev/go-gen-backend/internal/remotes"
	"github.com/gin-gonic/gin"
)

type BackendMsg struct {
	Message string
	Retc    chan gin.H
}

// must be initialized as go routine
// it consumes in chanell messages
// each message contains the request information plus the result channel
// each request must inizialize a BackendMsg message and post it to the backend client
// the response is posted on the return channel

func SetupSingleBackendClient(in <-chan BackendMsg, idx int) {
	log.Printf("%15s | %5d | %20s | %20s ", "BackEnd Started #", idx, "", "")
	// first time is opened a health service is called
	for rin := range in {
		time.Sleep(3 * time.Second)
		result := remotes.InvokeBck(idx, rin.Message)
		rin.Retc <- result
	}
}
