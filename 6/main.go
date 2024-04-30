package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "os"

	chi "github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	t, err := template.ParseFiles("HomePage.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name string
	}{"Arman Ala"}
	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func faqPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<h1>faq</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userID")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Me!</h1><br/><p><a href=\"mailto:arman17gb@gmail.com\">arman17gb@gmail.com</a><br/>User ID =", userId, "<br/>")
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
	fmt.Println(r)
	fmt.Println("Starting the Server on port 3000...")
	http.ListenAndServe(":3000", r)
}
