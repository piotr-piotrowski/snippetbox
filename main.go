package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Handler for display home page
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL pag exactly matches "/".
	// If it does not, use the http.NotFound() function to sent a 404 response to the client.
	// Importantly, we then return from the handler.
	// If we do not return the handler would keep executing and also write the message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Display the home page"))
}

// Handler for display a specific snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string
	// and try to convert it to an integer using the strconv.Atoi() function.
	// It it can't be converted to an integer, or the value is less than 1,
	// we return a 404 page not found response
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Handler for display a new snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	if r.Method != http.MethodPost {
		// It it's not, use the w.WriteHeader() method to send 405 status code
		// and the w.Write() method to write a "Method Not Allowed" response body.
		// We then return from the function so that the subsequent code is not executed.
		w.Header().Set("Allow", http.MethodPost)
		// Use the http.Error() function to send a 405 status code and "Method Not Allowed"
		// string as the response body.
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux,
	// then register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Print a log a message to say that the server is starting
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server.
	// We pass in two parameters:
	// - the TCP network address to listen on (in this case ":4000")
	// - the servemux we just created.
	// If http.ListenAndServe() returns an error we use the log.Fatal() function
	// to log  the error message and exit.
	// Note that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
