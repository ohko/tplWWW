// 异步请求

const FAKE = process.env.NODE_ENV == 'development'
const BASE_URL = ''
// const BASE_URL = 'http://127.0.0.1:8080'

const getJSON = (url, params, success, error) => {
   if (FAKE) {
      const fake = require("./fake.js")
      for (var i = 0; i < fake.length; i++) {
         if (fake[i].url.test(url)) return success(fake[i].data)
      }
   }

   $.ajax({
      dataType: "json",
      url: BASE_URL + url,
      data: params,
      success: x => {
         if (x.no == -1 && window.vue) return window.vue.$router.push({ name: 'login', params: { callback: encodeURIComponent(location.href) } })
         if (x.no != 0) return toastr.error(x.data)
         success(x)
      },
      error: (jqXHR, textStatus, errorThrown) => {
         toastr.error(errorThrown)
         if (error) error(textStatus, errorThrown)
      }
   });
}

const post = (url, params, success) => {
   if (FAKE) {
      const fake = require("./fake.js")
      for (var i = 0; i < fake.length; i++) {
         if (fake[i].url.test(url)) return success(fake[i].data)
      }
   }

   $.ajax({
      method: "POST",
      url: BASE_URL + url,
      data: params,
      success: x => {
         if (x.no == -1 && window.vue) return window.vue.$router.push({ name: 'login', params: { callback: encodeURIComponent(location.href) } })
         if (x.no != 0) return toastr.error(x.data)
         success(x)
      },
      error: (jqXHR, textStatus, errorThrown) => {
         toastr.error(errorThrown)
         if (error) error(textStatus, errorThrown)
      }
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