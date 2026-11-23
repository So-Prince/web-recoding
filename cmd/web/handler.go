package main

import (
	"html/template"
	"net/http"
	"strings"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "http.StatusMethodNotAllowed", http.StatusMethodNotAllowed)
	}

	if r.URL.Path != "/" {
		http.Error(w, "http.StatusInternalServerError", http.StatusInternalServerError)
	}
	ts, err := template.ParseFiles("./ui/html/main.html")
	if err != nil {
		http.Error(w, "http.StatusBadRequest", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)

	err = ts.ExecuteTemplate(w, "main.html", nil)
	if err != nil {
		http.Error(w, "http.StatusBadRequest", http.StatusBadRequest)
	}
}

func asciiArt(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "http.StatusMethodNotAllowed", http.StatusMethodNotAllowed)
	}

	if r.URL.Path != "/asciiart" {
		http.Error(w, "http.StatusInternalServerError", http.StatusInternalServerError)
	}

	ts, err := template.ParseFiles("./ui/html/main.html")
	if err != nil {
		http.Error(w, "http.StatusBadRequest", http.StatusBadRequest)
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, errorMessage(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	text := strings.TrimSpace(r.FormValue("text"))
	banner := strings.TrimSpace(r.FormValue("banner"))

	if text == "" || banner == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	// switch banner {
	// case "standard", "thinkertoy", "shadow":
	// default:
	// 	http.Error(w, "Error banner", http.StatusNotFound)
	// 	return
	// }

	result, _ := AsciiartFont(text, banner)
	// log.Print(result)

	ts, err = template.ParseFiles("./ui/html/main.html")
	if err != nil {
		http.Error(w, errorMessage(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	ts.ExecuteTemplate(w, "main.html", result)
}
