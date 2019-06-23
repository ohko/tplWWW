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
		Menu{Class: "fa-home", Text: "Home", Href: "/admin"},
		Menu{Class: "fa-dashboard", Text: "菜单组",
			Child: []Menu{
				Menu{Class: "fa-circle-o", Text: "功能1", Href: "/admin/func1"},
				Menu{Class: "fa-circle-o", Text: "功能2", Href: "/admin/func2"},
			}},
		Menu{Class: "fa-share", Text: "Logout:" + who, Href: "javascript:vueMenu.logout()"},
	}
}
