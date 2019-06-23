package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/ohko/hst"
	"github.com/ohko/logger"
	"github.com/ohko/tpler/controller"
)

var (
	addr        = "0.0.0.0:8080"
	sessionPath = "/tmp/hst_session"
	ll          = logger.NewLogger()
)

func start() {
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
		&controller.IndexController{},
		&controller.AdminController{},
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

func main() {
	ll.Log4Trace("=== START WEB SERVER ===")

	start()
}
