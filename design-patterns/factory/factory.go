package factory

import (
	"fmt"
)

const (
	Cash       = 1
	DebitCard = 2
)

type PaymentMethod interface {
	Pay(amount float64) string
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m { 
        case Cash: 
        return new(CashPM), nil 
        case DebitCard: 
        return new(DebitCardPM), nil 
		default: 
		return nil, fmt.Errorf("payment method %d not recognized", m) 
	}
}

type CashPM struct {}
type DebitCardPM struct {}

func (c *CashPM) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float64) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}