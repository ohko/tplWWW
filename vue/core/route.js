// 路由

import Dashboard from '../pages/dashboard.vue'
import Login from '../pages/login.vue'
import Logout from '../pages/logout.vue'

import Form from '../pages/form.vue'
import Table from '../pages/table.vue'
import Password from '../pages/password.vue'
import Success from '../pages/success.vue'
import Error from '../pages/error.vue'

import AdminUserList from '../user/list.vue'
import AdminUserAdd from '../user/add.vue'
import AdminUserEdit from '../user/edit.vue'

module.exports = [
   { path: '/', component: Dashboard },
   { path: '/admin', component: Dashboard },
   { path: '/admin/dashboard', component: Dashboard },
   { path: '/admin/login', name: "login", component: Login },
   { path: '/admin/logout', component: Logout },
   { path: '/admin/form', component: Form },
   { path: '/admin/table', component: Table },
   { path: '/admin/password', component: Password },
   { path: '/admin/success', component: Success },
   { path: '/admin/error', component: Error },
   { path: '/admin/user/list', component: AdminUserList },
   { path: '/admin/user/add', component: AdminUserAdd },
   { path: '/admin/user/edit/:user', component: AdminUserEdit },
   { path: '*', component: Error },
]