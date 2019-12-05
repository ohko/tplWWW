package controller

import (
	"net/http"
	"strconv"
	"tpler/model"

	"github.com/ohko/hst"
)

// AdminSettingController 配置管理控制器
type AdminSettingController struct {
	controller
}

// List 配置列表
func (o *AdminSettingController) List(ctx *hst.Context) {
	us, err := model.DBSetting.List()
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	if ctx.IsAjax() {
		ctx.JSON2(200, 0, us)
	}

	o.renderAdmin(ctx, map[string]interface{}{"us": us}, "admin/setting/list.html")
}

// Add 增加配置
func (o *AdminSettingController) Add(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		o.renderAdmin(ctx, nil, "admin/setting/add.html")
	}

	i, _ := strconv.Atoi(ctx.R.FormValue("Int"))
	u := &model.Setting{
		Key:    ctx.R.FormValue("Key"),
		Desc:   ctx.R.FormValue("Desc"),
		Int:    i,
		String: ctx.R.FormValue("String"),
		Bool:   ctx.R.FormValue("Bool") == "on" || ctx.R.FormValue("Bool") == "true",
	}
	if err := model.DBSetting.Save(u); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	if ctx.IsAjax() {
		ctx.JSON2(200, 0, "ok")
	}
	http.Redirect(ctx.W, ctx.R, "/admin_setting/list", http.StatusFound)
}

// Detail 查看配置
func (o *AdminSettingController) Detail(ctx *hst.Context) {
	key := ctx.R.FormValue("Key")
	u, err := model.DBSetting.Get(key)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	if ctx.IsAjax() {
		ctx.JSON2(200, 0, u)
	}
	o.renderAdmin(ctx, u, "admin/setting/edit.html")
}

// Edit 编辑配置
func (o *AdminSettingController) Edit(ctx *hst.Context) {
	key := ctx.R.FormValue("Key")
	u, err := model.DBSetting.Get(key)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	if ctx.R.Method == "GET" {
		if ctx.IsAjax() {
			ctx.JSON2(200, 0, u)
		}
		o.renderAdmin(ctx, u, "admin/setting/edit.html")
	}

	u.Desc = ctx.R.FormValue("Desc")
	u.Int, _ = strconv.Atoi(ctx.R.FormValue("Int"))
	u.String = ctx.R.FormValue("String")
	u.Bool = (ctx.R.FormValue("Bool") == "on" || ctx.R.FormValue("Bool") == "true")
	if err := u.Save(u); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	if ctx.IsAjax() {
		ctx.JSON2(200, 0, "ok")
	}
	http.Redirect(ctx.W, ctx.R, "/admin_setting/list", http.StatusFound)
}

// Delete 删除配置
func (o *AdminSettingController) Delete(ctx *hst.Context) {
	key := ctx.R.FormValue("Key")
	if err := model.DBSetting.Delete(key); err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	if ctx.IsAjax() {
		ctx.JSON2(200, 0, "ok")
	}
	http.Redirect(ctx.W, ctx.R, "/admin_setting/list", http.StatusFound)
}
