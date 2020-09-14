const express = require('express');

const srv = express()

const renderer = require('vue-server-renderer').createRenderer({
   template: require('fs').readFileSync('./ssr/layout.html', 'utf-8'),
});

const createApp = require('../dist/entry-server').default

srv.use("/", express.static('./dist'))
srv.get('*', (req, res) => {
   const context = {
      url: req.url,
      title: 'vue ssr',
      metas: `
   <meta charset="UTF-8">
   <meta name="keyword" content="vue,ssr">
   <meta name="description" content="vue srr demo">
   `,
   };

   createApp(context).then(app => {
      renderer.renderToString(app, context, (err, html) => {
         if (err) {
            console.log(err)
            if (err.code === 404) {
               res.status(404).end('Page not found')
            } else {
               res.status(500).end('Internal Server Error')
            }
         } else {
            res.end(html)
         }
      })
   }).catch(app => {
      console.log(app)
      res.status(404).end('Page not found')
   })
})

const port = process.env.PORT || 8080
srv.listen(port, () => {
   console.log(`server started at http://localhost:${port}`)
})