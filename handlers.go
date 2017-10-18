package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type torrentDownload struct {
    Magnet string
}

// Handler object type
type Handler struct {
	stream	*Stream
}

// NewHandler returns
func NewHandler() *Handler{
	return &Handler{stream: NewStream()}
}

// Index returns a useless string
func (h *Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

// Torrent starts downloading a torrent with a magnet link
// POST
func (h *Handler) Torrent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    decoder := json.NewDecoder(r.Body)
    var td torrentDownload
    err := decoder.Decode(&td)
    if err != nil {
        panic(err)
    }

    h.stream.NewMagnet(td.Magnet)

    fmt.Fprintf(w, "finished downloading: %s!\n", td.Magnet)
}