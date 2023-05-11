//@Author: Dennis
//@Date: 2023/5/11

package wallet

import (
	"fmt"
	"os"
	"testing"
)

func TestWithdraw(t *testing.T) {
	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{
			Balance: 100,
		}
		err := wallet.Withdraw(90)
		if err != nil {
			t.Errorf("balance error with %v", wallet)
		}
		if wallet.Balance != 10 {
			t.Errorf("balance error with %v", wallet)
		}
	})
	t.Run("test withdraw not enough balance", func(t *testing.T) {
		wallet := Wallet{Balance: 10}
		err := wallet.Withdraw(90)
		if err == NotEnoughAmount {
			_, _ = fmt.Fprintf(os.Stdout, "withdraw error:%v\n", err)
		}
		if wallet.Balance < 0 {
			t.Errorf("withdraw error")
		}
	})
}
