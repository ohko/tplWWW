package iface

// IBackend ...
type IBackend interface {
	Register() error // 注册
	Status() string  // 状态
}
