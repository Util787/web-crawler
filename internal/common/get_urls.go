package common

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// in current state func return both internal and external urls
// it returns every url (not only unique)
func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	urls := []string{}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	treeNode, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	for n := range treeNode.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						continue
					}
					urls = append(urls, baseURL.ResolveReference(href).String())
					break
				}
			}
		}
	}

	return urls, nil
}
