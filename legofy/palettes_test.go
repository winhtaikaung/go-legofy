package legofy

import (
	"testing"
)

func extend_palette_test(t *testing.T) {
	numbers := make([]int, 0)
	numbers = append(numbers, 0)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	if len(ExtendPalette(numbers, 256, 3)) != 768 {
		t.Log("Should match with Zawgyi string")
		t.Fail()
	}

}
