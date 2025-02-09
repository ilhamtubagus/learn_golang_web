package learn_golang_web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(wr http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type") // key are case-insensitive

	fmt.Println(contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Custom-Header", "Custom Value")
	fmt.Fprintf(w, "Custom Response Header!")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()

	assert.Equal(t, response.Header.Get("X-Custom-Header"), "Custom Value")
}
