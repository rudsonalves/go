package text

import "testing"

func Test_CountCharacteres(t *testing.T) {
	total, err := CountCharacteres("test_data/sample1.txt")
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if total != 35 {
		t.Error("expected 35, got", total)
	}

	_, err = CountCharacteres("test_data/no_file.txt")
	if err == nil {
		t.Error("expected an error")
	}
}
