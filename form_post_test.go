package learn_golang_web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Error parsing form", http.StatusBadRequest)
		return
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	// or we can invoke request.PostFormValue() method to get values for multiple keys without calling ParseForm(

	fmt.Fprintf(writer, "Hello, %s %s!", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader(`first_name=John&last_name=Doe`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, string(body), "Hello, John Doe!")
}
