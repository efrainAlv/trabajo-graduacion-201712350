package types

type Message struct {
    Time int `json:"current_time"`
	Symbol string `json:"symbol"`
	Price string `json:"price"`
}