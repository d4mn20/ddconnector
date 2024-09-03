package types

import "time"

type Test struct {
	ID            int      `json:"id"`
	Tags          []string `json:"tags"`
	TestTypeName  string   `json:"test_type_name"`
	FindingGroups []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Test      int    `json:"test"`
		JiraIssue struct {
			ID           int       `json:"id"`
			URL          string    `json:"url"`
			JiraID       string    `json:"jira_id"`
			JiraKey      string    `json:"jira_key"`
			JiraCreation time.Time `json:"jira_creation"`
			JiraChange   time.Time `json:"jira_change"`
			JiraProject  int       `json:"jira_project"`
			Finding      int       `json:"finding"`
			Engagement   int       `json:"engagement"`
			FindingGroup int       `json:"finding_group"`
		} `json:"jira_issue"`
	} `json:"finding_groups"`
	ScanType             string    `json:"scan_type"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	TargetStart          string    `json:"target_start"`
	TargetEnd            string    `json:"target_end"`
	EstimatedTime        string    `json:"estimated_time"`
	ActualTime           string    `json:"actual_time"`
	PercentComplete      int64     `json:"percent_complete"`
	Updated              time.Time `json:"updated"`
	Created              time.Time `json:"created"`
	Version              string    `json:"version"`
	BuildID              string    `json:"build_id"`
	CommitHash           string    `json:"commit_hash"`
	BranchTag            string    `json:"branch_tag"`
	Engagement           int       `json:"engagement"`
	Lead                 int       `json:"lead,omitempty"`
	TestType             int       `json:"test_type"`
	Environment          int       `json:"environment,omitempty"`
	APIScanConfiguration int       `json:"api_scan_configuration,omitempty"`
}

type RespTests struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Test `json:"results"`
}

type TestResult struct {
	ScanDate                          string   `json:"scan_date"`
	MinimumSeverity                   string   `json:"minimum_severity"`
	Active                            bool     `json:"active"`
	Verified                          bool     `json:"verified"`
	ScanType                          string   `json:"scan_type"`
	EndpointToAdd                     int      `json:"endpoint_to_add"`
	File                              string   `json:"file"`
	ProductTypeName                   string   `json:"product_type_name"`
	ProductName                       string   `json:"product_name"`
	EngagementName                    string   `json:"engagement_name"`
	EngagementEndDate                 string   `json:"engagement_end_date"`
	SourceCodeManagementURI           string   `json:"source_code_management_uri"`
	Engagement                        int      `json:"engagement"`
	TestTitle                         string   `json:"test_title"`
	AutoCreateContext                 bool     `json:"auto_create_context"`
	DeduplicationOnEngagement         bool     `json:"deduplication_on_engagement"`
	Lead                              int      `json:"lead"`
	Tags                              []string `json:"tags"`
	CloseOldFindings                  bool     `json:"close_old_findings"`
	CloseOldFindingsProductScope      bool     `json:"close_old_findings_product_scope"`
	PushToJira                        bool     `json:"push_to_jira"`
	Environment                       string   `json:"environment"`
	Version                           string   `json:"version"`
	BuildID                           string   `json:"build_id"`
	BranchTag                         string   `json:"branch_tag"`
	CommitHash                        string   `json:"commit_hash"`
	APIScanConfiguration              int      `json:"api_scan_configuration"`
	Service                           string   `json:"service"`
	GroupBy                           string   `json:"group_by"`
	CreateFindingGroupsForAllFindings bool     `json:"create_finding_groups_for_all_findings"`
	Test                              int      `json:"test"`
	TestID                            int      `json:"test_id"`
	EngagementID                      int      `json:"engagement_id"`
	ProductID                         int      `json:"product_id"`
	ProductTypeID                     int      `json:"product_type_id"`
	Statistics                        struct {
		Before struct {
			Info struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"info"`
			Low struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"low"`
			Medium struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"medium"`
			High struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"high"`
			Critical struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"critical"`
			Total struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"total"`
		} `json:"before"`
		Delta struct {
			Created struct {
				Info struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"info"`
				Low struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"low"`
				Medium struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"medium"`
				High struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"high"`
				Critical struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"critical"`
				Total struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"total"`
			} `json:"created"`
			Closed struct {
				Info struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"info"`
				Low struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"low"`
				Medium struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"medium"`
				High struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"high"`
				Critical struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"critical"`
				Total struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"total"`
			} `json:"closed"`
			Reactivated struct {
				Info struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"info"`
				Low struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"low"`
				Medium struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"medium"`
				High struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"high"`
				Critical struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"critical"`
				Total struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"total"`
			} `json:"reactivated"`
			LeftUntouched struct {
				Info struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"info"`
				Low struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"low"`
				Medium struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"medium"`
				High struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"high"`
				Critical struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"critical"`
				Total struct {
					Active       int `json:"active"`
					Verified     int `json:"verified"`
					Duplicate    int `json:"duplicate"`
					FalseP       int `json:"false_p"`
					OutOfScope   int `json:"out_of_scope"`
					IsMitigated  int `json:"is_mitigated"`
					RiskAccepted int `json:"risk_accepted"`
					Total        int `json:"total"`
				} `json:"total"`
			} `json:"left untouched"`
		} `json:"delta"`
		After struct {
			Info struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"info"`
			Low struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"low"`
			Medium struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"medium"`
			High struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"high"`
			Critical struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"critical"`
			Total struct {
				Active       int `json:"active"`
				Verified     int `json:"verified"`
				Duplicate    int `json:"duplicate"`
				FalseP       int `json:"false_p"`
				OutOfScope   int `json:"out_of_scope"`
				IsMitigated  int `json:"is_mitigated"`
				RiskAccepted int `json:"risk_accepted"`
				Total        int `json:"total"`
			} `json:"total"`
		} `json:"after"`
	} `json:"statistics"`
}
