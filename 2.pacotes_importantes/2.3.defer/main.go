package main

import (
	"net/http"
	"io"
)

func main() {
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
}