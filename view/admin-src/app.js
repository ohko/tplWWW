import '../../public/css/bootstrap.min.css'
import '../../public/css/font-awesome.min.css'
import '../../public/css/ionicons.min.css'
import '../../public/css/AdminLTE.min.css'
import '../../public/css/dataTables.bootstrap.min.css'
import '../../public/css/_all-skins.min.css'
import '../../public/css/icheck.css'
import '../../public/css/css.css'

window.$ = window.jQuery = require('../../public/js/jquery.min.js')
window.Vue = require('../../public/js/vue.min.js')
require('../../public/js/bootstrap.min.js')
require('../../public/js/adminlte.min.js')
require('../../public/js/fastclick.js')
require('../../public/js/jquery.dataTables.js')
require('../../public/js/dataTables.bootstrap.js')
require('../../public/js/icheck.min.js')

Vue.config.productionTip = false
const routes = require('./core/route.js')

new Vue({
   data: {
      currentRoute: window.location.pathname
   },
   render(h) {
      return h(this.ViewComponent)
   },
   computed: {
      ViewComponent() {
         return routes[this.currentRoute] || routes['/']
      }
   },
   created() {
      window.addEventListener('popstate', _ => { this.$emit("go", location.href) }, false);
   },
   methods: {
      go(href) {
         try { event.preventDefault() } catch (e) { }
         var path = href.split("?")[0]
         this.$root.currentRoute = path;
         try { window.history.pushState(null, "", href); } catch (e) { }
         this.$emit("go", href)
      }
   },
}).$mount('#app')
