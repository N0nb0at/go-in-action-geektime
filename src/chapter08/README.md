---
title: Go-08-条件与循环
date: 2019/06/09 16:39:25
categories: 
  - [Go-in-Action]
tags: 
  - [geektime]
  - [Golang]
---

## 循环

与其他主要语言差异

Go 语言仅支持循环关键字 `for`

```go
for i := 0; i <= 10; i++ 
```

## if 条件

与其他主要语言差异

1. condition 表达式结果必须为 boolean 值
2. 支持变量赋值：
```go
if var declaration; condition { ... }
```

## switch 条件

与其他主要语言差异

1. 条件表达式不限制为常量或整数
2. 单个 case 中，可以出现多个结果选项，使用逗号分隔
3. 与 C 语言等规则相反，Go 语言不需要用 break 来明确退出一个 case
4. 可以不设定 switch 之后的条件表达式，在此种情况下，整个 switch 结构与多个 if ... else ... 的逻辑作用等同
