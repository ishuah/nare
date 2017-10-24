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

type torrentProgress struct {
	BytesCompleted int64
	Length         int64
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

// Magnet starts downloading a torrent when given a magnet link
// POST
func (h *Handler) Magnet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var t torrentDownload
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	h.stream.NewMagnet(&t)
	h.jsonResponse(w, http.StatusCreated, t)
}

// Progress returns a torrent's length and bytesCompleted when given a torrent hash
func (h *Handler) Progress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	length, bytesCompleted, err := h.stream.TorrentProgress(ps.ByName("hash"))
	if err != nil {
		panic(err)
	}

	progress := torrentProgress{BytesCompleted: bytesCompleted, Length: length}

	h.jsonResponse(w, http.StatusOK, progress)
}

func (h *Handler) jsonResponse(w http.ResponseWriter, status int, object interface{}) {
	response, err := json.Marshal(object)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
