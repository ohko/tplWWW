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

module.exports = {
   '/': Dashboard,
   '/admin': Dashboard,
   '/admin/dashboard': Dashboard,
   '/admin/login': Login,
   '/admin/logout': Logout,
   '/admin/form': Form,
   '/admin/table': Table,
   '/admin/password': Password,
   '/admin/success': Success,
   '/admin/error': Error,
   '/admin/user/list': AdminUserList,
   '/admin/user/add': AdminUserAdd,
   '/admin/user/edit': AdminUserEdit,
}