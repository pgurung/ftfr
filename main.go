package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	args := os.Args[1:]

	copyFile(args[0], args[1])
}

func copyFile(src, dst string) {
	in, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	createDir(dst)

	out, err := os.Create(dst)

	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		log.Fatal(err)
	}

}

//Parse the second argument and create any directory necessary
func createDir(fp string) {
	var sep string
	if runtime.GOOS == "windows" {
		sep = "\\"
	} else {
		sep = "/"
	}
	ps := strings.Split(fp, sep)

	if len(ps) == 1 {
		return
	}
	path := filepath.Join(ps[:len(ps)-1]...)

	err := os.MkdirAll(path, 0776)
	if err != nil {
		log.Fatal(err)
	}

}
