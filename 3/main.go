package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello World!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Me!</h1><br/><p><a href=\"mailto:arman17gb@gmail.com\">arman17gb@gmail.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404 Not Found</h1>")
	fmt.Fprint(w, "<p>The page you are looking for could not be found.</p>")
	fmt.Fprintf(w, "Host: %s<br/>Path: %s<br/>RawPath: %s", r.URL.Host, r.URL.Path, r.URL.RawPath)
	fmt.Fprint(w, "<p><a href=\"/\">Back to Home</a></p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
		case "/":
			homeHandler(w, r)
		case "/contact":
			contactHandler(w, r)
		default:
			notFoundHandler(w, r)
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	http.HandleFunc("/contact", pathHandler) 
	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", nil)
}
