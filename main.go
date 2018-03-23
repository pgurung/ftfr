package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	args := os.Args[1:]

	write(read(args), args)

}

//Read the file passed in as the first argument
func read(s []string) []byte {
	src, err := ioutil.ReadFile(s[0])

	if err != nil {
		log.Fatal(err)
	}

	return src
}

//Writes the file to the destination (pased in as second argument) and create any directory necessary
func write(b []byte, ss []string) {
	createDir(ss)
	err := ioutil.WriteFile(ss[1], b, 0666)

	if err != nil {
		log.Fatal(err)
	}
}

//Parse the second argument and create any directory necessary
func createDir(ss []string) {
	var sep string
	if runtime.GOOS == "windows" {
		sep = "\\"
	} else {
		sep = "/"
	}
	ps := strings.Split(ss[1], sep)
	path := filepath.Join(ps[:len(ps)-1]...)

	err := os.MkdirAll(path, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
