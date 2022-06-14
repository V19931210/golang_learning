package slice

import "testing"

func TestSlice(t *testing.T) {
	//切片的声明及初始化方式
	var s0 []int
	t.Log(len(s0), cap(s0))

	s1 := []int{2, 4, 6, 8}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	s2[0] = 11
	t.Log(len(s2), cap(s2))
}
