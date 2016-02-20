package utils

import "os"

// Exists is return true if the file exists
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
