package utils

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ishuah/batian/config"
)

// Download downloads files via http
func Download(uri string) (string, error) {
	c, err := config.GetConfig()
	segments := strings.Split(uri, "/")
	filename := c.DownloadDir + segments[len(segments)-1]

	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	res, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}
