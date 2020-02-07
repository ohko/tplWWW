// Package backend 这个目录放一些后台自动化执行的代码
package backend

import (
	"tpler/backend/demo1"
	"tpler/backend/iface"
)

// ...
var (
	Backends = []iface.IBackend{
		&demo1.Demo1{},
	}
)

// Start ...
func Start() error {
	for _, v := range Backends {
		if err := v.Register(); err != nil {
			return err
		}
	}

	return nil
}
