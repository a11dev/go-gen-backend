package remotes

import (
	"testing"

	"github.com/segmentio/ksuid"
)

type responseMsg struct {
	response  int
	requestID string
	goruntime int
}

func TestInvokeBck(t *testing.T) {
	msgID := ksuid.New().String()
	bckID := 1

	result := InvokeBck(bckID, msgID)

	if result == nil || result["requestID"] == msgID {
		t.Fatalf(`%v, want match for %#q, nil`, msgID, result["requestID"])
	}

}
