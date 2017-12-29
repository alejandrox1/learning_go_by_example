package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "test.txt"
	newPath := "test2.txt"

	if err := os.Rename(originalPath, newPath); err != nil {
		err = os.Rename(newPath, originalPath)
		if err != nil {
			log.Fatal("Error renaming back to original name: ", err)
			return
		}
		if err != nil {
			log.Fatal("Error renaming: ", err)
		}
	}
}
