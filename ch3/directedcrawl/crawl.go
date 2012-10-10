// Package crawl is a solution to AIMA 3.19
package directedcrawl

import (
	"github.com/PuerkitoBio/purell"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// The path type records crawl results as paths to the crawl starting point.
type path struct {
	url  string
	rest *path // nil at start point
}

// Convert URLs to a standard form for comparison.
func normalizeURL(url string) string {
	if n, err := purell.NormalizeURLString(url, purell.FlagsUsuallySafe); err != nil {
		return n
	}
	return url
}

// Crawl finds a path from one URL to another, returning the path from one to the other.
func Crawl(from, to string, log *log.Logger) []string {
	from = normalizeURL(from)
	to = normalizeURL(to)
	log.Printf("Crawling from %v to %v.\n", from, to)
	
	// Create terminal paths at source and destination
	fromEndpoint := &path{from, nil}
	toEndpoint := &path{to, nil}

	// Create a map to record which URLs we have seen, and map them to their paths.
	prefix := map[string]*path{from: fromEndpoint}
	suffix := map[string]*path{to: toEndpoint}

	// Create forward and backward crawlers, with buffering for acks.
	reqForward := make(chan *path)
	ackForward := make(chan []*path, 1)
	go crawl(4, reqForward, ackForward, enumerateLinksForward, log)

	reqBack := make(chan *path)
	ackBack := make(chan []*path, 1)
	go crawl(4, reqBack, ackBack, enumerateLinksBackward, log)

	// Send crawlers their first request
	reqForward <- fromEndpoint
	reqBack <- toEndpoint

	// Listen for acks, which carry discovered paths, and record the discoved paths.
	// Stop when both a prefix and suffix are know for any URL.
	var p *path
outer:
	for {
		select {
		case pp := <-ackForward:
			for _, p = range pp {
				if prefix[p.url] == nil {
					prefix[p.url] = p
					if _, ok := suffix[p.url]; ok {
						break outer
					}
					log.Printf("---> %s\n", p.url)
					reqForward <- p
				}
			}
		case pp := <-ackBack:
			for _, p = range pp {
				if suffix[p.url] == nil {
					suffix[p.url] = p
					if _, ok := prefix[p.url]; ok {
						break outer
					}
					log.Printf("<--- %s\n", p.url)
					reqBack <- p
				}
			}
		}
	}
	// Terminate the crawler goroutines
	close(reqForward)
	close(reqBack)
	// Report results.
	return pathForPrefixSuffix(prefix[p.url], suffix[p.url])
}

// Combine prefix and suffix into a single path.
func pathForPrefixSuffix(prefix *path, suffix *path) []string {
	// Add prefixes to result.
	ret := []string{}
	for i := prefix; i != nil; i = i.rest {
		ret = append(ret, i.url)
	}
	// Reverse the backwards prefix order
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	// Add suffixes to result.
	for i := suffix.rest; i != nil; i = i.rest {
		ret = append(ret, i.url)
	}
	return ret
}

// An enumLinksFn finds all links one hop past path p, and passes them to channel ack.
type enumLinksFn func(p *path, ack chan<- []*path, l *log.Logger)

// Function crawl records paths passed to it on channel req, and
// crawls them N at a time using enumLinks.  The resulting next-hop
// paths are passed to channel ACK.
func crawl(n int, req <-chan *path, ack chan<- []*path, enumerateLinks enumLinksFn, log *log.Logger) {
	pending := newPathQ()             // Paths queued for crawling
	children := make(chan []*path, n) // Buffer acks so children can be GC'd.
	for {
		select {
		case p, ok := <-req:
			// Record requests and process them
			// asynchronously if not too busy.
			if !ok {
				return // Input was closed.
			}
			pending.add(p)
			if n > 0 && !pending.empty() {
				go enumerateLinks(pending.rm(), children, log)
				n--
			}
		case cc := <-children:
			// When our request complete, start more if
			// possible, and pass up results.
			if next := pending.rm(); next != nil {
				go enumerateLinks(next, children, log)
			} else {
				n++
			}
			ack <- cc
		}
	}
}

// A precompiled regular expression for matching URLs.
var aRegexp = regexp.MustCompile("<[aA][^>]*>")
var urlRegexp = regexp.MustCompile("http://[^ ;)<>#'%\"\t\n]*")

// pageURLs() returns the URLs in anchors in a web page.
func pageURLs(url string, log *log.Logger) []string {
	// Download the page
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	// Extract the anchors
	// Extract the URLs
	all, err := ioutil.ReadAll(resp.Body);
	if err != nil {
		log.Println("ReadAll failed: ", err)
		return nil
	}
	aa := aRegexp.FindAll(all, -1)
	uu := []string{}
	for _, a := range aa {
		url := normalizeURL(string(urlRegexp.Find(a)))
		if len(url) > 0 {
			uu = append(uu, url)
		}
	}
	return uu
}

// Function crawlForward reads the page at the head of the path, and
// reports all paths linked to it.
func enumerateLinksForward(from *path, ack chan<- []*path, log *log.Logger) {
	urls := pageURLs(from.url, log)
	paths := []*path{}
	for _, url := range urls {
		paths = append(paths, &path{string(url), from})
	}
	ack <- paths
}

// Function crawlBackward reports URLs that refer to the page.
func enumerateLinksBackward(to *path, ack chan<- []*path, log *log.Logger) {
	// Find pages mentioned in the destination.  These might have
	// links to the destination.  It would be much better to use a
	// search engine to find the links, but Google et al. have
	// started blocking or charging for automatic queries.
	dest := to.url
	urls := pageURLs(dest, log)
	paths := []*path{}
	// Load each of those pages and see if they actually contain the forward link
	for _, url := range urls {
		urls2 := pageURLs(string(url), log)
		// Report the match if the to.url is present.
		for _, url2 := range urls2 {
			// Match if the cleaned-up URLS match.
			if dest == string(url2) {
				paths = append(paths, &path{url, to})
				break
			}
		}
	}
	ack <- paths
}
