package crawler

import (
	"testing"
)

func TestLinksAreNotEmpty(t *testing.T) {
	c, err := NewGrabber(testSite)
	if err != nil {
		t.Log(err.Error())
		t.Fatal("error creating crawler")
	}
	links := c.getLinks()
	if len(links) < 1 {
		t.Log(err.Error())
		t.Error("No links found on page")
	}
}

func TestRelativeUrlMadeAbsolute(t *testing.T) {
	c, _ := NewGrabber(testSite)
	url, _ := c.transformRelativeUrl("./test/")
	if url.String() != testSite+"test/" {
		t.Error("Unexpected value calculating relative link")
	}
}

func TestErrorWhenUsingAbsoluteUrl(t *testing.T) {
	c, _ := NewGrabber(testSite)
	url, err := c.transformRelativeUrl("http://www.metal-archives.com/")
	if err == nil {
		t.Log(url.String())
		t.Error("No error parsing absolute url")
	}
}

func TestGetAbsLinks(t *testing.T) {
	c, _ := NewGrabber(testSite)
	for _, url := range c.GetAbsLinks() {
		if url.Host != c.baseUrl.Host {
			t.Log(url.String())
			t.Error("url did not contain the proper host")
		}
	}
}
