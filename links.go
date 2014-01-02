package halo

import (
	"regexp"
)

const urlRoot = "http://money.cnn.com"

var linkRE = regexp.MustCompile(
	`href="(/magazines/fortune/fortune_archive.*?htm)"`)

// Extract a list of links to Fortune stories from (for now at least)
// a Fortune ToC page, ala:
// http://money.cnn.com/magazines/fortune/fortune_archive/2004/03/22/toc.html
//
// Return a uniquified list of absolute URL's.
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
