package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/andrewarrow/feedback/router"
)

func Core(c *router.Context, second, third string) {
	if second == "asin" && third != "" && c.Method == "POST" {
		handleAsinPost(c, third)
		return
	}
	// Why did the Amazon ASIN ID get sent to confession?
	if second == "asins" && third == "" && c.Method == "GET" {
		handleAsins(c, third)
		return
	}
	// Because it committed "A SIN" of having too many digits!
	if second == "asin" && third != "" && c.Method == "GET" {
		handleAsin(c, third)
		return
	}
	if second == "about" && third == "" && c.Method == "GET" {
		handleAboutUs(c)
		return
	}
	if second == "privacy" && third == "" && c.Method == "GET" {
		handlePrivacy(c)
		return
	}
	if second == "terms" && third == "" && c.Method == "GET" {
		handleTerms(c)
		return
	}
	if second == "register" && third == "" && c.Method == "GET" {
		handleRegister(c)
		return
	}
	if second == "login" && third == "" && c.Method == "GET" {
		handleLogin(c)
		return
	}
	if second == "register" && third == "" && c.Method == "POST" {
		router.HandleCreateUserAutoForm(c, "")
		return
	}
	if second == "login" && third == "" && c.Method == "POST" {
		router.HandleCreateSessionAutoForm(c)
		return
	}
	if second == "logout" && third == "" && c.Method == "DELETE" {
		router.DestroySession(c)
		return
	}
	if router.NotLoggedIn(c) {
		return
	}
	if second == "start" && third == "" && c.Method == "GET" {
		handleCoreStart(c)
		return
	}
	if second == "add" && third == "" && c.Method == "POST" {
		handleAddPost(c)
		return
	}
	c.NotFound = true
}

func handleIndex(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("welcome.html", send, 200)
}

func handleRegister(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("register.html", send, 200)
}
func handleLogin(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("login.html", send, 200)
}

func handlePrivacy(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("privacy.html", send, 200)
}
func handleTerms(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("terms.html", send, 200)
}
func handleAboutUs(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("about.html", send, 200)
}
func handleCoreStart(c *router.Context) {
	send := map[string]any{}
	items := c.All("product", "where user_id=$1 order by created_at desc", "", c.User["id"])
	send["items"] = items
	c.SendContentInLayout("start.html", send, 200)
}
func handleAddPost(c *router.Context) {
	c.ReadJsonBodyIntoParams()
	asin, _ := c.Params["asin"].(string)
	one := c.One("product", "where asin=$1", asin)
	send := map[string]any{}
	if len(one) > 0 {
		send["error"] = "try another asin"
		c.SendContentAsJson(send, 422)
		return
	}
	one = c.One("product", "where user_id=$1", c.User["id"])
	if len(one) > 0 {
		ca := one["created_at"].(int64)
		if time.Now().UTC().Unix()-ca < 86400 {
			send["error"] = "try again tomorrow"
			c.SendContentAsJson(send, 422)
			return
		}
	}

	url := "https://www.amazon.com/dp/" + asin

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	imageURL, title := parseAmazon(resp.Body)

	//fmt.Println("11", imageURL, title, asin)
	//c.FreeFormUpdate("update products set photo=$1,original_title=$2 where asin=$3", imageURL, title, asin)
	if imageURL == "" {
		send["error"] = "try another asin"
		c.SendContentAsJson(send, 422)
		return
	}
	c.Params["photo"] = imageURL
	c.Params["original_title"] = title

	c.Params["user_id"] = c.User["id"]
	c.Params["scheduled_for"] = time.Now()
	items := c.FreeFormSelect("select max(scheduled_for) as scheduled_for from products")
	if len(items) > 0 {
		max := items[0]["scheduled_for"]
		if max != nil {
			maxTime := max.(time.Time)
			fmt.Println("scheduled_for", maxTime)
			c.Params["scheduled_for"] = maxTime.Add(time.Hour * 24)
		}
	}
	msg := c.ValidateAndInsert("product")
	if msg != "" {
		send["error"] = "try another asin"
		c.SendContentAsJson(send, 422)
		return
	}
	c.SendContentAsJson(send, 200)
}
