/*
	Truncate a file to 100 bytes. If file contains less than a 100 bytes then 
	rest of the space will be filled wih null bytes.
*/
package main

import (
	"log"
	"os"
)

func main() {
	if err := os.Truncate("test.txt", 100); err != nil {
		log.Fatal(err)
	}
}
