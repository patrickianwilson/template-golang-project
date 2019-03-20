package router

import (
	"github.com/gorilla/mux"
	"github.com/patrickianwilson/logging"
	"net/http"
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

	return router
}

var Routes = make([]Route, 0)
