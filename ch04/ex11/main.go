package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type githubIssue struct {
	Number  int    `json:"number,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`
	Title   string `json:"title,omitempty"`
	Body    string `json:"body,omitempty"`
	State   string `json:"state,omitempty"`
}

// APIURL Github api
const APIURL = "https://api.github.com"

// GitHubToken access token
const GitHubToken = "b5bfed530f1d4ff86a80a9d3308a51e46e02a8cc"

const (
	createCom = "Create"
	readCom   = "Read"
	editCom   = "Edit"
	closeCom  = "Close"
)

func main() {
	var (
		owner  string
		rep    string
		number int
		com    string
	)
	var issue githubIssue
	flag.StringVar(&owner, "o", "", "リポジトリのオーナー名")
	flag.StringVar(&rep, "r", "", "リポジトリ名")
	flag.IntVar(&number, "n", 0, "issueの番号")
	flag.StringVar(&com, "c", "", "実行するコマンド")
	flag.StringVar(&issue.Title, "title", "", "json タイトル")
	flag.StringVar(&issue.Body, "body", "", "json body")
	flag.Parse()

	switch com {
	case createCom:
		result, err := createIssue(owner, rep, &issue)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		resStr, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		fmt.Printf("%s\n", resStr)
	case readCom:
		result, err := readIssue(owner, rep, number)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		resStr, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		fmt.Printf("%s\n", resStr)
	case closeCom:
		result, err := closeIssue(owner, rep, number)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		resStr, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		fmt.Printf("%s\n", resStr)
	case editCom:
		result, err := editIssue(owner, rep, number, &issue)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		resStr, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		fmt.Printf("%s\n", resStr)
	default:
		log.Fatal("unkown command")
	}
}

func readIssue(owner, repos string, number int) (*githubIssue, error) {
	q := "/repos/" + owner + "/" + repos + "/issues/" + strconv.Itoa(number)
	resp, err := http.Get(APIURL + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var issue githubIssue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func createIssue(owner, repos string, input *githubIssue) (*githubIssue, error) {
	url := APIURL + "/repos/" + owner + "/" + repos + "/issues"
	jsonbytes, err := json.Marshal(input)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonbytes))
	if err != nil {
		return nil, err
	}

	resp, err := sendRequest(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("%s\n", resp.Status)
	}
	var issue githubIssue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func editIssue(owner, repos string, number int, input *githubIssue) (*githubIssue, error) {
	url := APIURL + "/repos/" + owner + "/" + repos + "/issues/" + strconv.Itoa(number)
	jsonbytes, err := json.Marshal(input)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonbytes))
	if err != nil {
		return nil, err
	}

	resp, err := sendRequest(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("%s\n", resp.Status)
	}
	var issue githubIssue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func closeIssue(owner, repos string, number int) (*githubIssue, error) {
	url := APIURL + "/repos/" + owner + "/" + repos + "/issues/" + strconv.Itoa(number)
	var st githubIssue
	st.State = "close"

	jsonStr, err := json.Marshal(st)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	resp, err := sendRequest(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("%s\n", resp.Status)
	}
	var issue githubIssue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func sendRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	token := "token " + GitHubToken
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
