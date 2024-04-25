package main

import (
	"fmt"
	"net/http"
)

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, http.StatusText(http.StatusNotFound))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<h1>Home Page</h1>")
}

func faqPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<h1>faq</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Me!</h1><br/><p><a href=\"mailto:arman17gb@gmail.com\">arman17gb@gmail.com</a>")
}

type Router struct {}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
		case "/":
			homepage(w, r)
		case "/contact":
			contactHandler(w, r)
		case "/faq":
			faqPage(w, r)
        default:
			notFoundPage(w, r)
	}
}

func main() {
	var router Router;
	fmt.Println("Starting the Server on port 3000...")
	http.ListenAndServe(":3000", router)
}
