package condition

import "testing"

func TestXXX(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("偶数")
		case 1, 3:
			t.Log("奇数")
		default:
			t.Log("o_o")
		}
	}
}

func TestXXX2(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("偶数")
		case i%2 == 1:
			t.Log("奇数")
		default:
			t.Log("o_o")
		}
	}
}
