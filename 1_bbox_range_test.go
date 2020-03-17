package slices_test

import (
	"fmt"
	"testing"

	"github.com/mark-ahn/slices"
)

func TestRange(t *testing.T) {
	list := slices.Uint16T.Range(20, 50)
	fmt.Println(list)
}
