package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func printFileInfo() {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %#v\n\n", fileInfo.Sys())
}


func main() {
	_, err := os.Stat("test.txt")
	if os.IsNotExist(err) {
		if _, err = os.Create("test.txt"); err != nil {
			log.Fatal(err)
		}
	}
	printFileInfo()

	// Change permissions.
	if err = os.Chmod("test.txt", 0777); err != nil {
		log.Println(err)
	}

	// Change ownership.
	if err = os.Chown("test.txt", os.Getuid(), os.Getgid()); err != nil {
		log.Println(err)
	}

	// Change timestamps.
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	if err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime); err != nil {
		log.Println(err)
	}

	printFileInfo()
}
