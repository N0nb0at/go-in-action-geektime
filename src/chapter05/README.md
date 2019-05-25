---
title: Go-05-变量、常量以及其他语言的差异
date: 2019/05/25 16:05:44
categories: 
  - [Go-in-Action]
tags: 
  - [geektime]
  - [Golang]
---

## 大师与初学者的差异

> The master has failed more times than the  beginner has tried.

## 编写测试程序

- 源码文件以 `_test` 结尾： `xxx_test.go`
- 测试方法名以 `Test` 开头： `func TestXXX(t *testing.T) { ... }`

Example: [first_test.go](test/first_test.go)

## 实现 Fibonacci 数列

Example: [fibonacci](fibonacci/fib_test.go)

## 变量赋值

与其他主要编程语言的差异：

- 赋值可以进行自动类型推断
- 在一个赋值语句中可以对多个变量进行同时赋值

## 常量定义

与其他主要编程语言的差异

快速设置连续值：

```go
const (
	Mondaky = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
```

```go
const(
    Open = 1 << iota
    Close
    Pending
)
```