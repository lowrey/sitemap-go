# sitemap.go
Build
-----
    go get github.com/PuerkitoBio/goquery
    go install

Usage
------
* Crawl a site and output a list of pages visited
    go run sitemap.go -url=http://httpbin.org/
* Limit the pages visited by a regex
    go run sitemap.go -url=http://httpbin.org/ -ignore=".*\/\d+ gzip"
* Output links to a file
    go run sitemap.go -url=http://httpbin.org/ -out="example.html"


Design
------
### Organization
Input is handled by the functionality in sitemap.go. The flag "url" is required.  The ignore flag tests multiple regular expression, delimted by a space, to determine if a URL is valid to visit.  URL fetching and crawling logic are contianed in the crawler subpackage.

Output of each URL visited is written to stdout as it is processed.  It will also output an HTML file containing each link when it finishes.  The output file can be set using the "out" flag.

### Tests
Automated unit testing has been created for the crawler package. All tests can be using the following *Go Tool* command in the root of the git repo:  
	go test ./crawler

### Externl libraries 
* [goquery](http://github.com/PuerkitoBio/goquery/)

License
------
MIT
