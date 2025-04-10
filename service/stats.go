package service

import (
	"fmt"
	"time"

	"github.com/ulngollm/time-report/api"
)

type StatsService struct {
	api *api.GitlabAPI
}

func NewStatsService(api *api.GitlabAPI) *StatsService {
	return &StatsService{api: api}
}

func (s StatsService) GetTotalTimeSpend() (time.Duration, error) {
	issues, err := s.api.GetIssues()
	if err != nil {
		return 0, fmt.Errorf("getIssues: %w", err)
	}

	totalTime := time.Duration(0)
	for _, issue := range issues {
		totalTime += time.Duration(issue.TimeStats.TotalTimeSpent) * time.Second
	}

	return totalTime, nil
}
