package service

import "github.com/ulngollm/time-report/api"

type Storage struct {
	issues []api.Issue
}

func NewStorage() *Storage {
	return &Storage{}
}
