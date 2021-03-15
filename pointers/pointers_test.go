package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, w, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(20)}
		err := w.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, w, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(0)
		w := Wallet{startingBalance}
		got := w.Withdraw(Bitcoin(10))
		want := ErrInsufficientFunds

		assertError(t, got, want)
		assertBalance(t, w, startingBalance)
	})
}

func assertBalance(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got.Error() != want.Error() {
		t.Errorf("got %q want %q", got.Error(), want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
