package controller

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/ohko/hst"
	"github.com/ohko/logger"
)

var ll = logger.NewLogger()

// Start 启动WEB服务
func Start(addr, sessionPath, oauth2Server string) {
	oauthServerHost = oauth2Server
	oauthStateString = time.Now().Format("20060102150405")

	// hst对象
	s := hst.New(nil)

	// 禁止显示Route日志
	// s.DisableRouteLog = true
	s.SetLogger(ioutil.Discard)

	// HTML模版
	s.SetDelims("{[{", "}]}")
	s.SetTemplatePath("./view/")

	// favicon.ico
	s.Favicon()

	// Session
	// s.SetSession(hst.NewSessionMemory())
	s.SetSession(hst.NewSessionFile("TPLER", sessionPath, time.Minute*30))

	// 静态文件
	s.StaticGzip("/public/", "./public/")

	// 注册自动路由
	s.RegisterHandle(nil,
		&IndexController{},
		&AdminController{},
		&Oauth2Controller{},
	)

	// 设置模版函数
	s.SetTemplateFunc(map[string]interface{}{
		"json": func(x interface{}) string {
			bs, err := json.Marshal(x)
			if err != nil {
				return err.Error()
			}
			return string(bs)
		},
	})

	// 启动web服务
	go s.ListenHTTP(addr)

	// 优雅关闭
	hst.Shutdown(time.Second*5, s)
}

type app struct{}

// 渲染错误页面
func (o *app) renderError(ctx *hst.Context, data interface{}) {
	ctx.HTML2(200, "layout/empty.html", data, "page/error.html")
}

func (o *app) loginSuccess(ctx *hst.Context, uid, user string) {
	ctx.SessionSet("", "/", "UID", uid, time.Minute*30)
	ctx.SessionSet("", "/", "User", user, time.Minute*30)
}
