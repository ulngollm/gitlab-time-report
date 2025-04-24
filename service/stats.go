package service

import (
	"fmt"
	"time"

	"github.com/ulngollm/time-report/api"
)

type StatsService struct {
	api   *api.GitlabAPI
	cache *Storage
}

func NewStatsService(api *api.GitlabAPI) *StatsService {
	return &StatsService{
		api:   api,
		cache: NewStorage(),
	}
}

func (s *StatsService) GetTotalTimeSpend() (time.Duration, error) {
	issues, err := s.getIssues()
	if err != nil {
		return 0, fmt.Errorf("getIssues: %w", err)
	}

	totalTime := time.Duration(0)
	for _, issue := range issues {
		totalTime += time.Duration(issue.TimeStats.TotalTimeSpent) * time.Second
	}
	return totalTime, nil
}

func (s *StatsService) GetReport() (string, error) {
	issues, err := s.getIssues()
	if err != nil {
		return "", fmt.Errorf("getIssues: %w", err)
	}

	report := ""
	for _, issue := range issues {
		rawTimeSpent := issue.TimeStats.TotalTimeSpent
		var hours, minutes int
		if rawTimeSpent == 0 {
			report += fmt.Sprintf("%d,%02d:%02d\n", issue.Iid, hours, minutes)
			continue
		}

		hours = rawTimeSpent / 3600
		minutes = (rawTimeSpent % 3600) / 60
		report += fmt.Sprintf("%d,%02d:%02d\n", issue.Iid, hours, minutes)
	}
	return report, nil
}

func (s *StatsService) getIssues() ([]api.Issue, error) {
	if len(s.cache.issues) > 0 {
		return s.cache.issues, nil
	}
	issues, err := s.api.GetIssues()
	if err != nil {
		return nil, fmt.Errorf("getIssues: %w", err)
	}
	s.cache.issues = issues

	return issues, nil
}
