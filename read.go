package main

import (
	"log"
	"net/http"
	"strings"
	"regexp"
	
	"github.com/sadbox/mediawiki"
	"github.com/shiyou0130011/p2mfmt"
	"github.com/PuerkitoBio/goquery"
)



func editPage(success chan bool, puki string, title string, changeTitle bool, client *mediawiki.MWApi) {
	if title == "" {
		log.Print("Error: Title is blank")
		success <- false
		return
	}

	pageTitle := title
	if strings.Contains(title, "/") {
		category := title[0:strings.Index(title, "/")]
		_, err := client.Read(category)

		if err == nil && category != "成句" && changeTitle {
			// 如果沒有 error ，表示此為頁面的 category
			pageTitle = title[strings.Index(title, "/")+1:]
		}
	}
	log.Printf("Read Page %s", title)

	text, categories := p2mfmt.Convert(readPuki(puki, title))
	for _, category := range categories {
		log.Printf("Add Category “%s”", category)
		text += "\n" + `[[category:` + category + `]]`
	}

	text = regexp.MustCompile(`(?i)(&amp;|&)color\((|black),(|black)\){[^}]*};`).ReplaceAllStringFunc(
		text,
		func(source string) string {
			s := source[strings.Index(source, "{")+1 : len(source)-2]

			return `{{censored|` + s + `}}`
		},
	)
	text = regexp.MustCompile(`(?i)(&amp;|&)color\((|black),(|black)\){[^;]*};`).ReplaceAllStringFunc(
		text,
		func(source string) string {
			s := source[strings.Index(source, "{")+1 : len(source)-2]

			return `{{censored|` + s + `}}`
		},
	)

	for _, line := range strings.Split(text, "\n") {
		trimedLine := strings.Trim(line, "\t\r ")
		if len(trimedLine) > 2 && trimedLine[0:1] == "=" && trimedLine[len(trimedLine)-1:] == "=" && trimedLine[1:2] != "=" {
			l := strings.Trim(trimedLine[1:len(trimedLine)-1], " \t")
			if l == pageTitle {
				text = strings.Replace(text, line, "", 1)
			} else {
				text = strings.Replace(text, line, "{{h0|"+strings.Replace(l, "=", "&#61;", -1)+"}}", 1)
			}

			break
		}
	}

	err := client.Edit(
		map[string]string{
			"title":   pageTitle,
			"summary": "機器人自動新增條目",
			"bot":     "true",
			//			"createonly": "true",
			"text": text,
		},
	)

	if err != nil {
		log.Printf("Error: %v", err)
		log.Printf("Create Page %s fail", pageTitle)
		success <- false
	}

	if title != pageTitle {
		err = client.Edit(
			map[string]string{
				"title": title,
				"bot":   "true",
				"text":  "#redirect [[" + pageTitle + "]]",
			},
		)
		if err != nil {
			log.Printf("Error: %v", err)
			log.Printf("Create Redirect Page %s fail", title)
			success <- false
		}
	}

	log.Printf("Create Page %s successfully", pageTitle)
	success <- true
}

func searchImg(success chan bool, puki, title string, client *mediawiki.MWApi) {
	doc, err := goquery.NewDocument(puki + "?" + title)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	imgs := []string{}

	doc.Find("img").Each(func(index int, n *goquery.Selection) {
		imgSrc, _ := n.Attr("src")

		imgarr := strings.Split(imgSrc, "/")
		if len(imgarr) > 1 && strings.Contains(imgarr[len(imgarr)-1], "img") {
			imgs = append(imgs, "https://wiki.komica.org/pix/"+imgarr[len(imgarr)-1])
		}

	})

	doc.Find("a").Each(func(index int, n *goquery.Selection) {
		aSrc, _ := n.Attr("href")

		aarr := strings.Split(aSrc, "/")
		if len(aarr) > 1 && strings.Contains(aarr[len(aarr)-1], "img") {
			imgs = append(imgs, "https://wiki.komica.org/pix/"+aarr[len(aarr)-1])
		}

	})

	log.Printf(`讀取圖片 %v`, imgs)

	a := make(chan bool)
	for _, v := range imgs {
		go (func(a chan bool, url string, client *mediawiki.MWApi) {
			resp, err := http.Get(url)
			if err != nil {
				a <- false
			}
			defer resp.Body.Close()

			us := strings.Split(url, "/")
			fn := us[len(us)-1]

			comment := "機器人自動上傳"
			if title != "" {
				comment = "[[" + title + "]]" + "條目用圖。\n\n" + comment
			}

			err = client.Upload(
				fn,
				resp.Body,
				map[string]string{
					"comment": comment,
				},
			)
			if err != nil {
				log.Print(`Error: Upload Fail "` + fn + `".`)
				a <- false
				return
			}
			log.Print("Upload file \"" + fn + `".`)

			a <- true
		})(a, v, client)
	}

	for range imgs {
		<-a
	}

	success <- true
}


func readPuki(puki string, title string) string {
	log.Printf(`read "%s"`, puki+"?cmd=edit&page="+title)
	doc, err := goquery.NewDocument(puki + "?cmd=edit&page=" + title)
	if err != nil {
		log.Printf("Error: %v", err)
		return ""
	}

	result := ""

	doc.Find("textarea").Each(func(index int, n *goquery.Selection) {
		if name, exists := n.Attr("name"); exists && name == "msg" {
			result = n.Text()
		}
	})

	result = strings.NewReplacer(
		"#contents", "",
	).Replace(result)
	return result
}
