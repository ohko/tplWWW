package controller

import (
	"net/http"
	"strings"

	"github.com/ohko/hst"
	"tpler/model"
)

// AdminController 默认管理控制器
type AdminController struct {
	app
}

// 渲染模版
func (o *AdminController) render(ctx *hst.Context, data interface{}, names ...string) {
	ctx.HTML2(200, "layout/admin.html", data, names...)
}

// Login 登录
func (o *AdminController) Login(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		ctx.HTML2(200, "layout/empty.html", map[string]string{
			"callback": ctx.R.FormValue("callback"),
		}, "admin/login.html")
	}

	ctx.R.ParseForm()
	user := ctx.R.FormValue("User")
	pass := ctx.R.FormValue("Password")
	callback := ctx.R.FormValue("callback")

	if err := users.Check(user, pass); err != nil {
		ctx.JSON2(200, 1, err.Error())
	}

	o.loginSuccess(ctx, "1", user)

	if strings.Contains(ctx.R.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
		ctx.JSON2(200, 0, "ok")
	} else if callback != "" {
		http.Redirect(ctx.W, ctx.R, callback, 302)
	} else {
		http.Redirect(ctx.W, ctx.R, "/admin/", 302)
	}
}

// Logout 登出
func (o *AdminController) Logout(ctx *hst.Context) {
	ctx.SessionDestory()
	http.Redirect(ctx.W, ctx.R, "/admin/login", 302)
	ctx.Close()
}

// Index ...
func (o *AdminController) Index(ctx *hst.Context) {
	user := &model.User{
		UID:  "123",
		User: "hello",
	}
	o.render(ctx, user, "admin/index.html")
}

// GetMenu 默认菜单
func (o *AdminController) GetMenu(ctx *hst.Context) {
	user, _ := ctx.SessionGet("User")
	ctx.JSON2(200, 0, new(model.Menu).GetAdminMenu(user.(string)))
}

// Form ...
func (o *AdminController) Form(ctx *hst.Context) {
	o.render(ctx, nil, "admin/form.html")
}

// Table ...
func (o *AdminController) Table(ctx *hst.Context) {
	o.render(ctx, nil, "admin/table.html")
}
