package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		t.Helper()
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		got := w.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
