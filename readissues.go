package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var ApiRoot = "https://api.github.com"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in MarkDown Format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// ReadIssue searches for the issue and prints it out to read
func ReadIssues(query []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(query, " "))
	u := fmt.Sprintf("%s/search/issues?q=%s", ApiRoot, q)

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, err
}
