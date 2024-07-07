package app

import (
	"fmt"
	"os"
	"time"

	"github.com/andrewarrow/feedback/router"
)

func handleAsinPost(c *router.Context, asin string) {
	devMode := os.Getenv("DEV_MODE") == "true"
	one := c.One("vote", "where user_id=$1", c.User["id"])
	if len(one) > 0 {
		ca := one["created_at"].(int64)
		delta := time.Now().Unix() - ca
		if devMode {
			delta -= 25200
		}
		fmt.Println(ca, delta)
		if delta < 3600 {
			send := map[string]any{}
			send["delta"] = fmt.Sprintf("You can vote again in: %d minutes", int(float64(3600-delta)/60.0))
			c.SendContentAsJson(send, 422)
			return
		}
	}
	c.FreeFormUpdate("insert into votes (user_id, asin) values ($1,$2)",
		c.User["id"], asin)
	c.FreeFormUpdate("update products set votes=votes+1 where asin=$1", asin)
	one = c.One("product", "where asin=$1", asin)
	c.SendContentAsJson(one, 200)
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
	items := c.All("product", "order by votes desc", "")
	c.Title = fmt.Sprintf("Next Up Products")
	send["items"] = items
	c.SendContentInLayout("schedule.html", send, 200)
}
