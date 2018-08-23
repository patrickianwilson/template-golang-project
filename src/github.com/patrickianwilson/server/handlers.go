package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentations"},
		Todo{Name: "Host meetup"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

