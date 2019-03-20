package router

import (
	"github.com/patrickianwilson/logging"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	STATIC_PATH = "/static/"
	STATIC_DIR  = "/assets/"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range Routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(logging.Logger(route.HandlerFunc, route.Name))
	}

	pwd, _ := os.Getwd()

	log.Printf("Serving Static Resources from path: %s", pwd+STATIC_DIR)

	router.
		PathPrefix(STATIC_PATH).
		Handler(http.StripPrefix(STATIC_PATH, http.FileServer(http.Dir(pwd+STATIC_DIR))))

	return router
}

var Routes = make([]Route, 0)
