package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) { // The formHandler function handles requests to the path "/form".
	if err := r.ParseForm(); err != nil { // Parse the form data and log any errors.
		fmt.Fprintf(w, "Error parsing form: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n") // If the path is "/form" and the method is "POST", write "POST request successful" to the response.
	name := r.FormValue("name")                 // Get the value of the "name" field from the form.
	address := r.FormValue("address")           // Get the value of the "address" field from the form.
	fmt.Fprintf(w, "Name = %s\n", name)         // Write the name to the response.
	fmt.Fprintf(w, "Address = %s\n", address)   // Write the address to the response.
}

func helloHandler(w http.ResponseWriter, r *http.Request) { // The helloHandler function handles requests to the path "/hello".
	if r.URL.Path != "/hello" { // If the path is not "/hello", return a 404.
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" { // If the method is not "GET", return a 405.
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!") // If the path is "/hello" and the method is "GET", write "Hello!" to the response.
} // Note that the http.ResponseWriter is an io.Writer, so we can use the fmt.Fprintf function to write to it.

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // Create a file server handler that serves files from the static directory (relative to the current working directory) and returns a 404 if the file is not found.
	http.Handle("/", fileServer)                        // Register the file server as the handler for all requests using the path "/".
	http.HandleFunc("/form", formHandler)               // Register the formHandler function as the handler for the path "/form".")
	http.HandleFunc("/hello", helloHandler)             // Register the helloHandler function as the handler for the path "/hello".")

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { // Start the server on port 8080 and log any errors.
		log.Fatal(err)
	}
}
