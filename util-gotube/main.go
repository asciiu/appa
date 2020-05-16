package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"path/filepath"

	. "github.com/kkdai/youtube"
)

func main() {
	flag.Parse()
	log.Println(flag.Args())
	usr, _ := user.Current()

	currentDir := fmt.Sprintf("%v/Movies/youtubedr", usr.HomeDir)
	log.Println("download to dir=", currentDir)

	youtube := NewYoutube(true)

	arg := flag.Arg(0)
	if err := youtube.DecodeURL(arg); err != nil {
		fmt.Println("err:", err)
	}

	filename := fmt.Sprintf("%s.mp4", youtube.VideoID)
	fmt.Println(filename)

	if err := youtube.StartDownload(filepath.Join(currentDir, filename)); err != nil {
		fmt.Println("err:", err)
	}
}
