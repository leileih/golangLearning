package main

import (
	"encoding/csv"
	"exercise/golangLearning/links"
	"exercise/golangLearning/protobuf"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/golang/protobuf/proto"
)

// tokens is a counting semaphore used to
// enforce a limit of 4 concurrent requests.
var tokens = make(chan struct{}, 4)
var maxDepth int
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	// print output to console
	fmt.Println(depth, url)

	// write to file -- start
	// transfer result to protobuf format
	link := protobuf.Link{
		Depth: int32(depth),
		Url:   url,
	}

	// 对数据进行序列化
	data, err := proto.Marshal(&link)
	if err != nil {
		log.Fatalln("Mashal data error:", err)
	}

	// create output file
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatalln("Cannot create file", err)
	}
	defer file.Close()

	// write to csv file
	w := csv.NewWriter(file)
	defer w.Flush()
	// convert []byte to []strings
	// might be improved
	for depth, url := range data {
		d := string(depth)
		u := string(url)
		output := [][]string{
			{d, u},
		}
		w.WriteAll(output)
	}
	// w.WriteAll(output)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	// write to file -- end

	/*
		// 对已经序列化的数据进行反序列化
		// 验证数据序列化是否正确
		var target protobuf.Link
		err = proto.Unmarshal(data, &target)
		if err != nil {
			log.Fatalln("UnMashal data error:", err)
		}

		// fmt.Println("depth: ", target.Depth, "url: ", target.Url)
	*/

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
	// signal.Notify(c, syscall.SIGTERM, os.Interrupt)
	// Block until a signal is received.
	// fmt.Println("启动")
	// s := <-c
	// fmt.Println("Got signal: ", s)
	// exitFunc()

	// flag 包实现了命令行参数的解析
	flag.IntVar(&maxDepth, "d", 3, "max crawl depth")
	flag.Parse()
	wg := &sync.WaitGroup{}
	for _, link := range flag.Args() {
		// tlink := link
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
	fmt.Println("保存进度...")
	fmt.Println("结束退出...")
	os.Exit(0)
}
