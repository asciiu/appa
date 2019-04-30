package main

import (
	"gopkg.in/libgit2/git2go.v27"
)

func main() {
	_, err := git.Clone("git://github.com/gopheracademy/gopheracademy-web.git", "web", &git.CloneOptions{})
	if err != nil {
		panic(err)
	}
}
