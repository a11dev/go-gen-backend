package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetFromTaskList(t *testing.T) {

	result, err := getFromTaskList(978897047)
	var wants int64
	wants = 978897047
	if err != nil || result.Key != 978897047 {
		t.Fatalf(`%v, want match for %#q, nil`, wants, result)
	}
}

func TestGetTask(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := gin.New()

	// Define the route similar to its definition in the routes file
	r.GET("/task/:id", GetTask)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/task/978897047", strings.NewReader(""))
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
	wants := `{"Key":978897047,"Body":{"Key":978897047,"Number":707389,"Refid":"REFID_000004","Status":"Deleted","Tpecategory":"AA"}}`
	if err != nil || strings.Index(string(p), wants) < 0 {
		t.Fatalf(`%v, want match for %#q, nil`, wants, string(p))
	}
}

func TestGetTaskList(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := gin.New()

	// Define the route similar to its definition in the routes file
	r.GET("/tasks", GetTasks)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/tasks", strings.NewReader(""))
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

	var f []map[string]interface{}
	json.Unmarshal(p, &f)
	var wants int64 = 978897386
	if err != nil || len(f) != 14 || int64(f[0]["Key"].(float64)) != wants {
		t.Fatalf(`%v, not match requirements`, f)
	}
}
