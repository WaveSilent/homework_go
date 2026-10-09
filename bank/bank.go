package bank

import (
	"errors"
	"fmt"
)

type Account struct {
	bill int
}

func NewAccount(bill int) Account {
	return Account{
		bill: bill,
	}
}

func (a *Account) ExtraditionMoney(sum int) (int, error) {
	if a.bill-sum < 0 {
		return 0, errors.New("Недостаточно средств")
	}

	a.bill -= sum

	return sum, nil
}

func (a Account) CheckBalance() {
	fmt.Println("Текущий баланс:", a.bill)
}

func (a *Account) OnlinePay(sum int) error {
	if a.bill-sum < 0 {
		return errors.New("Недостаточно средств")
	}

	fmt.Println("Покупка успешно сделана")
	fmt.Println("Со счета снято", sum, "тыс. рублей")

	a.bill -= sum

	return nil
}
