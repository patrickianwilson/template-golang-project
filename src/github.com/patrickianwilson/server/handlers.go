package main

import (
	"fmt"
	"net/http"

)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, "<title>Loading...</title>")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, "<p>Loading Application...</p>")
	fmt.Fprintln(w, "<script src=\"/static/main.js\">")
	fmt.Fprintln(w, "</script>")
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")

}

func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Excellante!")
}

