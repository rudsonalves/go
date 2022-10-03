package arith

import "testing"

func Test_SubNumbers(t *testing.T) {
	r := SubNumbers(5, 3)
	if r != 2 {
		t.Error("incorrect result: expectec 2, got", r)
	}
}
