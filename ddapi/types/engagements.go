package types

import "time"

type Engagement struct {
	ID                         int       `json:"id,omitempty"`
	Tags                       []string  `json:"tags,omitempty"`
	Name                       string    `json:"name"`
	Description                string    `json:"description"`
	Version                    string    `json:"version"`
	FirstContacted             string    `json:"first_contacted"`
	TargetStart                string    `json:"target_start"`
	TargetEnd                  string    `json:"target_end"`
	Reason                     string    `json:"reason"`
	Updated                    time.Time `json:"updated"`
	Created                    time.Time `json:"created"`
	Active                     bool      `json:"active"`
	Tracker                    string    `json:"tracker"`
	TestStrategy               string    `json:"test_strategy"`
	ThreatModel                bool      `json:"threat_model"`
	APITest                    bool      `json:"api_test"`
	PenTest                    bool      `json:"pen_test"`
	CheckList                  bool      `json:"check_list"`
	Status                     string    `json:"status,omitempty"`
	Progress                   string    `json:"progress"`
	TmodelPath                 string    `json:"tmodel_path"`
	DoneTesting                bool      `json:"done_testing"`
	EngagementType             string    `json:"engagement_type"`
	BuildID                    string    `json:"build_id"`
	CommitHash                 string    `json:"commit_hash"`
	BranchTag                  string    `json:"branch_tag"`
	SourceCodeManagementURI    string    `json:"source_code_management_uri"`
	DeduplicationOnEngagement  bool      `json:"deduplication_on_engagement"`
	Lead                       *int      `json:"lead,omitempty"`
	Requester                  *int      `json:"requester,omitempty"`
	Preset                     *int      `json:"preset,omitempty"`
	ReportType                 *int      `json:"report_type,omitempty"`
	Product                    int       `json:"product"`
	BuildServer                *int      `json:"build_server,omitempty"`
	SourceCodeManagementServer *int      `json:"source_code_management_server,omitempty"`
	OrchestrationEngine        *int      `json:"orchestration_engine,omitempty"`
	RiskAcceptance             []int     `json:"risk_acceptance,omitempty"`
}

type RespEngagements struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Engagement
}
