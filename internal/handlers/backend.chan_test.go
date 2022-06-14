package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/a11dev/go-gen-backend/internal/middleware"
	"github.com/a11dev/go-gen-backend/internal/runtimes"
	"github.com/gin-gonic/gin"
)

func TestGetRuntimeBackend(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := gin.New()

	in := make(chan runtimes.BackendMsg, 10)

	go runtimes.SetupSingleBackendClient(in, 0)

	chans := r.Group("/chans")
	chans.Use(middleware.Backend1ChannelsMiddleware(in))
	{
		chans.GET("/routine/:id", InvokeBackend)
	}
	// rr := r.Routes()
	// fmt.Printf("rr: %v\n", rr)
	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/chans/routine/707389", strings.NewReader(""))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len("")))

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 200
	if w.Code != http.StatusOK {
		t.Fail()
	}

	// Test that the page title is "Successful Login"
	// You can carry out a lot more detailed tests using libraries that can
	// parse and process HTML pages
	p, err := ioutil.ReadAll(w.Body)
	gh := gin.H{}

	json.Unmarshal([]byte(string(p)), &gh)
	wants := "737389"
	if err != nil || gh["requestID"] != "707389" {
		t.Fatalf(`%v, want match for %#q, nil`, wants, gh["requestID"])
	}
}
