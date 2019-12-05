# 网站脚手架

> 方便快速的构建一个网站。Bootstrap3 + AdminLTE + Golang + Vue

- 前端展示使用Golang template
- 后端管理模式1:使用Golang template
- 后端管理模式2:使用Vue前后端分离开发

## 目录结构
```
- root
 ∟ api        程序接口/第三方服务
 ∟ backend    后端自动化执行代码
 ∟ controller 业务逻辑Golang代码
 ∟ db         sqlite3数据库
 ∟ dist       前端开发的零时输出文件
 ∟ model      数据库操作Golang代码
 ∟ public     公用静态资源
 ∟ util       公用Golang方法
 ∟ vendor     第三方Golang依赖库
 ∟ view       模式1的前端html模版
 ∟ vue        模式2的前端分离开发vue源文件
```

## 克隆
```
git clone https://github.com/ohko/tpler.git
```

## Golang开发
```
fswatch
# 或
DEBUG=1 go run .
```

## 前端vue开发
```
parcel vue/index.html
```

## 前端编译vue
```
parcel build --public-url=/public/admin/ --out-dir=./public/admin/ --no-source-maps --no-cache --no-minify vue/index.html
```