package constant

import "testing"

const (
	Mon = iota + 1
	Tue
	Wed
)

func TestConst(t *testing.T) {
	t.Log(Mon, Tue)
}
