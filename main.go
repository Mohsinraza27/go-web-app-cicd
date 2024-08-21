package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func skillsPage(w http.ResponseWriter, r *http.Request) {
	// Render the skills html page from static folder
	http.ServeFile(w, r, "static/skills.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page from static folder
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page from static folder
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	// Serve static files (CSS, JS, images, etc.) directly from the "static" folder
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve specific routes
	http.HandleFunc("/", homePage)          // Serve home page at the root URL
	http.HandleFunc("/skills", skillsPage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	// Start the server
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
