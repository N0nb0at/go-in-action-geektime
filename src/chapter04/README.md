---
title: Go-04-编写第一个 Go 程序
date: 2019/05/22 23:45:00
categories: 
  - [Go-in-Action]
tags: 
  - [geektime]
  - [Golang]
---


## 开发环境构建

`GOPATH`

- 1.8 版本以前必须设置环境变量
- 1.8 及之后版本，如果没有设置则使用默认值
    - Unix 上默认为 `$HOME/go`
    - Windows 上默认为 `%USERPROFILE%/go`
    - Mac 上可以通过 `~/.bash_profile` 设置

## 基本程序结构

```go
package main // 包，表明代码所在的模块（包）

import "fmt" // 引入代码依赖

// 功能实现
func main() {
   fmt.Println("Hello World!")
}
```

## 应用程序入口

- 必须是 main 包：`package main`
- 必须是 main 方法：`func main()`
- 文件名不一定是 `main.go`

## 退出返回值

- Go 中 main 函数不支持任何返回值
- 通过 os.Exit 来返回状态

## 获取命令行餐厨

- main 函数不支持传入参数
- 在程序中直接通过 os.Args 获取命令行参数