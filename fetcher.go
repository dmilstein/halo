package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"halo"
)

const (
	pageCacheDir = "/Users/danmil/Projects/Halo/fortune"
	rootPage     = "testdata/fortune_toc_2004_03_22.html"
)

func DownloadPage(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Couldn't fetch %v: %v", url, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	destFile := pageCacheDir + "/" +
		base64.StdEncoding.EncodeToString([]byte(url))
	err = ioutil.WriteFile(destFile, body, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	links := ExtractLinksFromFile(rootPage)
	for _, url := range links {
		fmt.Printf("Attempting to download %v\n", url)
		DownloadPage(url)
	}
}
