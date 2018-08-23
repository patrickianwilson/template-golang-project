package main


import (
	"net/http"

	"github.com/gorilla/mux"
	"log"
	"os"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

const (
	STATIC_PATH = "/static/"
	STATIC_DIR = "/assets/"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(Logger(route.HandlerFunc, route.Name))

		log.Printf("Registering Handler: %s on path %s with method %s", route.Name, route.Pattern, route.Method)

	}


	pwd,_ := os.Getwd()

	log.Printf("Serving Static Resources from path: %s", pwd + STATIC_DIR)

	router.
		PathPrefix(STATIC_PATH).
		Handler(http.StripPrefix(STATIC_PATH, http.FileServer(http.Dir(pwd + STATIC_DIR))))


	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Status",
		"GET",
		"/status",
		Status,
	},
}