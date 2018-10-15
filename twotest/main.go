package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println("Hello World")
}

func GetPathBase(url string) string {
	return path.Base(path.Clean(url))
}
