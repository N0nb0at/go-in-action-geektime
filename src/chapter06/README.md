---
title: Go-06-数据类型
date: 2019/06/09 13:40:40
categories: 
  - [Go-in-Action]
tags: 
  - [geektime]
  - [Golang]
---

## 基本数据类型

```
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32, represents a Unitcode code point
float32 float64
complex64 complex128
```

## 类型转换

与其他主要语言的差异：

1. Go 语言不允许隐式类型转换
2. 别名和原有类型也不能进行隐式类型转换

## 类型的预定义值

1. math.MaxInt64
2. math.MaxFloat64
3. math.MaxUint32

## 指针类型

与其他主要语言的差异：

1. 不支持指针运算
2. string 是值类型，器默认的初始化值为空字符串，而不是 nil
