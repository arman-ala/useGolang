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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	templatePath := fp.Join("templates", "HomePage.gohtml")
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("parsing template error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Name string
	}{"Arman Ala"}
	err = t.Execute(w, data)
	if err != nil {
		log.Printf("executing template error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func faqPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	templatePath := fp.Join("templates", "FAQ.gohtml")
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userID")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templatePath := fp.Join("templates", "Contact.gohtml")
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	data := struct {
		UserId string
	}{userId}
	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}

	// fmt.Fprint(w, "<h1>Contact Me!</h1><br/><p><a href=\"mailto:arman17gb@gmail.com\">arman17gb@gmail.com</a><br/>User ID =", userId, "<br/>")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact/{userID}", contactHandler)
	r.Get("/faq", faqPage)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, http.StatusText(http.StatusNotFound))
	})
	fmt.Println("Starting the Server on port 3000...")
	http.ListenAndServe(":3000", r)
}
