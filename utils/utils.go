package utils

import (
	"io/ioutil"
	"log"
)

// ReadFile allows to read file from file path string
func ReadFile(file string) (data []byte) {
	data, readError := ioutil.ReadFile(file)
	if readError != nil {
		log.Fatal(readError)
	}

	return
}
