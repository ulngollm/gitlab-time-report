package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/ulngollm/time-report/api"
	"github.com/ulngollm/time-report/service"
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
	Mode      mode   `short:"m" long:"mode" default:"stats" choice:"stats" choice:"report"`
	ApiToken  string `long:"token" env:"TOKEN" required:"true" description:"Gitlab personal access token"`
	ApiHost   string `long:"host" env:"API_HOST" required:"true" description:"Gitlab base API URL"`
	Labels    string `long:"labels" env:"LABELS" required:"true" default:"any"`
	ProjectID int    `long:"project" env:"PROJECT_ID" required:"true" description:"Gitlab project ID"`
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

		fmt.Println("total time spend", spend)
	case modeReport:
		repo, err := cmd.s.GetReport()
		if err != nil {
			log.Printf("getTotalTimeSpend: %s", err)
			os.Exit(1)
		}
		fmt.Println(repo)
	}
}
