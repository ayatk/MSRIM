package utils

import (
	"reflect"
	"testing"
)

type existsTest struct {
	path    string
	boolean bool
}

func TestExists(t *testing.T) {
	var existsTests = []existsTest{
		{"../main.go", true},
		{"../scripts/compile", true},
		{"../test_assets/test.jar", true},
		{"../unknown-dir", false},
		{"../unknown-file.txt", false},
		{"../unknown-dir/unknown-file.txt", false},
	}

	for i := range existsTests {
		test := &existsTests[i]
		if Exists(test.path) != test.boolean {
			t.Errorf("Failed chack exists: %v => %t wants %t", test.path, Exists(test.path), test.boolean)
		}
	}
}

func TestStripBOM(t *testing.T) {
	var bom = [][]byte{
		{239, 187, 191, 244, 145, 128, 216, 130, 239, 116},
		{105, 155, 243, 209, 202, 129, 156, 114, 236, 138},
		{149, 183, 173, 138, 182, 134, 197, 127, 120, 100},
		{239, 187, 191, 255, 194, 159, 105, 195, 241, 167},
	}

	var answer = [][]byte{
		{244, 145, 128, 216, 130, 239, 116},
		{105, 155, 243, 209, 202, 129, 156, 114, 236, 138},
		{149, 183, 173, 138, 182, 134, 197, 127, 120, 100},
		{255, 194, 159, 105, 195, 241, 167},
	}

	for i := range bom {
		st := bom[i]
		ans := answer[i]
		if !reflect.DeepEqual(StripBOM(st), ans) {
			t.Errorf("Failed remove BOM")
		}
	}
}
