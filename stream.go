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
func (s *Stream) NewMagnet(td *torrentDownload) {
	t, _ := s.client.AddMagnet(td.Magnet)
	td.Hash = t.InfoHash().HexString()
	go s.startTorrent(t)
}

// TorrentProgress returns the specified torrent's progress
func (s *Stream) TorrentProgress(hashString string) (int64, error) {
	hash := metainfo.NewHashFromHex(hashString)
	t, ok := s.client.Torrent(hash)
	if !ok {
		return 0, errors.New("Error retrieving torrent with hash " + hashString)
	}
	return t.BytesCompleted(), nil
}

func (s *Stream) startTorrent(t *torrent.Torrent) {
	<-t.GotInfo()
	t.DownloadAll()
}
