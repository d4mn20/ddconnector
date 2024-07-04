package types

import "time"

type Product struct {
	ID            int      `json:"id"`
	FindingsCount int      `json:"findings_count"`
	FindingsList  []int    `json:"findings_list"`
	Tags          []string `json:"tags"`
	ProductMeta   []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"product_meta"`
	Name                          string    `json:"name"`
	Description                   string    `json:"description"`
	Created                       time.Time `json:"created"`
	ProdNumericGrade              int64     `json:"prod_numeric_grade"`
	BusinessCriticality           string    `json:"business_criticality"`
	Platform                      string    `json:"platform"`
	Lifecycle                     string    `json:"lifecycle"`
	Origin                        string    `json:"origin"`
	UserRecords                   int64     `json:"user_records"`
	Revenue                       string    `json:"revenue"`
	ExternalAudience              bool      `json:"external_audience"`
	InternetAccessible            bool      `json:"internet_accessible"`
	EnableProductTagInheritance   bool      `json:"enable_product_tag_inheritance"`
	EnableSimpleRiskAcceptance    bool      `json:"enable_simple_risk_acceptance"`
	EnableFullRiskAcceptance      bool      `json:"enable_full_risk_acceptance"`
	DisableSLABreachNotifications bool      `json:"disable_sla_breach_notifications"`
	ProductManager                int       `json:"product_manager"`
	TechnicalContact              int       `json:"technical_contact"`
	TeamManager                   int       `json:"team_manager"`
	ProdType                      int       `json:"prod_type"`
	SLAConfiguration              int       `json:"sla_configuration"`
	Members                       []int     `json:"members"`
	AuthorizationGroups           []int     `json:"authorization_groups"`
	Regulations                   []int     `json:"regulations"`
	Prefetch                      struct {
		AuthorizationGroups struct {
		} `json:"authorization_groups"`
	} `json:"prefetch"`
}

type RespProducts struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Product
}
