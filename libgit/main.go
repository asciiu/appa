package main

import (
	"fmt"

	"gopkg.in/libgit2/git2go.v27"
)

func main() {
	// init a new git repo under the web directory
	repo, err := git.InitRepository("web", false)
	//repo, err := git.Clone("git://github.com/gopheracademy/gopheracademy-web.git", "web", &git.CloneOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(repo)
}
