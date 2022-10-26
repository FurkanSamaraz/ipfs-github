package main

import (
	"fmt"
	"log"
	"main/pulls"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

var (
	path = "/Users/furkan/Desktop/ipfs-repo"
)

func main() {

	pulls.Pullrepo()

	f, _ := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)

	sh := shell.NewShell("localhost:5001")

	cid, _ := sh.AddDir(f.Name())

	fmt.Printf("https://ipfs.io/ipfs/%s", cid)

	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}

}
