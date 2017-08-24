// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"fmt"
	"sync"
	"testing"

	memo "github.com/hrNakamura/go_learn/ch09/ex03"
)

func TestMemo(t *testing.T) {
	cancel := fmt.Errorf("cancelled")
	fin := make(chan string)
	f := func(key string, done chan struct{}) (interface{}, error) {
		select {
		case <-done:
			return nil, cancel
		case <-fin:
			return key, nil
		}
	}

	keys := []string{"alpha", "beta", "gamma"}
	m := memo.New(memo.Func(f))
	defer m.Close()
	wg := &sync.WaitGroup{}

	fmt.Println("test cancel")
	wg.Add(len(keys))
	for _, key := range keys {
		done := make(chan struct{})
		go func(key string) {
			v, err := m.Get(key, done)
			wg.Done()
			if v != nil || err != cancel {
				t.Errorf("key=%v :got %v, %v\n", key, v, err)
			}
		}(key)
		close(done)
	}
	wg.Wait()

	fmt.Println("test normal")
	wg.Add(len(keys))
	for _, key := range keys {
		done := make(chan struct{})
		go func(key string) {
			v, err := m.Get(key, done)
			wg.Done()
			if v != key || err != nil {
				t.Errorf("key=%v :got %v, %v\n", key, v, err)
			}
		}(key)
		fin <- "ok"
	}
	wg.Wait()
}
