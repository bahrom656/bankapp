package payment

import "bank/pkg/bank/types"

func Max(payments []types.Payment) types.Payment {
	max := payments[0]
	for _, payment := range payments {
		if max.Amount < payment.Amount {
			max.Amount = payment.Amount
			max.ID = payment.ID
		}
	}
	return max
}
