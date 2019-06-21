package controller

import (
	"github.com/ohko/hst"
	"github.com/ohko/tpler/model"
)

// IndexController 默认主页控制器
type IndexController struct {
	app
}

// 渲染模版
func (o *IndexController) render(ctx *hst.Context, name string, data interface{}, names ...string) {
	names = append(names, "layout/default_header.html")
	names = append(names, "layout/default_footer.html")
	ctx.HTML2(200, name, data, names...)
}

// Index 默认主页
func (o *IndexController) Index(ctx *hst.Context) {
	o.render(ctx, "page/index.html", nil)
}

// Func1 ...
func (o *IndexController) Func1(ctx *hst.Context) {
	user := &model.User{
		UID:  "123",
		Name: "hello",
	}
	o.render(ctx, "page/func1.html", user)
}

// Func2 ...
func (o *IndexController) Func2(ctx *hst.Context) {
	o.render(ctx, "page/func2.html", nil)
}
