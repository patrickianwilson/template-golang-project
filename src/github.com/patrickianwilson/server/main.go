package main

import (
	_ "github.com/patrickianwilson/controllers"
	"github.com/patrickianwilson/router"
	"log"
	"net/http"
)

func main() {

	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
