package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got == nil {
			t.Fatal("Wanted an error but didn't got one")
		}

		if got.Error() != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("balance", func(t *testing.T) {
		assertBalance(t, Wallet{balance: Bitcoin(10)}, Bitcoin(10))
	})

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient ammount", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		errorString := fmt.Sprintf("only %s available. Can't withdraw %s", wallet.Balance(), Bitcoin(100))

		assertError(t, err, errorString)
		assertBalance(t, wallet, startingBalance)
	})
}
