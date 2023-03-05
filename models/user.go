package models

import "time"

type User struct {
	ID                string    `json:"id"`
	UserName          string    `json:"userName,omitempty"`
	RemainingTokens   int       `json:"remainingTokens"`
	AccountStatus     string    `json:"accountStatus"`
	AmtListingsAdded  int       `json:"amtListingsAdded"`
	AmtSwapsCompleted int       `json:"amtSwapsCompleted"`
	Timestamp         time.Time `json:"timestamp"`
}
