package main

import (
	"errors"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
)

// Stream object
type Stream struct {
	client *torrent.Client
}

// NewStream returns a Stream instance
func NewStream() *Stream {
	c, _ := torrent.NewClient(nil)
	return &Stream{client: c}
}

// NewMagnet starts downloading from a magnet link
func (s *Stream) NewMagnet(d *download) (Torrent, error) {
	t, err := s.client.AddMagnet(d.Source)

	if err != nil {
		return Torrent{}, err
	}

	tt := Torrent{Name: t.Name(),
		Hash: t.InfoHash().HexString()}

	go s.startTorrent(t)
	return tt, nil
}

// Torrents returns all the torrents added to a client
func (s *Stream) Torrents() []Torrent {
	var ts []Torrent
	for _, t := range s.client.Torrents() {
		tt := Torrent{Name: t.Name(),
			Hash:           t.InfoHash().HexString(),
			Length:         t.Length(),
			BytesCompleted: t.BytesCompleted(),
			Files:          t.Info().Files}
		ts = append(ts, tt)
	}
	return ts
}

// Torrent returns the specified torrent's progress
func (s *Stream) Torrent(hashString string) (Torrent, error) {
	hash := metainfo.NewHashFromHex(hashString)
	t, ok := s.client.Torrent(hash)
	if !ok {
		return Torrent{}, errors.New("Error retrieving torrent with hash " + hashString)
	}

	tt := Torrent{Name: t.Name(),
		Hash:           t.InfoHash().HexString(),
		Length:         t.Length(),
		BytesCompleted: t.BytesCompleted(),
		Files:          t.Info().Files}
	return tt, nil
}

func (s *Stream) startTorrent(t *torrent.Torrent) {
	<-t.GotInfo()
	t.DownloadAll()
}
