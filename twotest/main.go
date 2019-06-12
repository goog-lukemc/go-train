package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	fmt.Println("Hello World")
}

func GetPathBase(httpurl string) (*string, error) {
	// Requirement:
	// - Must pass all test cases
	// - Must be a valid request URI
	// Hint see golang package net/url ParseRequestURL() and the path package

	u, err := url.ParseRequestURI(httpurl)
	if err != nil {
		return nil, err
	}
	parameter := path.Base(u.Path)
	return &parameter, nil
}
