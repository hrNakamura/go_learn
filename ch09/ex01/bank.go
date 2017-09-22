// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrawal = make(chan int)
var results = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

//TODO 複数人が同時にWithDrawをCallするとresultがどちらのものか保証されない
//TODO 構造体に引き出し金額と結果をまとめたチャネルを関数内で生成する
func Withdraw(amount int) bool { withdrawal <- amount; return <-results }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdrawal:
			if balance < amount {
				results <- false
			} else {
				balance -= amount
				results <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
