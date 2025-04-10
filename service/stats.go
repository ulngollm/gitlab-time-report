package service

import (
	"github.com/ulngollm/time-report/api"
	"time"
)

type StatsService struct {
	api *api.GitlabAPI
}

func NewStatsService(api *api.GitlabAPI) *StatsService {
	return &StatsService{api: api}
}

func (s StatsService) GeTotalTimeSpend() (time.Duration, error) {
	return time.Duration(3600 * time.Second), nil
}
