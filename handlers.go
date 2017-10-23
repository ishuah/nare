package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type torrentDownload struct {
	Magnet string
	Hash   string
}

// Handler object type
type Handler struct {
	stream *Stream
}

// NewHandler returns
func NewHandler() *Handler {
	return &Handler{stream: NewStream()}
}

// Index returns a useless string
func (h *Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Magnet starts downloading a torrent with a magnet link
// POST
func (h *Handler) Magnet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var t torrentDownload
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	h.stream.NewMagnet(&t)
	response, err := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
