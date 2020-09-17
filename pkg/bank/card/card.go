package card

import (
	"bank/pkg/bank/types"
)

const (
	withDrawLimit = 50_000_00
	bonusLimit    = 5_000_00
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}

func Deposit(card *types.Card, amount types.Money) {

	if amount < 0 {
		return
	}
	if !(*card).Active {
		return
	}
	if amount > withDrawLimit {
		return
	}
	(*card).Balance = amount + (*card).Balance
}
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	bonus := int(card.MinBalance) * percent * daysInMonth / daysInYear / 100
	if !card.Active {
		return
	}
	if bonus > bonusLimit {
		return
	}
	if (*card).Balance <= 0 {
		return
	}
	(*card).Balance = (*card).Balance + types.Money(bonus)
}

func Withdraw(card types.Card, amount types.Money) types.Card {

	if !card.Active {
		return card
	}
	if card.Balance < amount {
		return card
	}
	if amount > 20_000_00 {
		return card
	}
	if amount < 0 {
		return card
	}
	card.Balance = card.Balance - amount
	return card
}
func Total(cards []types.Card) types.Money {
	var sum types.Money
	for _, card := range cards {
		if card.Active == true && card.Balance >= 0{
			sum = sum + card.Balance
		}
	}
	return sum
}
func PaymentSources(cards []types.Card) []types.PaymentSource{
	var source []types.PaymentSource
	var indexSource int
	source = make([]types.PaymentSource, len(cards))
	for index, card := range cards {
		if card.Active == true && card.Balance > 0 {
			source[indexSource].Number = string(cards[index].PAN)
			source[indexSource].Balance = cards[index].Balance
			source[indexSource].Type = "card"
			indexSource++ 
		}
	}
	return source
}