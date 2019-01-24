package main

import (
	"encoding/json"
	"log"
	"time"
	"tplwww/controller"

	"github.com/ohko/hst"
)

var addr = "0.0.0.0:8080"

func main() {
	log.Println(log.Ldate | log.Ltime | log.Lshortfile)

	// hst对象
	s := hst.New(nil)

	// favicon.ico
	s.Favicon()

	// Session
	s.SetSession(hst.NewSessionMemory())

	// 静态文件
	s.StaticGzip("/public/", "./public/")

	// 注册自动路由
	s.RegisterHandle(
		&controller.IndexController{},
		&controller.AdminController{},
	)

	// 重定义转义符
	s.SetDelims("<!-- {{", "}} -->")

	// 设置Layout模版
	s.SetLayout("layoutDefault", "./view/layout/default.html")
	s.SetLayout("layoutAdmin", "./view/layout/admin.html")
	s.SetLayout("layoutEmpty", "./view/layout/empty.html")

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
