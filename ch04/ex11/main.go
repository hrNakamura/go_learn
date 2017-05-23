package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type singleIssue struct {
	Number  int
	HTMLURL string `json:"html_url"`
	State   string
	Title   string
	Body    string
}

type createIssue struct {
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Assignee string   `json:"assignee,omitempty"`
	Labels   []string `json:"labels,omitempty"`
}

// APIURL Github api
const APIURL = "https://api.github.com"

// GitHubToken access token
const GitHubToken = "e4fc10d742090777175b340b00f13b5bcf9a67fe"

func main() {
	createNewIssue()
	data, err := getSingleIssue("hrNakamura", "go_learn", 1)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	mar, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", mar)
}

func getSingleIssue(owner, repos string, number int) (*singleIssue, error) {
	q := "/repos/" + owner + "/" + repos + "/issues/" + strconv.Itoa(number)
	resp, err := http.Get(APIURL + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get query failed: %s", resp.Status)
	}
	var result singleIssue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func createNewIssue() {
	url := APIURL + "/repos/hrNakamura/go_learn/issues"
	var st createIssue
	st.Title = "new issue"
	st.Body = "test"
	st.Assignee = "hrNakamura"
	// st.Labels = append(st.Labels, "bug")
	jsonStr, err := json.Marshal(st)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Println(string(jsonStr))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	token := "token " + GitHubToken
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("%s", resp.Status)
	}
	fmt.Printf("%s\n", resp.Status)
}

func updateIssue() {

}

func closeIssue() {

}
