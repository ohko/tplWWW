// Package backend 这个目录放一些后台自动化执行的代码
package backend

import (
	"tpler/backend/demo1"
	"tpler/backend/demo2"
	"tpler/backend/iface"
)

// ...
var (
	Backends = []iface.IBackend{
		&demo1.Demo1{},
		&demo2.Demo2{},
	}
)

// Start ...
func Start() error {
	// 取消注释，可激活backend服务
	// for _, v := range Backends {
	// 	if err := v.Register(); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
