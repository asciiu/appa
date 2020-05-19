package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"path/filepath"
)

func main() {
	flag.Parse()
	log.Println(flag.Args())

	usr, _ := user.Current()
	downloadDir := fmt.Sprintf("%v/Movies/youtubedr", usr.HomeDir)
	log.Println("download to dir=", downloadDir)

	// convert mp4
	path := filepath.Join(downloadDir, "LWE79K2Ii-s.mp4")

	fmt.Println(path)
}
