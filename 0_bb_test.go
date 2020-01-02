package slices_test

import (
	"testing"

	"github.com/mark-ahn/slices"
)

func TestT(t *testing.T) {
	ss := slices.OfString([]string{
		"0", "1", "2", "3",
	})

	if "3" != ss.At(3) {
		t.Errorf("expect %v, got %v", "3", ss.At(3))
	}
}
