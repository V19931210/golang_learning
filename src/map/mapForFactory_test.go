package _map

import "testing"

//map的value作为函数 实现工厂模式
func TestMapForFactory(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op << 1 }
	m[3] = func(op int) int { return op << 2 }
	t.Log(m[1](2), m[2](2), m[3](2))
}

//map[type]bool实现set
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Log("key is existed")
	} else {
		t.Log("key is not existed")
	}
	delete(mySet, 1)
	if mySet[n] {
		t.Log("key is existed")
	} else {
		t.Log("key is not existed")
	}
}
