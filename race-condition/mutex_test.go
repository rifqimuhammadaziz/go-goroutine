package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceConditionUsingMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	// the result will be same, because the counter only run in one goroutine (lock then unlock)
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock() // used if has shared variable (accessed by multiple goroutine) to avoid race condition
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x)
}

/**
CASE
If the struct will be accessed by multiple goroutine
*/
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

// WRITE PROCESS
func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock() // avoid race condition write data
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

// READ PROCESS
func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

// TEST
func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance:", account.GetBalance())
}

/**
CASE
Deadlock
*/
type UserBalance struct {
	sync.Mutex // Mutex sync.Mutex
	Name       string
	Balance    int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) ChangeBalance(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.ChangeBalance(-amount) // user1 send amount

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.ChangeBalance(amount)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Rifqi",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "Bagas",
		Balance: 1000000,
	}

	// DEADLOCK, 2 goroutine running at same time. the result does not match
	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)
	time.Sleep(5 * time.Second)

	fmt.Println("User", user1.Name, ", Balance", user1.Balance)
	fmt.Println("User", user2.Name, ", Balance", user2.Balance)

}
