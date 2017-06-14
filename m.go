package main

import (
	"flag"
	"log"
	
	"github.com/sadbox/mediawiki"
)


func main() {
	var title string
	var changeTitle bool
	var uploadImgOnly bool
	var createPageOnly bool
	
	flag.BoolVar(&changeTitle, "changetitle", false, "是否轉換標題，例如 角色/ABC 轉成 ABC 這樣")
	flag.BoolVar(&uploadImgOnly, "ci", false, "是否僅上傳圖片")
	flag.BoolVar(&createPageOnly, "cp", false, "是否僅建立條目")
	flag.StringVar(&title, "t", "", "wiki 標題")
	flag.StringVar(&title, "title", "", "wiki 標題")

	flag.Parse()
	
	if title == "" {
		log.Print("Error: Title is empty")
		return
	}
	
	wiki, _ := loadWikiData()
	
	client, err := mediawiki.New(wiki.Api, "")
	err = client.Login(wiki.Account, wiki.Password)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Login %s as %s", wiki.Wiki, wiki.Account)
	
	defer client.Logout()
	
	success := make(chan bool)
	action := 0
	
	if uploadImgOnly || (!createPageOnly && !uploadImgOnly){
		go searchImg(success, wiki.Puki, title, client)
		action++
	}
	if createPageOnly || (!createPageOnly && !uploadImgOnly){
		go editPage(success, wiki.Puki, title, changeTitle, client) 
		action++
	}
	
	for i:= 0; i < action; i++{
		<- success
	}
}
