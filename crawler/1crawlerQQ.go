package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"practice/golangLearning/links"
	"sync"
)

// tokens is a counting semaphore used to
// enforce a limit of 4 concurrent requests.
var tokens = make(chan struct{}, 4)
var maxDepth int
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(depth, url)
	if depth >= maxDepth {
		return
	}
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		seenLock.Lock()
		if seen[link] {
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
}

func main() {
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, syscall.SIGTERM, syscall.SIGTRAP)
	// Block until a signal is received.
	// fmt.Println("启动")
	// s := <-c
	// fmt.Println("Got signal: ", s)

	// flag 包实现了命令行参数的解析
	flag.IntVar(&maxDepth, "d", 2, "max crawl depth")
	flag.Parse()
	wg := &sync.WaitGroup{}
	for _, link := range flag.Args() {
		// tlink := link
		//fmt.Println("used for breakpoint")
		wg.Add(1)
		// go func() {
		// 	fmt.Println("start go func")
		// 	for s := range c {
		// 		fmt.Println("start range chan")
		// 		switch s {
		// 		case syscall.SIGTERM, syscall.SIGTRAP:
		// 			fmt.Println("退出", s)
		// 			exitFunc()
		// 		}
		// 	}
		// 	crawl(tlink, 0, wg)
		// }()
		go crawl(link, 0, wg)
	}
	wg.Wait()
}

// method after received SIGTERM
func exitFunc() {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}
