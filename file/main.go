package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fileName := "test.txt"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	s := "hello \n world!"
	dstFile.WriteString(s + "\n")
	path, err := filepath.Abs(dstFile.Name())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(path)
}
