package main

import (
	"fmt"
	"os"
    "log"
    "practice/links"
)

var maxDepth int = 3
func crawl(url string, depth int) []string {
    fmt.Println(depth, url)
    list, err := links.Extract(url)
    if err != nil {
        log.Print(err)
    }
    //depth += 1
    return list
}

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	
	// Add command-line arguments to worklist.
    go func() { worklist <- os.Args[1:] }()

	// Create 4 crawler goroutines to fetch each unseen link.
    for i := 0; i < 4; i++ {
        var depth int = 0
        
        go func() {
            for link := range unseenLinks {
                if depth >= maxDepth {
                    fmt.Println("done")
                    return
                }
                foundLinks := crawl(link,depth)
                go func() { worklist <- foundLinks }()
                depth++
            }
        }()
    }

    // The main goroutine de-duplicates worklist items
    // and sends the unseen ones to the crawlers.
    seen := make(map[string]bool)
    for list := range worklist {
        for _, link := range list {
            if !seen[link] {
                seen[link] = true
                unseenLinks <- link
            }
        }
    }








	/*
	// Crawl the web concurrently.
    seen := make(map[string]bool)
    for list := range worklist {
        for _, link := range list {
            if !seen[link] {
                seen[link] = true
                go func(link string) {
                    worklist <- crawl(link)
                }(link)
            }
        }
	}
	*/
}



/*
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
    seen := make(map[string]bool)
    for len(worklist) > 0 {
        items := worklist
        worklist = nil
        for _, item := range items {
            if !seen[item] {
                seen[item] = true
                worklist = append(worklist, f(item)...)
            }
        }
    }
}
*/

/*
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 4)

func crawl(url string) []string {
    fmt.Println(url)
    tokens <- struct{}{} // acquire a token
    list, err := Extract(url)
    <-tokens // release the token
    if err != nil {
        log.Print(err)
    }
    return list
}
*/

