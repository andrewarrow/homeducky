package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/router"
)

func Sitemap(c *router.Context, second, third string) {
	if second == "products" && third == "" && c.Method == "GET" {
		handleSitemap(c)
		return
	}
	c.NotFound = true
}

func handleSitemap(c *router.Context) {
	urls := []string{}
	list := c.All("product", "order by id", "")
	for _, item := range list {
		thing := fmt.Sprintf(itemLayout, item["asin"], time.Now().Format("2006-01-02"))
		urls = append(urls, fmt.Sprintf(layout, thing))
	}
	send := strings.Join(urls, "\n")
	//c.Writer.WriteHeader(200)
	c.Writer.Header().Set("Content-Type", "application/xml")
	c.Writer.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + send + `</urlset>`))
}

func handleZipSitemap(c *router.Context) {
	c.Writer.Header().Set("Content-Type", "application/xml")
	buffer := []string{}
	for i := 0; i < 25; i++ {
		thing := fmt.Sprintf("<sitemap><loc>https://homeducky.com/%d</loc></sitemap>", i)
		buffer = append(buffer, thing)
	}
	theLetters := strings.Join(buffer, "\n")
	c.Writer.Write([]byte(containerSitemap + theLetters + "</sitemapindex>"))
}

var containerSitemap = `<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <sitemap><loc>/all/sitemap/a</loc></sitemap>`

var itemLayout = `<loc>https://homeducky.com/core/asin/%s</loc><lastmod>%s</lastmod>`
var layout = `<url>%s<changefreq>weekly</changefreq><priority>0.8</priority></url>`
