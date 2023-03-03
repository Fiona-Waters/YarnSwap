package models

type User struct {
	ID              string `json:"id,omitempty"`
	UserName        string `json:"userName,omitempty"`
	RemainingTokens int    `json:"remainingTokens"`
}
