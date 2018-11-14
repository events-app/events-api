package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Info)
	http.HandleFunc("/api/v1/content/main/", MainContent)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Listening on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Printf("error: listing and serving: %s", err)
		return
	}
}

type Content struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET: https://%s/api/v1/content/main/\n", r.Host)
}

func MainContent(w http.ResponseWriter, r *http.Request) {
	content := Content{
		Name: "main",
		Text: `# The New Event

The New Event is the best event ever.
You should definitelly attend!

+ Register at [The New Event](http://thenewevent.com/).
+ Come
+ Have fun

We are waiting for you!`,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(content); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}
