package app

import (
	"fmt"

	"github.com/andrewarrow/feedback/router"
)

func handleAsinPost(c *router.Context, asin string) {
	send := map[string]any{}
	c.FreeFormUpdate("update products set votes=votes+1 where asin=$1", asin)
	c.SendContentAsJson(send, 200)
}
func handleAsin(c *router.Context, asin string) {
	send := map[string]any{}
	item := c.One("product", "where asin=$1", asin)
	if len(item) == 0 {
		c.SendContentInLayout("404.html", send, 404)
		return
	}
	c.LayoutMap["og"] = item["photo"]
	c.Title = fmt.Sprintf("%s | homeducky.com", item["title"])
	send["item"] = item
	c.SendContentInLayout("asin.html", send, 200)
}
func handleAsins(c *router.Context, asin string) {
	send := map[string]any{}
	items := c.All("product", "order by scheduled_for desc", "")
	c.Title = fmt.Sprintf("Next Up Products")
	send["items"] = items
	c.SendContentInLayout("schedule.html", send, 200)
}
