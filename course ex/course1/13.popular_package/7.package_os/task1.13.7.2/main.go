package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func main() {
	filepath := "course1/13.popular_package/7.package_os/task1.13.7.2/file.txt"

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error open file: ", err)
		return
	}

	defer file.Close()

	err = WriteFile(strings.NewReader("Hello, World!"), file)
	if err != nil {
		fmt.Println("Error write file: ", err)
	} else {
		fmt.Println("file wirte successfully")
	}

}

func WriteFile(data io.Reader, fd io.Writer) error {
	file, ok := fd.(*os.File)
	if !ok {
		return fmt.Errorf("invalid file descriptor")
	}

	filepath := file.Name()
	dir := path.Dir(filepath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if _, err := io.Copy(file, data); err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}
	return nil
}
