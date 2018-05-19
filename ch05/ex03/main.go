package main

import (
	"io"
	"os"

	. "github.com/golang-study/ch05/errorHandling"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	CheckError(err)
	ShowTexts(os.Stdout, doc)
}

func ShowTexts(w io.Writer, n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.TextNode {
		// s := strings.TrimSpace(n.Data)
		io.WriteString(w, n.Data)
		// s := strings.TrimSpace(n.Data)
		// s = strings.Trim(s, " \t\n")
		// textList = append(textList, s)
	}

	if !(n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style")) {
		ShowTexts(w, n.FirstChild)
	}

	ShowTexts(w, n.NextSibling)
	// return ShowTexts(textList, n.NextSibling)
}
