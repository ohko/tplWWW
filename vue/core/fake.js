// 模拟数据

const menu = {
   "no": 0,
   "data": [
      { "Class": "fa-home", "Text": "仪表盘", "Href": "#/admin/dashboard" },
      { "Class": "fa-home", "Text": "用户管理", "Href": "#/admin/user/list" },
      {
         "Class": "fa-dashboard", "Text": "示例", "Href": "", "Child": [
            { "Class": "fa-circle-o", "Text": "表单", "Href": "#/admin/form" },
            { "Class": "fa-circle-o", "Text": "表格", "Href": "#/admin/table" }]
      },
      { "Class": "fa-home", "Text": "系统配置", "Href": "#/admin/setting/list" },
      { "Class": "fa-home", "Text": "修改密码", "Href": "#/admin/password" },
      { "Class": "fa-share", "Text": "退出:admin", "Href": "#/admin/logout" }
   ],
}

const userList = {
   "no": 0, "data": [
      { "User": "user-1", "Email": "email1@qq.com" },
      { "User": "user-2", "Email": "email2@qq.com" },
   ]
}

const settingList = {
   "no": 0, "data": [
      { "Key": "key-1", "Int": 1, "String": "a", "Bool": true },
      { "Key": "key-2", "Int": 2, "String": "b", "Bool": false },
   ]
}

const userDetail = {
   "no": 0, "data": { "User": "user-1", "Email": "email1@qq.com" }
}

const settingDetail = {
   "no": 0, "data": { "Key": "key-1", "Int": 1, "String": "a", "Bool": true }
}

const login = { "no": 0, "data": "ok" }
const logout = { "no": 0, "data": "ok" }
const resultOK = { "no": 0, "data": "ok" }

module.exports = [
   { url: /\/admin\/get_adm_menu/, data: menu },
   { url: /\/admin\/login/, data: login },
   { url: /\/admin\/logout/, data: logout },
   { url: /\/admin_user\/add/, data: resultOK },
   { url: /\/admin_user\/list/, data: userList },
   { url: /\/admin_user\/delete/, data: resultOK },
   { url: /\/admin_user\/detail/, data: userDetail },
   { url: /\/admin_user\/edit/, data: resultOK },
   { url: /\/admin_setting\/add/, data: resultOK },
   { url: /\/admin_setting\/list/, data: settingList },
   { url: /\/admin_setting\/delete/, data: resultOK },
   { url: /\/admin_setting\/detail/, data: settingDetail },
   { url: /\/admin_setting\/edit/, data: resultOK },
]