package links

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	// For control over HTTP client headers, redirect policy, and other settings, create a Client:
	// https://golang.org/pkg/net/http/
	client := new(http.Client)

	request, err := http.NewRequest("Get", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error when requesting %s: %s", url, err)
	}
	request.Header.Add("Accept-Encoding", "gzip")

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println("error when debugging request: ", url, err)
	}
	fmt.Println(string(requestDump))

	response, err := client.Do(request)
	if err != nil {
		//fmt.Println("shown - error here")
		return nil, fmt.Errorf("response from %s: %s", url, err)
	}

	// Save a copy of this response for debugging.
	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		fmt.Println("error when debugging response: ", url, err)
	}
	fmt.Println(string(responseDump))

	// 判断 response 状态应该放在读取 response 之后
	if response.StatusCode != http.StatusOK {
		//fmt.Println(string(response.Status))
		response.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, response.Status)
	}
	defer response.Body.Close()

	// Check that the server actual sent compressed data
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, fmt.Errorf("reading gzip page %s: %v", url, err)
		}
		defer reader.Close()
	default:
		reader = response.Body
	}

	// // 判断 response 状态应该放在读取response之后
	// if response.StatusCode != http.StatusOK {
	// 	response.Body.Close()
	// 	return nil, fmt.Errorf("getting %s: %s", url, response.Status)
	// }

	// html.Parse 读取的是 io.Reader，即 response.Body 和 reader 同类型
	// 这里可以直接读取 html.Parse(reader)
	doc, err := html.Parse(reader)
	//response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// Extract links from content read
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := response.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
