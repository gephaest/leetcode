package p13_roman_to_integer

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	input := "MCMXCIV" // 44
	//input := "CM" // 44

	//input := "XLIV" // 44
	//input := "CXL"

	res := romanToInt(input)

	fmt.Printf("input: %s\n"+
		"resust: %d\n", input, res)
}
