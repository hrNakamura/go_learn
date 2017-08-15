package bank_test

import (
	"fmt"
	"testing"
	"time"

	bank "myProject/go_learn/ch09/ex01"
)

func TestWithdraw(t *testing.T) {
	done := make(chan struct{})

	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()
	var result bool
	go func() {
		time.Sleep(1 * time.Second)
		result = bank.Withdraw(100)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()
	<-done
	<-done
	if got, want := result, true; got != want {
		t.Errorf("Balance = %d, result = %v", bank.Balance(), got)
	}

	go func() {
		time.Sleep(1 * time.Second)
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()
	go func() {
		result = bank.Withdraw(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()
	<-done
	<-done
	if got1, want1, got2, want2 := result, false, bank.Balance(), 300; got1 != want1 || got2 != want2 {
		t.Errorf("result = %v, want %v", got1, false)
		t.Errorf("Balance = %d, want %d", got2, 300)
	}
	//預金をリセット
	bank.Withdraw(bank.Balance())
}

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
