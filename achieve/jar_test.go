package achieve

import (
	"os"
	"testing"
)

func TestUnJar(t *testing.T) {
	jarSrc := "../test_assets/test.jar"
	jarDest := "../test_assets/dest"

	if err := UnJar(jarSrc, jarDest); err != nil {
		t.Fatalf("Un zipping error: %v", err)
	}
	os.RemoveAll(jarDest)
}
