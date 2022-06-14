package array

import "testing"

//数组 及其遍历
func TestArray(t *testing.T) {
	var a [3][3]int
	a[0][0] = 5
	b := [4]int{1, 2, 3, 4}
	c := [...]int{1, 3, 5, 7, 9}
	t.Log(a[0], a[1], a[2])
	t.Log(a[0:len(a)])
	t.Log(b[0], b[1], b[2])
	t.Log(c[0], c[1], c[2])
	for idx, e := range c {
		t.Log(idx, e)
	}
	for _, e := range c {
		t.Log(e)
	}
}
