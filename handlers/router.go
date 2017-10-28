package handlers

import (
	"github.com/julienschmidt/httprouter"
)

// NewRouter returns a httprouter
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	for _, route := range routes() {
		router.Handle(route.Method, route.Path, route.Handle)
	}
	return router
}
