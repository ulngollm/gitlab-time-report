package api

import "time"

type Issue struct {
	ProjectId int `json:"project_id"`
	Author    struct {
		State     string      `json:"state"`
		WebUrl    string      `json:"web_url"`
		AvatarUrl interface{} `json:"avatar_url"`
		Username  string      `json:"username"`
		Id        int         `json:"id"`
		Name      string      `json:"name"`
	} `json:"author"`
	Description string `json:"description"`
	State       string `json:"state"`
	Assignee    struct {
		AvatarUrl interface{} `json:"avatar_url"`
		WebUrl    string      `json:"web_url"`
		State     string      `json:"state"`
		Username  string      `json:"username"`
		Id        int         `json:"id"`
		Name      string      `json:"name"`
	} `json:"assignee"`
	Type       string    `json:"type"`
	Labels     []string  `json:"labels"`
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
	ClosedAt   time.Time `json:"closed_at"`
	DueDate    string    `json:"due_date"`
	WebUrl     string    `json:"web_url"`
	References struct {
		Short    string `json:"short"`
		Relative string `json:"relative"`
		Full     string `json:"full"`
	} `json:"references"`
	TimeStats struct {
		TimeEstimate        int     `json:"time_estimate"`
		TotalTimeSpent      int     `json:"total_time_spent"`
		HumanTimeEstimate   *string `json:"human_time_estimate"`
		HumanTotalTimeSpent *string `json:"human_total_time_spent"`
	} `json:"time_stats"`
	IssueType string `json:"issue_type"`
}
