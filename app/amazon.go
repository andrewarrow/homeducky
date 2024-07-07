package app

import (
	"io"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

func parseAmazon(reader io.Reader) (string, string) {

	bodyBytes, _ := ioutil.ReadAll(reader)
	bodyString := string(bodyBytes)
	//fmt.Println(bodyString)

	tokenizer := html.NewTokenizer(strings.NewReader(bodyString))

	var (
		title    string
		imageURL string
	)

	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return imageURL, title
		case tt == html.StartTagToken || tt == html.SelfClosingTagToken:
			t := tokenizer.Token()

			// Check for the product title
			if t.Data == "span" {
				for _, a := range t.Attr {
					if a.Key == "id" && a.Val == "productTitle" {
						tokenizer.Next() // Move to the text within the tag
						title = strings.TrimSpace(html.UnescapeString(string(tokenizer.Text())))
					}
				}
			}

			// Check for the product image
			if t.Data == "img" {
				for _, a := range t.Attr {
					if a.Key == "id" && a.Val == "landingImage" {
						for _, imgAttr := range t.Attr {
							if imgAttr.Key == "src" {
								imageURL = imgAttr.Val
							}
						}
					}
				}
			}
		}
	}

	return imageURL, title
}
