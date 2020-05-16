package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"

	gotube "github.com/asciiu/appa/lib/youtube"
)

func main() {
	flag.Parse()
	log.Println(flag.Args())
	usr, _ := user.Current()

	downloadDir := fmt.Sprintf("%v/Movies/youtubedr", usr.HomeDir)
	log.Println("download to dir=", downloadDir)

	arg := flag.Arg(0)
	if err := gotube.Download(arg, downloadDir); err != nil {
		fmt.Println("err:", err)
	}
}
