package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ohko/hst"
	"github.com/ohko/logger"
	"github.com/ohko/tpler/model"
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
	s.SetSession(hst.NewSessionFile(sessionPath, time.Minute*30))

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

// CheckLogined ...
func (o *app) CheckLogined(ctx *hst.Context) bool {
	if v, _ := ctx.SessionGet("User"); v != nil {
		return true
	}
	return false
}

// Login 登录
func (o *app) Login(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		ctx.HTML2(200, "layout/empty.html", nil, "admin/login.html")
	}

	ctx.R.ParseForm()
	user := ctx.R.FormValue("User")
	pass := ctx.R.FormValue("Password")

	u := new(model.User)
	if err := u.Check(user, pass); err != nil {
		ctx.JSON2(200, 1, err.Error())
	}

	o.loginSuccess(ctx, "1", user)
	ctx.JSON2(200, 0, "ok")
}

func (o *app) loginSuccess(ctx *hst.Context, uid, user string) {
	ctx.SessionSet("UID", uid, time.Minute*30)
	ctx.SessionSet("User", user, time.Minute*30)
}

// Logout 登出
func (o *app) Logout(ctx *hst.Context) {
	ctx.SessionDestory()
	http.Redirect(ctx.W, ctx.R, "/admin/login", 302)
	ctx.Close()
}
