package controller

import (
	"net/http"
	"time"

	"github.com/ohko/hst"
	"github.com/ohko/tplwww/model"
)

// AdminController 默认管理控制器
type AdminController struct {
	app
}

// 渲染模版
func (o *AdminController) render(ctx *hst.Context, name string, data interface{}, names ...string) {
	names = append(names, "layout/admin_header.html")
	names = append(names, "layout/admin_footer.html")
	ctx.HTML2(200, name, data, names...)
}

// Index ...
func (o *AdminController) Index(ctx *hst.Context) {
	if !o.CheckLogined(ctx) {
		http.Redirect(ctx.W, ctx.R, "/admin/login", 302)
		ctx.Close()
		return
	}
	o.render(ctx, "admin/index.html", nil)
}

// CheckLogined ...
func (o *AdminController) CheckLogined(ctx *hst.Context) bool {
	if v, _ := ctx.SessionGet("Name"); v != nil {
		return true
	}
	return false
}

// Login 登录
func (o *AdminController) Login(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		ctx.HTML2(200, "admin/login.html", nil, "layout/empty_header.html", "layout/empty_footer.html")
	}

	ctx.R.ParseForm()
	account := ctx.R.FormValue("Account")
	password := ctx.R.FormValue("Password")

	if account == "demo" && password == "demo" {
		ctx.SessionSet("UID", 1, time.Minute*30)
		ctx.SessionSet("Name", "管理员", time.Minute*30)
		ctx.JSON2(200, 0, "ok")
		return
	}
	ctx.JSON2(200, 1, "密码错误")
}

// Logout 登出
func (o *AdminController) Logout(ctx *hst.Context) {
	ctx.SessionDestory()
	http.Redirect(ctx.W, ctx.R, "/admin/login", 302)
	ctx.Close()
}

// GetMenu 默认菜单
func (o *AdminController) GetMenu(ctx *hst.Context) {
	if !o.CheckLogined(ctx) {
		ctx.JSON2(200, -1, "请先登陆")
		return
	}

	accountName, _ := ctx.SessionGet("Name")
	ctx.JSON2(200, 0, []model.Menu{
		model.Menu{Class: "fa-home", Text: "Home " + accountName.(string), Href: "/admin"},
		model.Menu{Class: "fa-dashboard", Text: "菜单组",
			Child: []model.Menu{
				model.Menu{Class: "fa-circle-o", Text: "功能1", Href: "/admin/func1"},
				model.Menu{Class: "fa-circle-o", Text: "功能2", Href: "/admin/func2"},
			}},
		model.Menu{Class: "fa-share", Text: "Exit " + accountName.(string), Href: "javascript:vueMenu.logout()"},
	})
}

// Func1 ...
func (o *AdminController) Func1(ctx *hst.Context) {
	user := &model.User{
		UID:  "123",
		Name: "hello",
	}
	o.render(ctx, "admin/func1.html", user)
}

// Func2 ...
func (o *AdminController) Func2(ctx *hst.Context) {
	o.render(ctx, "admin/func2.html", nil)
}
