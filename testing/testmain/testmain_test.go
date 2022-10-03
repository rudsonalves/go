package arithmetic_test

import (
	"fmt"
	"os"
	"rudsonalves/testing/testmain/arithmetic"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("set up stuff for tests here")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exitVal)
}

func Test_AddNumbers(t *testing.T) {
	r := arithmetic.AddNumbers(2, 3)
	e := 2 + 3
	fmt.Println("AddNumbers:", testTime)
	if r != e {
		t.Errorf("incorrect result: expectec %d, got %d\nTime: %v", e, r, testTime)
	}
}

func Test_SubNumbers(t *testing.T) {
	r := arithmetic.SubNumbers(5, 3)
	e := 5 - 3
	fmt.Println("SubNumbers:", testTime)
	if r != e {
		t.Errorf("incorrect result: expectec %d, got %d\n%v", e, r, testTime)
	}
}
