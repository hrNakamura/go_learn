package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

//TODO 要修正
func main() {
	var url string
	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		fmt.Println("no url")
		return
	}
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("words: %v\n", words)
	fmt.Printf("images: %v", images)
}

// CountWordsAndImages 指定したHTMLの単語数と画像数をカウントする
func CountWordsAndImages(url string) (words, imagaes int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, imagaes = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	visit(&words, &images, n)
	return
}

func visit(words, images *int, n *html.Node) (w, i int) {
	*words += countWords(n)
	*images += countImages(n)
	//TODO nodeを返す必要はない
	//タプルでは+=ができないため
	for c := n.FirstChild; c != nil; c = n.NextSibling {
		w, i = visit(words, images, c)
		*words += w
		*images += i
	}
	return *words, *images
}

func countWords(n *html.Node) int {
	var count int
	if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			count++
		}
	}
	return count
}

func countImages(n *html.Node) int {
	var count int
	if n.Type == html.ElementNode && n.Data == "img" {
		// for _, img := range n.Attr {
		// 	if img.Key == "src" {
		count++
		// 	}
		// }
	}
	return count
}
