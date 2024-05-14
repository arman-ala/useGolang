package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	fp "path/filepath"

	// "os"

	chi "github.com/go-chi/chi/v5"
)

func templateExecuter(w http.ResponseWriter, filePath string) (err error) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	t, err := template.ParseFiles(filePath)
	if err != nil {
		log.Printf("parsing template error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("executing template error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := fp.Join("templates", "HomePage.gohtml")
	templateExecuter(w, templatePath)
}

func faqPage(w http.ResponseWriter, r *http.Request) {
	templatePath := fp.Join("templates", "FAQ.gohtml")
	templateExecuter(w, templatePath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := fp.Join("templates", "Contact.gohtml")
	templateExecuter(w, templatePath)

	// fmt.Fprint(w, "<h1>Contact Me!</h1><br/><p><a href=\"mailto:arman17gb@gmail.com\">arman17gb@gmail.com</a><br/>User ID =", userId, "<br/>")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqPage)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, http.StatusText(http.StatusNotFound))
	})
	fmt.Println("Starting the Server on port 3000...")
	http.ListenAndServe(":3000", r)
}