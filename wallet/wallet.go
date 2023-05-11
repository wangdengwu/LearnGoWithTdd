//@Author: Dennis
//@Date: 2023/5/11

package wallet

var NotEnoughAmount = NotEnoughAmountError("not enough money")

type NotEnoughAmountError string

func (e NotEnoughAmountError) Error() string {
	return string(e)
}

type Wallet struct {
	Balance int
}

func (w *Wallet) Withdraw(amount int) error {
	if w.Balance < amount {
		return NotEnoughAmount
	}
	w.Balance -= amount
	return nil
}
