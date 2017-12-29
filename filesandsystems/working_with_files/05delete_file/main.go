package main

import (
	"log"
	"os"
)

func main() {
	if err := os.Remove("test.txt"); err != nil {
		log.Fatal(err)
	}
}
