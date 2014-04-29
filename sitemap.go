package main

import (
	"lowrey.me/sitemap-go/crawler"
	"fmt"
	"flag"
	"net/url"
)

func main() {
	ignoreFilter := flag.String("ignore", "", "ignore url that match regex")
	u := flag.String("url", "", "target")
	out := flag.String("out", "", "output file")
	flag.Parse()
	ur, _ := url.Parse(*u)
	c, _ := crawler.NewCrawler(*u, *ignoreFilter)
    outfile := *out
    if outfile == "" {
        outfile = fmt.Sprintf("%s.html", ur.Host) 
    }
	c.GetAllLinks().WriteHtml(outfile)
}
