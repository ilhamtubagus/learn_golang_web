package learn_golang_web

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = r.URL.Query().Get("username")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	// ... can be set multiple times to add multiple cookies
	fmt.Fprintf(w, "Cookie set successfully")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			fmt.Fprintf(w, "No cookie found")
		} else {
			fmt.Printf("Error: %v\n", err)
		}
		return
	}
	fmt.Fprintf(w, "Username: %s", cookie.Value)
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set", SetCookie)
	mux.HandleFunc("/get", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/set?username=john", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()
	var isCookieFound bool = false
	var indexCookie int = -1
	for i, cookie := range cookies {
		if cookie.Name == "username" {
			isCookieFound = true
			indexCookie = i
			break
		}
	}
	assert.True(t, isCookieFound)
	assert.Equal(t, cookies[indexCookie].Value, "john")
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/get", nil)
	cookie := &http.Cookie{Name: "username", Value: "john", Path: "/"}
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, string(body), "Username: john")
}
