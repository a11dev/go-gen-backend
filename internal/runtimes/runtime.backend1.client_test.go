package runtimes

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRuntimeBackend1(t *testing.T) {

	in := make(chan BackendMsg, 10)
	go SetupSingleBackendClient(in, 0)

	res := make(chan gin.H)

	var msgID string = "TESTMSG"

	go func() {
		in <- BackendMsg{msgID, res}
	}()
	ginret := <-res
	t.Logf("risultato %+v", ginret)
	if ginret == nil || ginret["requestID"] != msgID {
		t.Fatalf(`%v, want match for %s, %T`, msgID, ginret["requestID"], ginret["requestID"])
	}
	t.Logf("risultato: %+v", ginret)
}
