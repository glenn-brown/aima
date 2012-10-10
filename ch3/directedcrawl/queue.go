package directedcrawl

import (
	"container/list"
	"regexp"
)

// A hostQ stores per-host URL information.
//
type hostQ struct {
	hostname string
	paths    []*path
}

func newHostQ(hostname string) *hostQ {
	return &hostQ{hostname, []*path{}}
}

func (h *hostQ) pushBack(p *path) {
	h.paths = append(h.paths, p)
}

func (h *hostQ) pop() (p *path) {
	if len(h.paths) == 0 {
		return nil
	}
	p = h.paths[0]
	h.paths = h.paths[1:]
	return p
}

func (h *hostQ) empty() bool {
	return len(h.paths) == 0
}

// A pathQ stores paths, grouped by hostname, in FIFO order per host.
// Paths are removed from host queues in round-robin fashion, so
// crawling does not pound single host, or get stuck in a honeypot.
//
type pathQ struct {
	hostQMap  map[string]*hostQ
	hostQList list.List
}

var hostRe, _ = regexp.Compile("http://([^/]*)")

func newPathQ() *pathQ {
	return &pathQ{map[string]*hostQ{}, list.List{}}
}

func (a *pathQ) add(p *path) {
	sub := hostRe.FindStringSubmatch(p.url)
	if len(sub) < 2 { // No hostname
		return
	}
	hq := a.hostQMap[sub[1]]
	if hq == nil {
		hq = newHostQ(string(sub[1]))
		a.hostQList.PushBack(hq)
	}
	hq.pushBack(p)
}

func (a *pathQ) rm() (p *path) {
	front := a.hostQList.Front()
	if front == nil {
		return nil
	}
	hq := a.hostQList.Remove(front).(*hostQ)
	p = hq.pop()
	if !hq.empty() {
		a.hostQList.PushBack(hq)
	}
	return p
}

func (a *pathQ) empty() bool {
	return a.hostQList.Front() == nil
}