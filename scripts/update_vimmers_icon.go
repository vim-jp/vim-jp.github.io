package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"os"
)

const (
	filename = "./vimmers/vimmers.json"
)

func main() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	var vimmers []map[string]interface{}
	err = json.NewDecoder(f).Decode(&vimmers)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	for _, vimmer := range vimmers {
		if twitter, ok := vimmer["twitter"]; ok {
			doc, err := goquery.NewDocument("https://twitter.com/" + twitter.(string))
			if err != nil {
				log.Println(err)
				continue
			}
			if icon, ok := doc.Find("img.ProfileAvatar-image").First().Attr("src"); ok {
				fmt.Println(icon)
				vimmer["twitter_icon"] = icon
			}
		}
	}
	b, err := json.MarshalIndent(vimmers, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
