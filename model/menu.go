package model

// Menu 菜单定义
type Menu struct {
	Class string // 图标类名
	Text  string // 文字
	Href  string // 链接
	Child []Menu `json:",omitempty"` // 子菜单
}
