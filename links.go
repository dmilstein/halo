package main

import (
	"io/ioutil"
	"regexp"
)

const urlRoot = "http://money.cnn.com"

var linkRE = regexp.MustCompile(
	`href="(/magazines/fortune/fortune_archive.*?htm)"`)

// Extract a list of links to Fortune stories from a string of HTML.
//
// Return a uniquified list of absolute URL's.
//
// An example of the kind of page we'd be pulling links from:
//
// http://money.cnn.com/magazines/fortune/fortune_archive/2004/03/22/toc.html
//
func ExtractLinks(html string) (result []string) {

	matches := linkRE.FindAllStringSubmatch(html, -1)

	foundLinks := make(map[string]bool)
	for _, aMatch := range matches {
		urlPath := aMatch[1]
		if _, present := foundLinks[urlPath]; !present {
			foundLinks[urlPath] = true
			result = append(result, urlRoot+urlPath)
		}
	}
	return
}

func ExtractLinksFromFile(fileName string) []string {

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return ExtractLinks(string(b))
}
