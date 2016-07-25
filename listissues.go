package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// listIssues queries the Github issue tracker and lists out the issues
func listIssues(term string) (*IssueSearchResult, error) {
	// query escape the term
	q := url.QueryEscape(term)

	// create the url
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

	return &result, nil
}
