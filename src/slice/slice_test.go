package slice

import "testing"

func TestSlice(t *testing.T) {
	//切片的声明及初始化方式
	//var s0 []int
	s0 := []int{}
	t.Log(len(s0), cap(s0))

	s1 := []int{2, 4, 6, 8}
	t.Log(len(s1), cap(s1))

	//make(type, len, cap)
	s2 := make([]int, 3, 5)
	s2[0] = 11
	t.Log(len(s2), cap(s2))
}

func TestSliceShareMemory(t *testing.T) {
	s0 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s0[3:6]
	t.Log(s1, len(s1), cap(s1))
	s2 := s0[5:8]
	t.Log(s2, len(s2), cap(s2))
	s2[0] = 100
	t.Log(s1)
}
