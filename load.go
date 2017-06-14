package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type WikiData struct {
	Wiki     string `json:"wiki"`
	Api      string `json:"api"`
	Puki     string `json:"puki"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func loadWikiData() (WikiData, error) {
	b, err := ioutil.ReadFile("wiki.json")
	if err != nil {
		log.Printf("Error: %v", err)
		return WikiData{}, err
	}
	
	w := WikiData{}
	err = json.Unmarshal(b, &w)
	
	if err != nil {
		log.Printf("Error: %v", err)
		return w, err
	}
	
	return w, err
}
