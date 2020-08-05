package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(filename string)(string, error) {
	if IsEmpty(filename) {
		return "", errors.New("Boş veri dosya adı olarak kullanılamaz.")
	}

	bytes, err := ioutil.ReadFile(filename)
	checkError(err)
	return string(bytes),nil

}
