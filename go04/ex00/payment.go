package main

import (
	"errors"
	"fmt"
)

type Payer interface{ Pay(float64) error }

type CreditCard struct {
	Number  string
	Holder  string
	Limit   float64
	Balance float64
}

type BankTransfer struct {
	IBAN  string
	Owner string
}

type Crypto struct {
	Wallet   string
	Currency string
	Rate     float64
}

func (c *CreditCard) Pay(amount float64) error {
	if amount > c.Limit {
		return errors.New("payment goes beyond limit")
	}
	if amount > c.Balance {
		return errors.New("not enough balance in the credit card")
	}
	c.Balance -= amount
	return nil
}

func (b *BankTransfer) Pay(amount float64) error {
	fmt.Println(b.String(), " succeeded payment")
	return nil
}

func (b *BankTransfer) String() string {
	return fmt.Sprintf("Owner: %s. IBAN: %s.", b.Owner, b.IBAN)
}

func (c *Crypto) Pay(amount float64) error {
	return nil
}

func ProcessPayment(p Payer, amount float64) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println(err.Error())
	}
}
