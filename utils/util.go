package utils

import (
	"bytes"
	"os"
)

var utf8bom = []byte{239, 187, 191}

// Exists is return true if the file exists
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// utf8のbomを消す
func StripBOM(in []byte) []byte {
	return bytes.TrimPrefix(in, utf8bom)
}
