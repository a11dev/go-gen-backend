package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetTask(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := gin.New()

	// Define the route similar to its definition in the routes file
	r.GET("/task/:id", GetTask)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/task/707389", strings.NewReader(""))
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
	wants := `{"Key":978897047,"Number":707389,"Status":"Deleted","Typecategory":"AA","Refid":"REFID_000004"}`
	if err != nil || strings.Index(string(p), wants) < 0 {
		t.Fatalf(`%v, want match for %#q, nil`, wants, string(p))
	}
}
