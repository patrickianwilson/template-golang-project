package controllers

import (
	"fmt"
	"github.com/patrickianwilson/router"
	"net/http"
)

type IndexModel struct {
	Name string
}

func IndexController(w http.ResponseWriter, r *http.Request) {
	model := &IndexModel{
		Name: "Foo Bar",
	}
	fmt.Fprintln(w, templateManager.RenderTemplate("index", model))
}

/*
* Bind this controller to the main routes
 */
func init() {
	router.Routes = append(router.Routes, router.Route{
		"Index",
		"GET",
		"/",
		IndexController,
	})
}
