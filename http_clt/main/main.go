package main

import (
	"io"
	"log"
	"net/http"
	//"strings"
)

func main() {
	requrl := "http://127.0.0.1:8080/abc.file"

	req, err := http.NewRequest("GET", requrl, nil)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	h := NewHashCheck(resp.Body)

	//io.Discard
	buf := make([]byte, 1024)
	io.CopyBuffer(io.Discard, h, buf)
}
