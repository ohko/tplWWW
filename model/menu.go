package model

// Menu 菜单定义
type Menu struct {
	Class string // 图标类名
	Text  string // 文字
	Href  string // 链接
	Child []Menu `json:",omitempty"` // 子菜单
}

// GetAdminMenu ...
func (Menu) GetAdminMenu(who string) []Menu {
	return []Menu{
		{Class: "fa-home", Text: "仪表盘", Href: "/admin/"},
		{Class: "fa-home", Text: "用户管理", Href: "/admin_user/list"},
		{Class: "fa-dashboard", Text: "示例",
			Child: []Menu{
				{Class: "fa-circle-o", Text: "表单", Href: "/admin/form"},
				{Class: "fa-circle-o", Text: "表格", Href: "/admin/table"},
			}},
		{Class: "fa-home", Text: "系统配置", Href: "/admin_setting/list"},
		{Class: "fa-home", Text: "修改密码", Href: "/admin/password"},
		{Class: "fa-share", Text: "退出:" + who, Href: "javascript:vueMenu.logout()"},
	}
}

// GetAdmMenu ...
func (Menu) GetAdmMenu(who string) []Menu {
	return []Menu{
		{Class: "fa-home", Text: "仪表盘", Href: "#/admin/dashboard"},
		{Class: "fa-home", Text: "用户管理", Href: "#/admin/user/list"},
		{Class: "fa-dashboard", Text: "示例",
			Child: []Menu{
				{Class: "fa-circle-o", Text: "表单", Href: "#/admin/form"},
				{Class: "fa-circle-o", Text: "表格", Href: "#/admin/table"},
			}},
		{Class: "fa-home", Text: "系统配置", Href: "#/admin/setting/list"},
		{Class: "fa-home", Text: "修改密码", Href: "#/admin/password"},
		{Class: "fa-share", Text: "退出:" + who, Href: "#/admin/logout"},
	}
}
