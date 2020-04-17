package controller

import (
	"tpler/common"

	"github.com/ohko/hst"
)

// IndexController 默认主页控制器
type IndexController struct {
	controller
}

// 渲染模版
func (o *IndexController) render(ctx *hst.Context, data interface{}, names ...string) {
	ctx.HTML2(200, "layout/default.html", data, names...)
}

// Index 默认主页
func (o *IndexController) Index(ctx *hst.Context) {
	o.render(ctx, nil, "page/index.html")
}

// BuildInfo 编译信息
func (o *IndexController) BuildInfo(ctx *hst.Context) {
	ctx.W.Write([]byte(common.BuildInfo))
}

// Ws websocket例子
func (o *IndexController) Ws(ctx *hst.Context) {
	/* "github.com/gorilla/websocket"
	// 建立连接
	ws, err := websocket.Upgrade(ctx.W, ctx.R, nil, 1024, 1024)
	if err != nil {
		log.Println(err)
		http.NotFound(ctx.W, ctx.R)
		ctx.Close()
	}

	// 接收数据
	go func() {
		var msg hst.JSONData
		for {
			if err := ws.ReadJSON(&msg); err != nil {
				common.LL.Log2Error(err)
				break
			}
			common.LL.Log0Debug(msg.No, msg.Data)
		}
	}()

	// 发送数据
	for {
		time.Sleep(time.Second * 3)
		if err := ws.WriteJSON(hst.JSONData{No: 0, Data: "ok"}); err != nil {
			break
		}
	}
	// */
}
