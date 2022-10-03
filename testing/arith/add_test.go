package arith

import "testing"

func Test_addNumbers(t *testing.T) {
	result := AddNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expectec 5, got", result)
	}
}
