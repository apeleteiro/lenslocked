package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func executeTemplate(w http.ResponseWriter, filePath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Printf("error parsing template: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("error executing template: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	executeTemplate(w, filepath.Join("templates", "home.gohtml"))
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	executeTemplate(w, filepath.Join("templates", "contact.gohtml"))
}

func faqHandler(w http.ResponseWriter, _ *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println(err)
		return
	}
}
