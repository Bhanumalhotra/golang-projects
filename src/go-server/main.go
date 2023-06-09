// Declare the main package
package main

// Import necessary packages
import (
	"fmt"
	"log"
	"net/http"
)

// Define a handler function for a form
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	// Print a success message
	fmt.Fprintf(w, "POST request successful")
	// Retrieve form values by name
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// Define a handler function for the "hello" path
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the path matches "/hello"
	if r.URL.Path != "/hello" {
		// If not, return a 404 error
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// Check if the HTTP method is GET
	if r.Method != "GET" {
		// If not, return an error message
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	// If everything is good, print "hello!"
	fmt.Fprintf(w, "hello!")
}

// Define the main function
func main() {
	// Create a file server to serve static files
	fileServer := http.FileServer(http.Dir("./static"))
	// Register the file server to the root URL path
	http.Handle("/", fileServer)
	// Register the formHandler to the "/form" URL path
	http.HandleFunc("/form", formHandler)
	// Register the helloHandler to the "/hello" URL path
	http.HandleFunc("/hello", helloHandler)

	// Start the server
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// If there's an error, log it
		log.Fatal(err)
	}
}
