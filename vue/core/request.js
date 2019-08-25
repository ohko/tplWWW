// 异步请求

const FAKE = process.env.NODE_ENV == 'development'
const BASE_URL = ''
// const BASE_URL = 'http://127.0.0.1:8080'

const getJSON = (url, params, callback) => {
   if (FAKE) {
      const fake = require("./fake.js")
      for (var i = 0; i < fake.length; i++) {
         if (fake[i].url.test(url)) return callback(fake[i].data)
      }
   }

   $.getJSON(BASE_URL + url, params, x => {
      if (x.no == -1 && window.vue) return window.vue.$router.push({ name: 'login', params: { callback: encodeURIComponent(location.href) } })
      if (x.no != 0) return alert(x.data)
      callback(x)
   });
}

const post = (url, params, callback) => {
   if (FAKE) {
      const fake = require("./fake.js")
      for (var i = 0; i < fake.length; i++) {
         if (fake[i].url.test(url)) return callback(fake[i].data)
      }
   }

   $.post(BASE_URL + url, params, x => {
      if (x.no == -1 && window.vue) return window.vue.$router.push({ name: 'login', params: { callback: encodeURIComponent(location.href) } })
      if (x.no != 0) return alert(x.data)
      callback(x)
   });
}

export default class Request {
   static install(Vue, options) {

      $.ajaxSetup({
         dataType: 'json',
         crossDomain: true,
         xhrFields: {
            withCredentials: true
         }
         // username: 'test',
         // password: 'test'
      });

      Vue.prototype.$post = post
      Vue.prototype.$getJSON = getJSON
   }
}