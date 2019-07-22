package controller

import (
	"github.com/ohko/hst"
	"tpler/model"
)

// IndexController 默认主页控制器
type IndexController struct {
	app
}

// 渲染模版
func (o *IndexController) render(ctx *hst.Context, data interface{}, names ...string) {
	ctx.HTML2(200, "layout/default.html", data, names...)
}

// Index 默认主页
func (o *IndexController) Index(ctx *hst.Context) {
	o.render(ctx, nil, "page/index.html")
}

// Func1 ...
func (o *IndexController) Func1(ctx *hst.Context) {
	user := &model.User{
		UID:  "123",
		User: "hello",
	}
	o.render(ctx, user, "page/func1.html")
}

// Func2 ...
func (o *IndexController) Func2(ctx *hst.Context) {
	o.render(ctx, nil, "page/func2.html")
}
