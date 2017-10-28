package handlers

import (
	"encoding/json"
	"text/template"
	"net/http"

	"github.com/ishuah/batian/stream"
	"github.com/julienschmidt/httprouter"
)

// Handler object type
type Handler struct {
	stream *stream.Stream
}

// NewHandler returns
func NewHandler() *Handler {
	return &Handler{stream: stream.NewStream()}
}

// Index returns the SPA
func (h *Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("frontend/index.html")
	t.Execute(w, nil)
}

// Torrents returns all the torrents loaded in the Client
func (h *Handler) Torrents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	torrents := h.stream.Torrents()
	h.jsonResponse(w, http.StatusOK, torrents)
}

// Torrent returns a torrent's length and bytesCompleted when given a torrent hash
func (h *Handler) Torrent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t, err := h.stream.Torrent(ps.ByName("hash"))
	if err != nil {
		h.jsonResponse(w, http.StatusNotFound, err)
	}

	h.jsonResponse(w, http.StatusOK, t)
}

// Magnet starts downloading a torrent when given a magnet link
// POST
func (h *Handler) Magnet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var d stream.Download
	err := decoder.Decode(&d)
	if err != nil {
		h.jsonResponse(w, http.StatusBadRequest, err)
	}

	t, err := h.stream.NewMagnet(&d)

	if err != nil {
		h.jsonResponse(w, http.StatusBadRequest, err)
	}
	h.jsonResponse(w, http.StatusCreated, t)
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
