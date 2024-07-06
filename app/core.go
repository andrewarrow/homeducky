package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Core(c *router.Context, second, third string) {
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

func handleCoreStart(c *router.Context) {

	send := map[string]any{}
	c.SendContentInLayout("start.html", send, 200)
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
func handleStart(c *router.Context) {
	send := map[string]any{}
	items := c.All("product", "where user_id=$1 order by created_at desc", "", c.User["id"])
	send["items"] = items
	c.SendContentInLayout("start.html", send, 200)
}
func handleAddPost(c *router.Context) {
	c.ReadJsonBodyIntoParams()
	c.Params["user_id"] = c.User["id"]
	c.ValidateAndInsert("product")
	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
