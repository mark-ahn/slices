package slices_test

import (
	"fmt"
	"testing"

	"github.com/mark-ahn/slices/v2"
)

func TestRange(t *testing.T) {
	list := slices.Range[uint16](20, 51, 3)
	fmt.Println(list)
}
