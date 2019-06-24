package controller

import (
	"net/http"

	"github.com/ohko/hst"
	"github.com/ohko/tpler/model"
)

// AdminController 默认管理控制器
type AdminController struct {
	app
}

// 渲染模版
func (o *AdminController) render(ctx *hst.Context, data interface{}, names ...string) {
	ctx.HTML2(200, "layout/admin.html", data, names...)
}

func checkLogined(ctx *hst.Context) bool {
	if v, _ := ctx.SessionGet("User"); v != nil {
		return true
	}
	return false
}

// Login 登录
func (o *AdminController) Login(ctx *hst.Context) {
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

	o.loginSuccess(ctx, "1", user)
	ctx.JSON2(200, 0, "ok")
}

// Logout 登出
func (o *AdminController) Logout(ctx *hst.Context) {
	ctx.SessionDestory()
	http.Redirect(ctx.W, ctx.R, "/admin/login", 302)
	ctx.Close()
}

// Index ...
func (o *AdminController) Index(ctx *hst.Context) {
	if !checkLogined(ctx) {
		http.Redirect(ctx.W, ctx.R, "/admin/login", 302)
		ctx.Close()
		return
	}
	o.render(ctx, nil, "admin/index.html")
}

// GetMenu 默认菜单
func (o *AdminController) GetMenu(ctx *hst.Context) {
	if !checkLogined(ctx) {
		ctx.JSON2(200, -1, "请先登陆")
		return
	}

	user, _ := ctx.SessionGet("User")
	ctx.JSON2(200, 0, new(model.Menu).GetAdminMenu(user.(string)))
}

// Func1 ...
func (o *AdminController) Func1(ctx *hst.Context) {
	user := &model.User{
		UID:  "123",
		User: "hello",
	}
	o.render(ctx, user, "admin/func1.html")
}

// Func2 ...
func (o *AdminController) Func2(ctx *hst.Context) {
	o.render(ctx, nil, "admin/func2.html")
}
