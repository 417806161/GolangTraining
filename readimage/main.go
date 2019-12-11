package main

import (
	"io"
	"log"
	"os"
)

func main() {
	imgPath := "./readimage/image/6TBf9r1c.jpg"
	file, err := os.Open(imgPath)
	if err != nil {
		log.Fatalln(err)
	}
}
