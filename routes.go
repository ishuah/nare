package main

import (
	"github.com/julienschmidt/httprouter"
)

// Route object
type Route struct {
	Method string
	Path   string
	Handle httprouter.Handle
}

// Routes is an array
type Routes []Route

func routes() Routes {
	handler := NewHandler()
	var routes = Routes{
		Route{
			"GET",
			"/",
			handler.Index,
		},
		Route{
			"POST",
			"/torrents/magnet",
			handler.Magnet,
		},
		Route{
			"GET",
			"/torrents/:hash",
			handler.Torrent,
		},
		Route{
			"GET",
			"/torrents",
			handler.Torrents,
		},
	}
	return routes
}
