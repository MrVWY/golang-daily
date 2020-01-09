package main

import "fmt"

//指针 指针方法
//非指针 非指针方法

func modify(num int) {
	num = 1000
}

func modify1(num *int) {
	*num = 1000
}

type testint int

//乘2
func (p *testint) testdouble() int {
	*p = *p * 2
	fmt.Println("testdouble p = ", *p)
	return 0
}

//平方
func (p testint) testsquare() int {
	p = p * p
	fmt.Println("square p = ", p)
	return 0
}

type Foo struct {
	    val int
	}

func (f Foo) Inc(inc int) {
	f.val += inc
}

type Foo2 struct {
	val int
}

func (f *Foo2) Inc(inc int) {
	f.val += inc
}

func main() {
	a := 1
	modify(a)
	fmt.Println(a)

	modify1(&a)
	fmt.Println(a)

	var i testint = 2
	i.testdouble()
	fmt.Println("i = ", i)
	i.testsquare()
	fmt.Println("i = ", i)

	var f Foo
	f.Inc(100)
	fmt.Println(f.val)

	var f2 Foo2
	f2.Inc(100)
	fmt.Println(f2.val)
}