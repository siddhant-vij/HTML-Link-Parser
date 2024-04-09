package parser

import (
	"io"

	"golang.org/x/net/html"
)

// Link represents a hyperlink (<a href="...">text</a>) with an Href.
type Link struct {
	Href string
}

// Parse will take in an HTML document and will return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var links []Link
	getLinks(doc, &links)
	return links, nil
}

func getLinks(n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				*links = append(*links, Link{
					Href: a.Val,
				})
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getLinks(c, links)
	}
}
