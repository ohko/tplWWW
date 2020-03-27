// 需要编译的资源文件
import '../public/css/css.css'

// 不需要编译的资源文件
const assets = [
   '../public/css/bootstrap.css',
   '../public/css/font-awesome.css',
   '../public/css/ionicons.css',
   '../public/css/AdminLTE.css',
   '../public/css/_all-skins.css',
   '../public/css/icheck.css',
   '../public/css/ag-grid.min.css',
   '../public/css/ag-theme-balham.min.css',
   '../public/css/toastr.min.css',

   '../public/js/jquery.min.js',
   '../public/js/bootstrap.min.js',
   '../public/js/adminlte.min.js',
   '../public/js/fastclick.js',
   '../public/js/icheck.min.js',
   '../public/js/ag-grid-enterprise.js',
   '../public/js/toastr.min.js',
]

import Vue from "vue"; Vue.config.productionTip = false
import Vuex from "vuex"; Vue.use(Vuex)
import VueRouter from "vue-router"; Vue.use(VueRouter)
import request from "./core/request.js";

// 依次加载资源文件
const loadAssets = () => {
   if (assets.length == 0) {
      Vue.use(request)

      // 资源文件加载结束后再初始化Vue
      new Vue({
         store: require('./core/store.js'),
         router: new VueRouter({
            routes: require('./core/route.js')
         }),
         created() { window.vue = this },
      }).$mount('#app')
      return
   }

   // 根据不同的文件类型加载
   const url = assets.splice(0, 1)[0]
   const exts = url.split(".")
   const ext = exts[exts.length - 1]
   if (ext == "css") {
      const css = document.createElement("link")
      css.href = url; css.rel = "stylesheet"; css.onload = loadAssets
      document.head.append(css)
   } else if (ext == "js") {
      const js = document.createElement("script")
      js.src = url; js.onload = loadAssets
      document.head.append(js)
   }
}
loadAssets()