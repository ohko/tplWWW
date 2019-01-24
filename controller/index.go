package controller

import (
	"tplwww/model"

	"github.com/ohko/hst"
)

// IndexController 默认主页控制器
type IndexController struct{}

// Index 默认主页
func (o *IndexController) Index(ctx *hst.Context) {
	// ctx.RenderFiles(nil, "./public/index.html", "./public/layoutDefault.html")
	ctx.LayoutRender("layoutDefault", nil, "./view/index.html")
}

// Func1 ...
func (o *IndexController) Func1(ctx *hst.Context) {
	user := &model.User{
		UID:  "123",
		Name: "hello",
	}
	ctx.LayoutRender("layoutDefault", user, "./view/func1.html")
}

// Func2 ...
func (o *IndexController) Func2(ctx *hst.Context) {
	ctx.LayoutRender("layoutDefault", nil, "./view/func2.html")
}
