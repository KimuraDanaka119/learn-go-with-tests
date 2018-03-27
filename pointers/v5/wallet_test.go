package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw over balance limit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		if err == nil {
			t.Errorf("expected an error to be returned when withdrawing too much")
		}

		if got, isWithdrawErr := err.(WithdrawError); isWithdrawErr {
			want := WithdrawError{
				AmountToWithdraw: Bitcoin(100),
				CurrentBalance:   Bitcoin(20),
			}
			if want != got {
				t.Errorf("got %#v, want %#v", got, want)
			}
		} else {
			t.Errorf("did not get a withdraw error %#v", err)
		}
	})

}
