package learn_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}

// TestServeMux ServeMux will combine all http handler function into single handler (commonly called router)
// Parameters:
//
//	t - The testing object used to manage
func TestServeMux(t *testing.T) {
	var mux *http.ServeMux = http.NewServeMux()

	// adding "/" after route name will make route with /* redirected to this handler
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "images")
		if err != nil {
			return
		}
	})

	mux.HandleFunc("/images/thumbnails", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "thumbnails")
		if err != nil {
			return
		}
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux, // mux implements Handler interface which has method ServeHTTP
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}
