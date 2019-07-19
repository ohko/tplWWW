package model

// Menu 菜单定义
type Menu struct {
	Class string // 图标类名
	Text  string // 文字
	Href  string // 链接
	Child []Menu `json:",omitempty"` // 子菜单
}

// GetAdminMenu ...
func (o *Menu) GetAdminMenu(who string) []Menu {
	return []Menu{
		Menu{Class: "fa-home", Text: "仪表盘", Href: "/admin/"},
		Menu{Class: "fa-dashboard", Text: "菜单组",
			Child: []Menu{
				Menu{Class: "fa-circle-o", Text: "表单", Href: "/admin/form"},
				Menu{Class: "fa-circle-o", Text: "表格", Href: "/admin/table"},
			}},
		Menu{Class: "fa-share", Text: "退出:" + who, Href: "javascript:vueMenu.logout()"},
	}
}
