package main

import (
	"fmt"
	"log"
	"os"

	"time"

	"gopl.io/ch4/github"
)

//!+
func main() {
	inMonthIssues := make(map[int]*github.Issue)
	inYearIssues := make(map[int]*github.Issue)
	beforeYearIssues := make(map[int]*github.Issue)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	beforeMonth := time.Now().AddDate(0, -1, 0)
	beforeYear := time.Now().AddDate(-1, 0, 0)
	for _, item := range result.Items {
		if beforeMonth.Before(item.CreatedAt) {
			inMonthIssues[item.Number] = item
		} else if beforeYear.Before(item.CreatedAt) {
			inYearIssues[item.Number] = item
		} else {
			beforeYearIssues[item.Number] = item
		}
	}
	fmt.Printf("\nIn a month issues:\n")
	for _, item := range inMonthIssues {
		fmt.Printf("%v\t#%-5d %9.9s %.55s\n", item.CreatedAt,
			item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("\nIn a year issues:\n")
	for _, item := range inYearIssues {
		fmt.Printf("%v\t#%-5d %9.9s %.55s\n", item.CreatedAt,
			item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("\nBefore a year issues:\n")
	for _, item := range beforeYearIssues {
		fmt.Printf("%v\t#%-5d %9.9s %.55s\n", item.CreatedAt,
			item.Number, item.User.Login, item.Title)
	}
}
