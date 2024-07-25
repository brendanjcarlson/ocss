package testutils

import (
	"os"
	"testing"
)

func MustReadFile(t *testing.T, filename string) []byte {
	contents, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return contents
}
