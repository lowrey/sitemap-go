package crawler

import (
	"net/url"
	"testing"
)

const (
	testSite    string = "http://httpbin.org/"
	testSubSite string = "http://httpbin.org/user-agent"
	testFilter  string = "user"
)

func TestUrlSetAddsOnlyOnce(t *testing.T) {
	urlSet := make(UrlSet)
	u, _ := url.Parse(testSite)
	urlSet.Add(u)
	u2, _ := url.Parse(testSite)
	urlSet.Add(u2)
	if len(urlSet) > 1 {
		t.Log(urlSet)
		t.Fatal("url set added more than one for the same url")
	}
}

func TestGetAllLinks(t *testing.T) {
	c, _ := NewCrawler(testSite, testFilter)
	u2 := c.GetAllLinks()
	if len(u2) == 0 {
		t.Error("error getting links")
	}
}

func TestValidUrl(t *testing.T) {
	c, _ := NewCrawler(testSubSite, testFilter)
	u, _ := url.Parse(testSubSite)
	if c.isValidUrl(u) {
		t.Error("filter didnt work")
	}
}
