package types

//Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы, и т.д.)

type Money int64

//Currency представляет код валюты
type Currency string

//коды валюты
const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
)

//PAN представляет номер карты
type PAN string

//Card представляет информацию о платёжной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

//Payment представляет информацию о платеже.
type Payment struct {
	ID     int
	Amount Money
}
//PaymentSource представляет информацию о источнике оплате 
type PaymentSource struct{
	Type 		string
	Number 		string
	Balance 	Money
}