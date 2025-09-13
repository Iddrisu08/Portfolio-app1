package main

import (
    "encoding/json"
    "html/template"
    "log"
    "net/http"
    "time"
)

type ContactMessage struct {
    Name    string `json:"name"`
    Email   string `json:"email"`
    Subject string `json:"subject"`
    Message string `json:"message"`
    Date    time.Time
}

func main() {
    // Serve static files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Home page handler
    http.HandleFunc("/", homeHandler)
    
    // Contact form handler
    http.HandleFunc("/contact", contactHandler)

    log.Println("Server running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Only serve home page for root path
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    tmpl, err := template.ParseFiles("internal/templates/home.html")
    if err != nil {
        log.Printf("Error parsing template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    data := map[string]string{
        "Title":       "Iddrisu Abdul Razak Iddrisu - Portfolio",
        "Description": "Portfolio of Iddrisu Abdul Razak Iddrisu - DevOps Engineer & Software Developer",
    }
    
    if err := tmpl.Execute(w, data); err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    // Set CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    var message ContactMessage
    if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
        log.Printf("Error decoding JSON: %v", err)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error":   "Invalid JSON data",
        })
        return
    }

    // Validate required fields
    if message.Name == "" || message.Email == "" || message.Message == "" {
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error":   "Name, email, and message are required",
        })
        return
    }

    // Add timestamp
    message.Date = time.Now()

    // Log the message (in a real application, you'd save to database or send email)
    log.Printf("Contact form submission:")
    log.Printf("Name: %s", message.Name)
    log.Printf("Email: %s", message.Email)
    log.Printf("Subject: %s", message.Subject)
    log.Printf("Message: %s", message.Message)
    log.Printf("Date: %s", message.Date.Format(time.RFC3339))


    // Respond with success
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
        "message": "Thank you for your message! I'll get back to you soon.",
    })
}
