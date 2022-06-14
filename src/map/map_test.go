package _map

import "testing"

func TestMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Logf("len m1=%d", len(m1))
	m2 := map[string]int{}
	t.Logf("len m2=%d", len(m2))
	//make(map,cap)
	m3 := make(map[string]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestMapTravel(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	//if判断条件的形式
	if value, ok := m1[3]; ok {
		t.Logf("key is existed, its value is:%d", value)
	} else {
		t.Log("key is not existed")
	}
	for key, value := range m1 {
		t.Log(key, value)
	}
}
