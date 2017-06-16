package main

import "fmt"

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("args empty")
	}
	val := vals[0]
	for i := 1; i < len(vals); i++ {
		if val < vals[i] {
			val = vals[i]
		}
	}
	return val, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("args empty")
	}
	val := vals[0]
	for i := 1; i < len(vals); i++ {
		if val > vals[i] {
			val = vals[i]
		}
	}
	return val, nil
}

func max2(first int, vals ...int) int {
	val := first
	for i := 0; i < len(vals); i++ {
		if val < vals[i] {
			val = vals[i]
		}
	}
	return val
}

func min2(first int, vals ...int) int {
	val := first
	for i := 0; i < len(vals); i++ {
		if val > vals[i] {
			val = vals[i]
		}
	}
	return val
}

func main() {
	//!+main
	fmt.Println(max())
	fmt.Println(min())
	// fmt.Println(max2()) //error
	// fmt.Println(min2()) //error
	fmt.Println(max(3))
	fmt.Println(min(3))
	fmt.Println(max2(3))
	fmt.Println(min2(3))
	fmt.Println(max(1, 2, 3, 4))
	fmt.Println(min(1, 2, 3, 4))
	fmt.Println(max2(1, 2, 3, 4))
	fmt.Println(min2(1, 2, 3, 4))
	//!-main

}
