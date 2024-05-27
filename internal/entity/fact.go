package entity

import "time"

type Fact struct {
	PeriodStart time.Time `json:"period_start"`
	PeriodEnd   time.Time `json:"period_end"`
	// PeriodKey           string    `json:"period_key"`
	// IndicatorToMoId     int       `json:"indicator_to_mo_id"`
	// IndicatorToMoFactId int       `json:"indicator_to_mo_fact_id"`
	// Value               int       `json:"value"`
	// FactTime            time.Time `json:"fact_time"`
	// IsPlan              int       `json:"is_plan"`
	// AuthUserId          int       `json:"auth_user_id"`
	Comment string `json:"comment"`
}

type Facts struct {
	Limit int     `json:"limit"`
	Facts []*Fact `json:"facts"`
}
