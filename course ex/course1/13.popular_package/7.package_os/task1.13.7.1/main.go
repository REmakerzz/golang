package main

import (
	"fmt"
	"os"
	"path"
)

func main() {

	err := WriteFile("/path/to/file.txt", []byte("Hello, World!"), os.FileMode(0644))
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("File written successfully!")
	}
}

func WriteFile(filepath string, data []byte, perm os.FileMode) error {
	dir := path.Dir(filepath)

	if err := os.MkdirAll(dir, perm); err != nil {
		return err
	}

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return err
	}
	return nil
}
