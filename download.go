package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-getter"
)

// download a git url
func download(p *GitPayload) error {
	path := fmt.Sprintf("data/%s/%s", p.Reposiory.Name, p.After)
	url := p.Reposiory.GitURL

	log.Println(path)

	err := getter.Get(path, url)
	if err != nil {
		log.Printf("error downloading: %e", err)
		return err
	}
	log.Println("download successful")
	return nil
}
