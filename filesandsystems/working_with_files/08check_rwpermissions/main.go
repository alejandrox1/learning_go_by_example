package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}

		if os.IsPermission(err) {
			log.Println("Error: Write permission denied")
		}
	}
	file.Close()

	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied")
		}
	}
	defer file.Close()
}
