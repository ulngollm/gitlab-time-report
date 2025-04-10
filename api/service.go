package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	projectIssuesEndpoint = "/api/v4/projects/%d/issues/?per_page=%d&labels=%s&state=opened&scope=all&order_by=relative_position&sort=asc"
	perPage               = 100
)

type Config struct {
	Host      string
	Token     string
	ProjectID int
	Labels    string
}

type GitlabAPI struct {
	conf   Config
	client http.Client
}

func NewGitlabAPI(conf Config) *GitlabAPI {
	return &GitlabAPI{
		conf:   conf,
		client: http.Client{Timeout: time.Second * 10},
	}
}

func (a GitlabAPI) GetIssues() ([]Issue, error) {
	endpoint := fmt.Sprintf(a.conf.Host+projectIssuesEndpoint, a.conf.ProjectID, perPage, a.conf.Labels)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("PRIVATE-TOKEN", a.conf.Token)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status is not ok: %s", resp.Status)
	}

	var response []Issue
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	return response, nil
}
