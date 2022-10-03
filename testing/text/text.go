package text

import (
	"io/ioutil"
	"unicode/utf8"
)

func CountCharacteres(fileName string) (int, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return 0, err
	}
	return utf8.RuneCount(data), nil
}
