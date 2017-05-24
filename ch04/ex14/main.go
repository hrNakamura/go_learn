// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 113.

// Issuesreport prints a report of issues matching the search terms.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"

	"gopl.io/ch4/github"
)

//!+template
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

//!-template

//!+daysAgo
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

//!-daysAgo

const UserURL = "https://api.github.com/search/users"

type UserResults struct {
	TotalCount int            `json:"total_count"`
	Users      []*github.User `json:"items"`
}

const usersTmpl = `{{.TotalCount}} issues:
{{range .Users}}----------------------------------------
User:	{{.Login}}
HTMLURL:{{.HTMLURL}}
{{end}}`

//!+exec
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

var usersReport = template.Must(template.New("userlist").Parse(usersTmpl))

func main() {
	// result, err := github.SearchIssues(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := report.Execute(os.Stdout, result); err != nil {
	// 	log.Fatal(err)
	// }

	userResult, err := searchUsers(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := usersReport.Execute(os.Stdout, userResult); err != nil {
		log.Fatal(err)
	}
}

//!-exec

func noMust() {
	//!+parse
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	//!-parse
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func searchUsers(terms []string) (*UserResults, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(UserURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result UserResults
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
