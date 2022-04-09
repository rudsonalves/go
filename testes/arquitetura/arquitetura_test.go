package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependencia(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em amd64")
	}
	// ...
	t.Fail()
}
