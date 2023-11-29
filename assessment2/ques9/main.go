package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

var myFile *os.File

func walk(s string, d fs.DirEntry, err error) error {
	file, err := os.OpenFile(myFile.Name(), os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if err != nil {
		return err
	}
	if !d.IsDir() {
		_, err := file.WriteString(s + " \n")
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func main() {

	file, err := os.Create("./file.txt")
	if err != nil {
		panic(err.Error())
	}

	myFile = file

	filepath.WalkDir("..\\learning-module", walk)
}
