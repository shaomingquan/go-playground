package oop

import "fmt"

type A interface {
	foo () int
	bar () int
}

type B struct {
	x int
}

func (this *B) foo () int {
	this.x = 1
	return this.x
}

func (this *B) bar () int {
	return this.x
}

func Struct1 () {
	var a A

	b := B{}

	a = &b
	// a = b // 报错

	fmt.Println(a.foo(), b.bar()) // 有副作用
}

type B2 struct {
	x int
}

func (this B2) foo () int {
	this.x = 1
	return this.x
}

func (this B2) bar () int {
	return this.x
}

func Struct2 () {
	var a A

	b := B2{}

	// a = &b // 不报错，虽然不报错，但也别这样写。并且同样是无副作用的
	a = b

	fmt.Println(a.foo(), b.bar()) // 无副作用
}

/**

为什么同样a的声明方式，Struct1需要指针，Struct2都可以？
或许可以这样理解
将b赋值给a的时候，原本是个空壳的a就有了调用实体值
因为这个值本身需要作为参数调用（语法糖），需要符合定义
在B中，这个值是一个指针类型，所以给a的也应该是指针

至于为什么Struct2不报错？语法糖吧，之所以可以这么设计
因为如果底层拿到指针，当然可以做值copy这样的事
反之如果底层拿到的就死值copy，它无法做指针相关的事

**/