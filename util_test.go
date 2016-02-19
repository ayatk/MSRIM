package main

import "testing"

type existsTest struct {
	path    string
	boolean bool
}

var existsTests = []existsTest{
	{"main.go", true},
	{"images/image.go", true},
	{"test_assets/test.jar", true},
	{"unknown-dir", false},
	{"unknown-file.txt", false},
	{"unknown-dir/unknown-file.txt", false},
}

func TestExists(t *testing.T) {
	for i := range existsTests {
		test := &existsTests[i]
		if Exists(test.path) != test.boolean {
			t.Errorf("Failed chack exists: %v => %t wants %t",
				test.path, Exists(test.path), test.boolean)
		}
	}
}
