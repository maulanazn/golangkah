package routine

import (
	"fmt"
	"sync"
	"testing"
)

type WalletAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (wallet *WalletAccount) AddBalance(amount int) {
	wallet.RWMutex.Lock()
	wallet.Balance = wallet.Balance + amount
	wallet.RWMutex.Unlock()
}

func (wallet *WalletAccount) GetBalance() int {
	wallet.RWMutex.RLock()
	balance := wallet.Balance
	wallet.RWMutex.RUnlock()
	return balance
}

func TestRwMutex(t *testing.T) {
	var wait sync.WaitGroup
	wallet := WalletAccount{}

	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			for j := 0; j < 10; j++ {
				wallet.AddBalance(1)
				fmt.Println(wallet.GetBalance())
			}
		}()
	}

	wait.Wait()
	fmt.Println("Total balance : ", wallet.GetBalance())
}

func SetName(name_channel chan<- interface{}, i int) {
	name_channel <- i
}

func GetName(name_channel <-chan interface{}) {
	data := <-name_channel
	fmt.Println(data)
}

func TestCallName(t *testing.T) {
	var wait sync.WaitGroup
	channel := make(chan interface{})
	defer close(channel)

	var x int = 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x += 1
				go SetName(channel, x)
				go GetName(channel)
				mutex.Unlock()
			}
		}()
	}

	wait.Wait()
	fmt.Println("Counter: ", x)
}

type GopayAccount struct {
	sync.Mutex
	sync.WaitGroup
	Name    string
	Balance int
}

func (gopayAccount *GopayAccount) Lock() {
	gopayAccount.Mutex.Lock()
}

func (gopayAccount *GopayAccount) Unlock() {
	gopayAccount.Mutex.Unlock()
}

func (gopayAccount *GopayAccount) Change(amount int) {
	gopayAccount.Balance = gopayAccount.Balance + amount
}

func Transfer(user1 *GopayAccount, user2 *GopayAccount, amount int) {
	defer user1.WaitGroup.Done()
	user1.Lock()
	user1.Change(-amount)

	defer user2.WaitGroup.Done()
	user2.Lock()
	user2.Change(amount)

	user1.Unlock()
	user2.Unlock()
}

func TestTransfer(t *testing.T) {
	user := GopayAccount{Name: "maulana", Balance: 10000}
	user1 := GopayAccount{Name: "fatih", Balance: 10000}

	user.WaitGroup.Add(1)
	user1.WaitGroup.Add(1)
	go Transfer(&user, &user1, 100)
	user.WaitGroup.Wait()

	fmt.Println("User : ", user.Name, "Balance", user.Balance)
	fmt.Println("User : ", user.Name, "Balance", user1.Balance)

}

func RunAsynchronously(group *sync.WaitGroup) {
	/**
	Done() must be on go routine function
	Add(int) must be outside of go routine function
	Wait() must be outside of go routine function and place it in the bottom after the function is executed
	*/
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello World")
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronously(group)
	}

	group.Wait()
	fmt.Println("Done")
}
