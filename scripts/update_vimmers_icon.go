package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	filename = "./vimmers/vimmers.json"
)

func download(u string) (string, error) {
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("content-type")
	ext := ""
	if strings.Index(ct, "/jpeg") != -1 {
		ext = "jpg"
	} else if strings.Index(ct, "/png") != -1 {
		ext = "png"
	} else if strings.Index(ct, "/gif") != -1 {
		ext = "gif"
	}
	sum := md5.Sum([]byte(u))
	filename := filepath.Join("assets/images/vimmers", fmt.Sprintf("%x", sum[:])+"."+ext)
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return filename, err
}

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
				if h, err := download(icon); err == nil {
					vimmer["twitter_icon"] = "/" + filepath.ToSlash(h)
				} else {
					vimmer["twitter_icon"] = icon
				}
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
