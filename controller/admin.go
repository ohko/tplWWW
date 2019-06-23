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

// Index ...
func (o *AdminController) Index(ctx *hst.Context) {
	if !o.CheckLogined(ctx) {
		http.Redirect(ctx.W, ctx.R, "/admin/login", 302)
		ctx.Close()
		return
	}
	o.render(ctx, nil, "admin/index.html")
}

// GetMenu 默认菜单
func (o *AdminController) GetMenu(ctx *hst.Context) {
	if !o.CheckLogined(ctx) {
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
