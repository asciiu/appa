package youtube

import (
	"fmt"
	"log"
	"path/filepath"

	. "github.com/kkdai/youtube"
)

//### Dependencies
// go - v1.14
// brew install ffmpeg (OSX)
// apt install ffmpeg (Linux)

func Download(url, downloadDir string) (string, error) {
	log.Println("downloading fiel...", url)

	youtube := NewYoutube(true)

	if err := youtube.DecodeURL(url); err != nil {
		log.Println("err:", err)
		return "", err
	}

	filename := fmt.Sprintf("%s.mp4", youtube.VideoID)
	path := filepath.Join(downloadDir, filename)

	if err := youtube.StartDownload(path); err != nil {
		log.Println("err:", err)
		return "", err
	}

	return path, nil
}
