// 异步请求

const BASE_URL = 'http://127.0.0.1:8080'

export const getJSON = (url, params, callback) => {
   const fake = require("./fake.js")
   for (var i = 0; i < fake.length; i++) {
      if (fake[i].url.test(url)) return callback(fake[i].data)
   }

   $.getJSON(BASE_URL + url, params, x => {
      if (x.no == -1) return location.href = "./login"
      callback(x)
   });
}

export const post = (url, params, callback) => {
   const fake = require("./fake.js")
   for (var i = 0; i < fake.length; i++) {
      if (fake[i].url.test(url)) return callback(fake[i].data)
   }

   $.post(BASE_URL + url, params, x => {
      if (x.no != 0) return alert(x.data)
      callback(x)
   });
}