// 模拟数据

const menu = {
   "no": 0,
   "data": [
      { "Class": "fa-home", "Text": "仪表盘", "Href": "/admin/dashboard" },
      { "Class": "fa-home", "Text": "用户管理", "Href": "/admin/user/list" },
      {
         "Class": "fa-dashboard", "Text": "示例", "Href": "", "Child": [
            { "Class": "fa-circle-o", "Text": "表单", "Href": "/admin/form" },
            { "Class": "fa-circle-o", "Text": "表格", "Href": "/admin/table" }]
      },
      { "Class": "fa-home", "Text": "修改密码", "Href": "/admin/password" },
      { "Class": "fa-share", "Text": "退出:admin", "Href": "/admin/logout" }
   ],
}

const login = { "no": 0, "data": "ok" }
const logout = { "no": 0, "data": "ok" }

module.exports = [
   { url: /\/admin\/get_menu/, data: menu },
   { url: /\/admin\/login/, data: login },
   { url: /\/admin\/logout/, data: logout },
]