package main

import (
	"flag"
	"log"
	"runtime"

	"tpler/backend"
	"tpler/common"
	"tpler/controller"
	"tpler/model"
)

var (
	buildInfo    = ""
	addr         = flag.String("s", ":8080", "server address")
	sessionPath  = flag.String("sp", "/tmp/hst_session", "session path")
	dbPath       = flag.String("db", "./db/sqlite3.db", "database path")
	resetAdmin   = flag.String("resetAdmin", "", "reset admin new password")
	oauth2Server = flag.String("o2", "http://127.0.0.1:8000", "oauth2 server")
	version      = flag.Bool("v", false, buildInfo)
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Flags() | log.Lshortfile)
	println("BuildInfo:", buildInfo)

	// 系统初始化
	if err := common.Init(); err != nil {
		common.LL.Log3Fatal(err)
	}

	// 初始化数据库
	if err := model.Init(*dbPath); err != nil {
		common.LL.Log3Fatal(err)
	}

	// 重设管理员密码
	if *resetAdmin != "" {
		if err := model.ResetAdmin(*resetAdmin); err != nil {
			common.LL.Log4Trace(err)
		}
		return
	}

	// 初始化后台程序
	if err := backend.Start(); err != nil {
		common.LL.Log3Fatal(err)
	}

	// 启动web服务
	controller.Start(*addr, *sessionPath, *oauth2Server)
}
