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
		Menu{Class: "fa-home", Text: "用户管理", Href: "/admin_user/list"},
		Menu{Class: "fa-dashboard", Text: "示例",
			Child: []Menu{
				Menu{Class: "fa-circle-o", Text: "表单", Href: "/admin/form"},
				Menu{Class: "fa-circle-o", Text: "表格", Href: "/admin/table"},
			}},
		Menu{Class: "fa-home", Text: "修改密码", Href: "/admin/password"},
		Menu{Class: "fa-share", Text: "退出:" + who, Href: "javascript:vueMenu.logout()"},
	}
}
