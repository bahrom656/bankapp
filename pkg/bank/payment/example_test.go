package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 9,
		},
		{
			ID:     2,
			Amount: 4,
		},
		{
			ID:     3,
			Amount: 9,
		},
		{
			ID:     5,
			Amount: 5,
		},
	}

	max := Max(payments)
	fmt.Println(max.ID)
	//Output: 1
}
