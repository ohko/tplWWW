package controller

import (
	"net/http"
	"time"
	"tplwww/model"

	"github.com/ohko/hst"
)

// Controller ...
type Controller interface {
	CheckLogined(c *hst.Context) bool // 判断是否已经登陆
	DoLogout(c *hst.Context) bool     // 登出，由/logout调用
	DoLogin(c *hst.Context)           // 登陆，在/login页面调用
	GetMenu(c *hst.Context)           // 菜单，在加载页面时调用
}

// AdminController 默认管理控制器
type AdminController struct{}

// Index ...
func (o *AdminController) Index(ctx *hst.Context) {
	if !o.CheckLogined(ctx) {
		http.Redirect(ctx.W, ctx.R, "/admin/login", 302)
		ctx.Close()
		return
	}
	// o.RenderFilesByTplDefault(c,nil, "../sysStatic/index.html")
	ctx.LayoutRender("layoutAdmin", nil, "./view/admin/index.html")
}
func (o *AdminController) CheckLogined(ctx *hst.Context) bool {
	if v, _ := ctx.SessionGet("Name"); v != nil {
		return true
	}
	return false
}

// Login 登录
func (o *AdminController) Login(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		ctx.LayoutRender("layoutEmpty", nil, "./view/admin/login.html")
	}

	ctx.R.ParseForm()
	account := ctx.R.FormValue("Account")
	password := ctx.R.FormValue("Password")

	if account == "demo" && password == "demo" {
		ctx.SessionSet("UID", 1, time.Minute*30)
		ctx.SessionSet("Name", "管理员", time.Minute*30)
		ctx.JSON2(0, "ok")
		return
	}
	ctx.JSON2(1, "密码错误")
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
		ctx.JSON2(-1, "请先登陆")
		return
	}

	accountName, _ := ctx.SessionGet("Name")
	ctx.JSON2(0, []model.Menu{
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
	ctx.LayoutRender("layoutAdmin", user, "./view/func1.html")
}

// Func2 ...
func (o *AdminController) Func2(ctx *hst.Context) {
	ctx.LayoutRender("layoutAdmin", nil, "./view/func2.html")
}
