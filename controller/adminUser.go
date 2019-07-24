package controller

import (
	"net/http"

	"tpler/model"
	"tpler/util"

	"github.com/ohko/hst"
)

// AdminUserController 用户管理控制器
type AdminUserController struct {
	controller
}

// List 用户列表
func (o *AdminUserController) List(ctx *hst.Context) {
	us, err := users.List()
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	o.renderAdmin(ctx, map[string]interface{}{"us": us}, "admin/user/list.html")
}

// Add 增加用户
func (o *AdminUserController) Add(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		o.renderAdmin(ctx, nil, "admin/user/add.html")
	}

	u := &model.User{
		User:  ctx.R.FormValue("User"),
		Pass:  string(util.Hash([]byte(ctx.R.FormValue("Pass")))),
		Email: ctx.R.FormValue("Email"),
	}
	if err := users.Save(u); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	http.Redirect(ctx.W, ctx.R, "/admin_user/list", http.StatusFound)
}

// Edit 编辑用户
func (o *AdminUserController) Edit(ctx *hst.Context) {
	user := ctx.R.FormValue("User")
	u, err := users.Get(user)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	if ctx.R.Method == "GET" {
		o.renderAdmin(ctx, u, "admin/user/edit.html")
	}

	pass := ctx.R.FormValue("Pass")
	if pass != "" {
		u.Pass = string(util.Hash([]byte(pass)))
	}
	u.Email = ctx.R.FormValue("Email")
	if err := u.Save(u); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	http.Redirect(ctx.W, ctx.R, "/admin_user/list", http.StatusFound)
}

// Delete 删除用户
func (o *AdminUserController) Delete(ctx *hst.Context) {
	user := ctx.R.FormValue("User")
	if err := users.Delete(user); err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	http.Redirect(ctx.W, ctx.R, "/admin_user/list", http.StatusFound)
}
