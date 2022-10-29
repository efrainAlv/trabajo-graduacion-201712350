package types

type Message struct {
    Time int 		`json:"current_time"`
	Symbol string 	`json:"symbol"`
	Price string 	`json:"price"`
}

type HistoryPrice struct {
	TimePeriodStart string `json:"time_period_start"`
	TimePeriodEnd 	string `json:"time_period_end"`
	TimeOpen 		string `json:"time_open"`
	TimeClose 		string `json:"time_close"`
	RateOpen 		string `json:"rate_open"`
	RateHigh 		string `json:"rate_high"`
	RateLow 		string `json:"rate_low"`
	RateClose 		string `json:"rate_close"`
}
