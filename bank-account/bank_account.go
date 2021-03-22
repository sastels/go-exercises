package account

import "sync"

type Account struct {
	mutex   sync.Mutex
	balance int64
	open    bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{sync.Mutex{}, initialDeposit, true}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()

	if a.open {
		a.open = false
		balance := a.balance
		a.balance = 0
		a.mutex.Unlock()
		return balance, true
	} else {
		a.mutex.Unlock()
		return 0, false
	}
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	if a.open {
		a.mutex.Unlock()
		return a.balance, true
	} else {
		a.mutex.Unlock()
		return 0, false
	}
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	if !a.open {
		a.mutex.Unlock()
		return 0, false
	} else if -amount > a.balance {
		a.mutex.Unlock()
		return 0, false
	} else {
		a.balance += amount
		a.mutex.Unlock()
		return a.balance, true
	}
}
