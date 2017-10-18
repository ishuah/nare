package main

import (
	"github.com/anacrolix/torrent"
)

// Stream object
type Stream struct {
	client 		*torrent.Client
	torrents	map[string]*torrent.Torrent
}

// NewStream returns a Stream instance
func NewStream() *Stream {
	c, _ := torrent.NewClient(nil)
	return &Stream{client: c, torrents: map[string]*torrent.Torrent{}}
}

// NewMagnet starts downloading from a magnet link
func (s *Stream) NewMagnet(magnet string) {
	t, _ := s.client.AddMagnet(magnet)
	hash := t.InfoHash().HexString()
	s.torrents[hash] = t
	<-t.GotInfo()
	t.DownloadAll()
}

// TorrentProgress returns the specified torrent's progress
func (s *Stream) TorrentProgress(hash string) int64 {
	t := s.torrents[hash]
	return t.BytesCompleted()
}