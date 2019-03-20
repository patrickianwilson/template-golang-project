package controllers

import (
	"fmt"
	"github.com/patrickianwilson/router"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Excellante!")
}

/*
* Bind this controller to the main routes
 */
func init() {
	router.Routes = append(router.Routes, router.Route{
		"Index",
		"GET",
		"/status",
		Status,
	})
}
