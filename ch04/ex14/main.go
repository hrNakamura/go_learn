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
const BugFilter = "+label:bug"

type UserResults struct {
	TotalCount int            `json:"total_count"`
	Users      []*github.User `json:"items"`
}

const usersTmpl = `Users
{{.TotalCount}} issues:
{{range .Users}}----------------------------------------
User:	{{.Login}}
HTMLURL:{{.HTMLURL}}
{{end}}
`

type BugResults struct {
	TotalCount int
	Items      []*BugItems
}

type BugItems struct {
	Title     string
	HTMLURL   string `json:"html_url"`
	State     string
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

const bugTempl = `Bug Report
{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Title: {{.Title}}
URL:   {{.HTMLURL}}
State:  {{.State}}
CreateAt:    {{.CreatedAt | daysAgo}} days
{{end}}
`

//!+exec
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

var usersReport = template.Must(template.New("userlist").
	Parse(usersTmpl))

var bugReport = template.Must(template.New("buglist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(bugTempl))

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-exec

func handler(w http.ResponseWriter, r *http.Request) {
	termBase := r.URL.Path[1:]
	terms := strings.Split(termBase, "+")

	// result, err := github.SearchIssues(terms)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := report.Execute(w, result); err != nil {
	// 	log.Fatal(err)
	// }
	userResult, err := searchUsers(terms)
	if err != nil {
		log.Fatal(err)
	}
	if err := usersReport.Execute(w, userResult); err != nil {
		log.Fatal(err)
	}
	bugResult, err := searchBugReports(terms)
	if err != nil {
		log.Fatal(err)
	}
	if err := bugReport.Execute(w, bugResult); err != nil {
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

func searchBugReports(terms []string) (*BugResults, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(github.IssuesURL + "?q=" + q + BugFilter)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result BugResults
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
