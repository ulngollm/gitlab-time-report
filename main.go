package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/ulngollm/time-report/api"
	"github.com/ulngollm/time-report/service"
	"log"
	"os"
)

type app struct {
	s *service.StatsService
}

type mode string

const (
	modeStats  mode = "stats"
	modeReport mode = "report"
)

type options struct {
	Mode      mode   `short:"m" long:"mode" default:"stats"`
	ApiToken  string `long:"token" env:"TOKEN" required:"true"`
	ApiHost   string `long:"host" env:"API_HOST" required:"true"`
	Labels    string `long:"labels" env:"LABELS" required:"true"`
	ProjectID int    `long:"project" env:"PROJECT_ID" required:"true"`
}

func main() {
	var opts options
	p := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash|flags.HelpFlag)
	if _, err := p.Parse(); err != nil {
		os.Exit(1)
	}
	config := api.Config{
		Host:      opts.ApiHost,
		Token:     opts.ApiToken,
		ProjectID: opts.ProjectID,
		Labels:    opts.Labels,
	}

	cmd := app{
		s: service.NewStatsService(
			api.NewGitlabAPI(config),
		),
	}

	switch opts.Mode {
	case modeStats:
		spend, err := cmd.s.GetTotalTimeSpend()
		if err != nil {
			log.Printf("getTotalTimeSpend: %s", err)
			os.Exit(1)
		}

		fmt.Println(spend)
	case modeReport:
		repo, err := cmd.s.GetReport()
		if err != nil {
			log.Printf("getTotalTimeSpend: %s", err)
			os.Exit(1)
		}
		fmt.Println(repo)
	}
}
