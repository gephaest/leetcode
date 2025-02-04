package p122_buy_sell_stock

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	in := []int{7, 6, 4, 3, 1}

	fmt.Printf("in: %v\n", in)

	res := maxProfit(in)

	fmt.Printf("res: %d\n", res)
}
