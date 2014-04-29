package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	//"log"
)

type Error struct {
	S string
}

func (e *Error) Error() string {
	return e.S
}

type Grabber struct {
	baseUrl *url.URL
	doc     *goquery.Document
}

func NewGrabber(baseUrl string) (*Grabber, error) {
	c := new(Grabber)
	var e error
	c.baseUrl, e = url.Parse(baseUrl)
	if e != nil {
		return nil, e
	}
	//log.Printf("Crawling from url [%s]", c.baseUrl.String())

	c.doc, e = goquery.NewDocument(c.baseUrl.String())
	if e != nil {
		return nil, e
	}
	return c, nil
}

func (c *Grabber) getLinks() []string {
	links := []string{}
	c.doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		links = append(links, href)
	})
	return links
}

//TODO does not work for www links
func (c *Grabber) GetAbsLinks() []*url.URL {
	urls := []*url.URL{}
	for _, link := range c.getLinks() {
		url, err := c.transformRelativeUrl(link)
		if err == nil {
			//log.Println(url.String())
			urls = append(urls, url)
		}
	}
	return urls
}

func (c *Grabber) transformRelativeUrl(weburl string) (*url.URL, error) {
	relUrl, err := url.Parse(weburl)
	if err != nil {
		return nil, err
	}
	if relUrl.Host != c.baseUrl.Host && relUrl.Host != "" {
		return nil, &Error{"URL not compatible with base URL"}
	}
	return c.baseUrl.ResolveReference(relUrl), nil
}
