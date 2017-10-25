package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/julienschmidt/httprouter"
)

type download struct {
	Source string
}

type Torrent struct {
	Name           string
	Hash           string
	Length         int64
	BytesCompleted int64
	Files          []metainfo.FileInfo
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
	var d download
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
