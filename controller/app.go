package controller

import (
	"net/http"
	"time"

	"github.com/ohko/hst"
	"github.com/ohko/tpler/model"
)

type app struct{}

// CheckLogined ...
func (o *app) CheckLogined(ctx *hst.Context) bool {
	if v, _ := ctx.SessionGet("User"); v != nil {
		return true
	}
	return false
}

// Login 登录
func (o *app) Login(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		ctx.HTML2(200, "layout/empty.html", nil, "admin/login.html")
	}

	ctx.R.ParseForm()
	user := ctx.R.FormValue("User")
	pass := ctx.R.FormValue("Password")

	u := new(model.User)
	if err := u.Check(user, pass); err != nil {
		ctx.JSON2(200, 1, err.Error())
	}

	ctx.SessionSet("UID", 1, time.Minute*30)
	ctx.SessionSet("User", user, time.Minute*30)
	ctx.JSON2(200, 0, "ok")
}

// Logout 登出
func (o *app) Logout(ctx *hst.Context) {
	ctx.SessionDestory()
	http.Redirect(ctx.W, ctx.R, "/admin/login", 302)
	ctx.Close()
}
