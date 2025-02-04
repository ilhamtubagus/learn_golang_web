package learn_golang_web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "Hello, World!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

func TestSayHello(t *testing.T) {
	// in request url we're passing name in query parameter
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=John", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, bodyString, "Hello, John!")
}

func TestSayHelloWithoutQueryParam(t *testing.T) {
	// in request url we're not passing any name in query parameter
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, bodyString, "Hello, World!")
}

func MultipleQueryParam(w http.ResponseWriter, r *http.Request) {
	// Get query parameter for specified key
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(w, "Name: %s, Age: %s", name, age)
}

func TestMultipleQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/multiple?name=John&age=30", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, bodyString, "Name: John, Age: 30")
}

func MultipleValueQueryParam(w http.ResponseWriter, r *http.Request) {
	// Query parameter value for each key are stored as string slice
	var query url.Values = r.URL.Query()
	// We will not use Get method to get all values for a key
	var names []string = query["name"]

	fmt.Fprintln(w, strings.Join(names, ", "))
}

func TestMultipleValueQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/multiple?name=John&name=Doe", nil)
	recorder := httptest.NewRecorder()

	MultipleValueQueryParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, bodyString, "John, Doe\n")
}
