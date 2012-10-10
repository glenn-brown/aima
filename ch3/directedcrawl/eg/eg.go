package main

import (
	"github.com/glenn-brown/directedcrawl"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	verbose := flag.Bool("verbose", false, "Verbose Output")
	from := flag.String("from", "http://www.meetup.com", "Start Page URL")
	to := flag.String("to", "http://www.caltech.edu", "End Page URL")
	flag.Parse()

	w := ioutil.Discard
	if *verbose {
		w = io.Writer(os.Stdout)
	}
	
	fmt.Println(directedcrawl.Crawl (*from, *to, log.New(w, "", 0)))
}
