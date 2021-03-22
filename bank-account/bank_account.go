package account

type Account struct {
	balance int64
	open    bool
}

// API:
//
// Open(initialDeposit int64) *Account
// (*Account) Close() (payout int64, ok bool)
// (*Account) Balance() (balance int64, ok bool)
// (*Account) Deposit(amount int64) (newBalance int64, ok bool)

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{initialDeposit, true}
}

func (a *Account) Close() (payout int64, ok bool) {
	if a.open {
		a.open = false
		balance := a.balance
		a.balance = 0
		return balance, true
	} else {
		return 0, false
	}
}

func (a *Account) Balance() (balance int64, ok bool) {
	if a.open {
		return a.balance, true
	} else {
		return 0, false
	}
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if !a.open {
		return 0, false
	} else if -amount > a.balance {
		return 0, false
	} else {
		a.balance += amount
		return a.balance, true
	}
}
