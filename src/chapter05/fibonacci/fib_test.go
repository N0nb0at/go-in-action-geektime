package fibonacci

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var foo = 1
	var bar = 1
	t.Log(foo)
	for i := 0; i < 10; i++ {
		t.Log(" ", bar)
		tmp := foo
		foo = bar
		bar = tmp + foo
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2

	tmp := a
	a = b
	b = tmp

	t.Log(a, b)

	a, b = b, a

	t.Log(a, b)
}
