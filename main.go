package main

import (
    "html/template"
    "log"
    "net/http"
)

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("internal/templates/home.html")
        if err != nil {
            log.Printf("Error parsing template: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        data := map[string]string{
            "Title":       "My Portfolio",
            "Description": "Welcome to my portfolio website! Explore my projects and get to know me better.",
        }
        tmpl.Execute(w, data)
    })

    log.Println("Server running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
