package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: 60 * time.Second}
	//get(c, "http://google.com")
	customGet(c, "http://google.com")
	//post(c, "http://google.com")
}

func get(c http.Client, url string) {
	res, err := c.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}

func customGet(c http.Client, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}

func post(c http.Client, url string) {
	jsonString := bytes.NewBuffer([]byte(`{"nome":"Luis", "idade":34}`))
	res, err := c.Post(url, "application/json", jsonString)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}
