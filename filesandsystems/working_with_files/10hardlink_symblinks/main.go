// A file will only be deleted from disk once all hard links are removed.
package main


import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a hard link. We will create two file names that point to the same
	// contents, changing the contents of one will change the other.
	// Deleteing/renaming one will not affect the other.
	err := os.Link("original.txt", "original_also.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Create a symlink.
	err = os.Symlink("original.txt", "original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Lstat will return file info/symlink.
	fileInfo, err := os.Lstat("original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Link info: %#v\n", fileInfo)

	// Change ownership of the symlink and not the file it points to.
	err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}
