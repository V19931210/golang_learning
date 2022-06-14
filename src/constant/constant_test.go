package constant

import "testing"

const (
	Mon = iota + 1
	Tue
	Wed
)

const (
	open = 1 << iota
	close
	pending
)

func TestConst(t *testing.T) {
	t.Log(Mon, Tue)
	t.Log(open, close)
}

func TestBitOperation(t *testing.T) {
	a := 7
	t.Log(a&open == open, a&close == close, a&pending == pending)
	a &^= open //等价于 a=a&^open  &^为按位清零运算符
	t.Log(a&open == open, a&close == close, a&pending == pending)
}
