package youtube_test

import (
	"fmt"
	"os"
	"os/user"
	"testing"

	gotube "github.com/asciiu/appa/lib/youtube"
	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	usr, _ := user.Current()
	downloadDir := fmt.Sprintf("%v/Movies/youtubedr", usr.HomeDir)
	url := "https://www.youtube.com/watch?v=LWE79K2Ii-s"

	filename, err := gotube.Download(url, downloadDir)
	assert.Nil(t, err, fmt.Sprintf("download failed with"))

	_, err = os.Stat(filename)
	assert.Nil(t, err, fmt.Sprintf("filename not found"))
}
