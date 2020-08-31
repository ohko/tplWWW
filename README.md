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
 ∟ component  组件
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
# 预处理静态文件
npm run postbuild
# 启动127.0.0.1:1234测试
parcel vue/index.html
# 或 配合后端服务动态监听文件变化
parcel watch --public-url=/adm/ --out-dir=./dist/ vue/index.html
```

## 前端编译vue
```
parcel build --public-url=/public/admin/ --out-dir=./public/admin/ --no-source-maps --no-cache --no-minify vue/index.html
```

## 后端日志流程跟踪
```golang
// 1. 定义日志写到文件
LLFile = logger.NewDefaultWriter(&logger.DefaultWriterOption{
   Clone: os.Stdout,
   Path:  "./log",
   Label: "tpler",
   Name:  "log_",
})

// 2. 执行流程中创建新的logger对象，在函数间作为参数或上下文参数传递
// 具体实例代码：/backend/demo2/demo2.go
func (o *Demo2) do() {
	for {
		time.Sleep(time.Second * 5)

      // 创建新的logger对象
      loger := logger.NewLogger(common.LLFile)
      // 用nano时间格式保证唯一性和了解启动时间点
      loger.SetPrefix(time.Now().Format(time.RFC3339Nano))
      // 传递logger
		o.step1(loger)
	}
}

// loger当作当前协程日志记录接口
func (o *Demo2) step1(loger *logger.Logger) {
   loger.Log0Debug("Demo2 :: step1")
   // 传递logger
	go o.step2(loger)
}

// loger当作当前协程日志记录接口
func (o *Demo2) step2(loger *logger.Logger) {
   loger.Log0Debug("Demo2 :: step2")
   // 传递logger
	go o.step3(loger)
}

// loger当作当前协程日志记录接口
func (o *Demo2) step3(loger *logger.Logger) {
	go loger.Log0Debug("Demo2 :: step3")
}
```

## 日志效果
```log
2020/02/12 16:52:38 demo2.go:44: [2020-02-12T16:52:38.78484+08:00:D] Demo2 :: step1
2020/02/12 16:52:38 demo2.go:44: [2020-02-12T16:52:38.784881+08:00:D] Demo2 :: step1
2020/02/12 16:52:38 demo2.go:49: [2020-02-12T16:52:38.784881+08:00:D] Demo2 :: step2
2020/02/12 16:52:38 demo2.go:55: [2020-02-12T16:52:38.784881+08:00:D] Demo2 :: step3
2020/02/12 16:52:38 demo2.go:49: [2020-02-12T16:52:38.78484+08:00:D] Demo2 :: step2
2020/02/12 16:52:38 demo2.go:55: [2020-02-12T16:52:38.78484+08:00:D] Demo2 :: step3
2020/02/12 16:52:43 demo2.go:44: [2020-02-12T16:52:43.789002+08:00:D] Demo2 :: step1
2020/02/12 16:52:43 demo2.go:44: [2020-02-12T16:52:43.789017+08:00:D] Demo2 :: step1 <- 789017:可讲上下日志串起来
2020/02/12 16:52:43 demo2.go:49: [2020-02-12T16:52:43.789017+08:00:D] Demo2 :: step2 <- 789017:可讲上下日志串起来
2020/02/12 16:52:43 demo2.go:49: [2020-02-12T16:52:43.789002+08:00:D] Demo2 :: step2
2020/02/12 16:52:43 demo2.go:55: [2020-02-12T16:52:43.789002+08:00:D] Demo2 :: step3
2020/02/12 16:52:43 demo2.go:55: [2020-02-12T16:52:43.789017+08:00:D] Demo2 :: step3 <- 789017:可讲上下日志串起来
```