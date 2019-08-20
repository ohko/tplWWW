import '../public/css/bootstrap.css'
import '../public/css/font-awesome.css'
import '../public/css/ionicons.css'
import '../public/css/AdminLTE.css'
import '../public/css/dataTables.bootstrap.css'
import '../public/css/_all-skins.css'
import '../public/css/icheck.css'
import '../public/css/css.css'

window.$ = window.jQuery = require('../public/js/jquery.min.js')
window.Vue = require('../public/js/vue.min.js')
window.Vuex = require('../public/js/vuex.js')
window.VueRouter = require('../public/js/vue-router.js')
require('../public/js/bootstrap.min.js')
require('../public/js/adminlte.min.js')
require('../public/js/fastclick.js')
require('../public/js/jquery.dataTables.js')
require('../public/js/dataTables.bootstrap.js')
require('../public/js/icheck.min.js')

Vue.use(VueRouter)
Vue.config.productionTip = false

import request from "./core/request.js"
Vue.use(request)

new Vue({
   store: require('./core/store.js'),
   router: new VueRouter({
      routes: require('./core/route.js')
   }),
   created() {
      window.vue = this
   }
}).$mount('#app')
