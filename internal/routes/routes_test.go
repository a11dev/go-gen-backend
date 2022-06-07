// handlers.user_test.go

package routes

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/a11dev/go-gen-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

// Test that a POST request to login returns a success message for
// an unauthenticated user
func TestHealthCheck(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := gin.New()

	// Define the route similar to its definition in the routes file
	r.GET("/healthcheck", handlers.HealthCheck)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/healthcheck", strings.NewReader(""))
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

	if err != nil || strings.Index(string(p), "{\"status\":\"SYSTEM OK\"}") < 0 {
		t.Fail()
	}
}
