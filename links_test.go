package halo

import (
	"reflect"
	"testing"
)

func TestExtractLinksMakesAbsoluteLinks(t *testing.T) {
	urlRoot := "http://money.cnn.com"
	aLink := "/magazines/fortune/fortune_archive/2004/03/22/365076/index.htm"

	htmlString := "blah href=\"" + aLink + "\" blah"

	expectedLinks := []string{urlRoot + aLink}
	foundLinks := ExtractLinks(htmlString)

	if !reflect.DeepEqual(foundLinks, expectedLinks) {
		t.Errorf("Found %#v instead of %#v", foundLinks, expectedLinks)
	}
}

func TestExtractLinksReturnsUniquifiedList(t *testing.T) {
	urlRoot := "http://money.cnn.com"
	aLink := "/magazines/fortune/fortune_archive/2004/03/22/365076/index.htm"

	htmlString := "first href=\"" + aLink + "\" blah " +
		"second href=\"" + aLink + "\" blah"

	expectedLinks := []string{urlRoot + aLink}
	foundLinks := ExtractLinks(htmlString)

	if !reflect.DeepEqual(foundLinks, expectedLinks) {
		t.Errorf("Found %#v instead of %#v", foundLinks, expectedLinks)
	}
}
