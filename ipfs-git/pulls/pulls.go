package pulls

import (
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
)

var (
	url  = "https://github.com/FurkanSamaraz/ipfs-git.git"
	path = "/Users/furkan/Desktop/ipfs-repo"
)

// Basic example of how to clone a repository using clone options.
func Pullrepo() {

	// Clone the given repository to the given directory
	Info("git clone " + url)

	git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

}
