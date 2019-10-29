package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// download a git url
func download(p *GitPayload) error {

	url := fmt.Sprintf("%s/archive/master.zip", p.Reposiory.URL)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(p.After)
	if err != nil {
		log.Println("file create err")
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Println("file download err")
		return err
	}

	return nil
}
