package links

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"

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
		return nil, err
	}
	request.Header.Add("Accept-Encoding", "gzip")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
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
		// 这里可以直接读取response吗？gzip 格式的 response可以直接读取吗？
		// 如果 response 可以直接读取，那么不需要放在 case “gzip” 里
		// if response.StatusCode != http.StatusOK {
		// 	response.Body.Close()
		// 	return nil, fmt.Errorf("getting %s: %s", url, response.Status)
		// }
		defer reader.Close()
	default:
		reader = response.Body
	}

	// ???????????????????????????
	// 如果可以直接读取不需要改动
	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, response.Status)
	}
	// 同理，这里的response可不可以直接读取
	// 如果可以直接读取不需要改动
	doc, err := html.Parse(response.Body)
	response.Body.Close()
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
