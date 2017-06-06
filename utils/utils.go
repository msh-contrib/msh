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

// InArray checks if certain value is in list of items
func InArray(value interface{}, list []interface{}) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

// GetIndex find index of element in slice
func GetIndex(element interface{}, list []interface{}) int {
	for pos, item := range list {
		if item == element {
			return pos
		}
	}

	return -1
}
