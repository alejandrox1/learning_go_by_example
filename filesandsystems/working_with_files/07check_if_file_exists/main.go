package main

import (
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	if _, err = os.Stat("test.txt"); os.IsNotExist(err) {
		log.Fatal("File does not exist")
	}
	log.Println("File does exist")
}
