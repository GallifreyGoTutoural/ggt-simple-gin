# ggt-simple-gin
[English](./README.md) || [中文](./README_zh.md)

ggt-simple-gin是一个模仿gin手写的Web框架项目，旨在更深入地了解gin的底层原理和设计模式，最终实现一个简易版本的Web框架——gsg 。

项目的名称中的 "ggt" 是 "Gallifrey's GoTutoural" 的简写，而 "gsg" 则是 "ggt-simple-gin" 的简写。

项目的主要参考来源是是[极客兔兔](https://geektutu.com/)大佬的博客：[7天用Go从零实现Web框架Gee教程](https://geektutu.com/post/gee.html)，如果想了解更多程序设计细节和考量，请查阅原博客。

## 开发计划

- [x] 构造`Engine`结构体实现http包的`ServeHTTP`接口，添加其构造器
- [x] `Engine`封装GET和POST请求的函数，实现基本的网络请求和响应功能
- [x] 抽离出`router`，方便后续功能开发
- [x] 设计`Context`，封装Request和Response ，提供对 JSON、HTML 等返回类型的支持
- [x] 使用Trie树实现动态路由解析
- [x] 添加两种模式匹配支持`:name`和`*filepath`
- [x] 实现路由分组控制
- [x] 设计并实现Web框架的中间件机制
- [x] 实现通用的`Logger`中间件，实现`Logger`能够记录请求到响应所花费的时间的功能
- [ ] 实现静态资源服务
- [ ] 支持HTML模板渲染
- [ ] 实现错误处理机制



