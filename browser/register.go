package browser

import (
	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	LogoutEvents()
	go checkTime()
	afterRegister := func(id int64) {
		Global.Location.Set("href", "/core/start")
	}
	afterLogin := func(id int64) {
		Global.Location.Set("href", "/core/start")
	}
	if Global.Start == "start.html" {
		a := wasm.NewAutoForm("add")
		a.Path = "/core/add"
		a.Clear = true
		a.Before = func() string {
			Document.Id("add-button").Set("value", "please wait...")
			return ""
		}
		a.After = func(content string) {
			Global.Location.Set("href", "/core/start")
		}
		Global.AddAutoForm(a)
	} else if Global.Start == "asin.html" || Global.Start == "schedule.html" {
		handleAsins()
	} else if Global.Start == "login.html" {
		Global.AutoForm("login", "core", nil, afterLogin)
	} else if Global.Start == "register.html" {
		Global.AutoForm("register", "core", nil, afterRegister)
	}
}

func LogoutEvents() {
	if Document.Id("logout") == nil {
		return
	}
	Global.Event("logout", Global.Logout("/core"))
}
