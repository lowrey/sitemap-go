package crawler

import (
	"net/url"
	"regexp"
	"strings"
	//"container/list"
	"fmt"
	"log"
	"os"
)

type UrlSet map[string]bool

//TODO trim trailing slashes?
func (urlSet UrlSet) Add(u *url.URL) bool {
	_, found := urlSet[u.String()]
	urlSet[u.String()] = true
	return !found
}

func (urlSet UrlSet) Get() []string {
	keys := make([]string, len(urlSet))
	i := 0
	for k := range urlSet {
		keys[i] = k
		i++
	}
	//sort.Strings(keys)
	return keys
}

type Crawler struct {
	startUrl     *url.URL
	filters      []string
	useQueries   bool
	useFragments bool
	useLog       bool
}

func NewCrawler(startUrl string, filter string) (*Crawler, error) {
	c := new(Crawler)
	u, err := url.Parse(startUrl)
	if err != nil {
		return nil, err
	}
	c.startUrl = u

	c.filters = strings.Fields(filter)
	c.useQueries = false
	c.useFragments = false
	c.useLog = true
	return c, nil
}

func getUrls(loc string, jobs chan int, urls chan *url.URL) {
	jobs <- 1
	grabber, err := NewGrabber(loc)
	if err == nil {
		links := grabber.GetAbsLinks()
		for _, link := range links {
			urls <- link
		}
	}
	jobs <- -1
}

func (crawler *Crawler) GetAllLinks() UrlSet {
	urlSet := make(UrlSet)
	jobCount := 0
	urls := make(chan *url.URL)
	jobs := make(chan int)
	go getUrls(crawler.startUrl.String(), jobs, urls)
	for {
		select {
		case c, more := <-jobs:
			if more {
				//log.Println(jobCount)
				jobCount += c
				if jobCount < 1 {
					//log.Println("done")
					close(jobs)
				}
			} else {
				return urlSet
			}
		case u := <-urls:
			if crawler.isValidUrl(u) {
				prev := len(urlSet)
				urlSet.Add(u)
				//havent seen the url before
				if len(urlSet) > prev {
					if crawler.useLog {
						log.Println("Visit: ", u)
					}
					go getUrls(u.String(), jobs, urls)
				}
			}
		}
	}
}

func (c *Crawler) isValidUrl(u *url.URL) bool {
	for _, f := range c.filters {
		if m, _ := regexp.MatchString(f, u.String()); m && f != "" {
			return false
		}
	}
	if len(u.Query()) != 0 && !c.useQueries {
		return false
	}
	if u.Fragment != "" && !c.useFragments {
		return false
	}
	if u.Host != c.startUrl.Host {
		return false
	}
	//TODO remove
	if strings.Contains(u.Path, u.Host) {
		return false
	}
	return true
}

func (urlSet UrlSet) Write(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, link := range urlSet.Get() {
		f.WriteString(link + "\n")
	}
	return nil
}

func (urlSet UrlSet) WriteHtml(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, link := range urlSet.Get() {
		a := fmt.Sprintf("<a href='%s'>%s</a><br/>\n", link, link)
		f.WriteString(a)
	}
	return nil
}
