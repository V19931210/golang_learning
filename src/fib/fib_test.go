package fib

import (
	"testing"
)

func TestFib(t *testing.T) {
	//三种赋值方式

	//var a int = 1
	//var b int = 1

	//var (
	//	a int = 1
	//	b     = 1
	//)

	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

//赋值同时对多个变量赋值
func TestExchangeVal(t *testing.T) {
	a := 1
	b := 2
	t.Log(a, b)
	a, b = b, a
	t.Log(a, b)
}
