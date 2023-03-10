package models

import "time"

type User struct {
	ID                string    `json:"id,omitempty"`
	UserName          string    `json:"userName,omitempty"`
	RemainingTokens   int       `json:"remainingTokens"`
	AccountStatus     string    `json:"accountStatus"`
	AmtListingsAdded  int       `json:"amtListingsAdded"`
	AmtSwapsCompleted int       `json:"amtSwapsCompleted"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
	ArchiveTimestamp  time.Time `json:"archiveTimestamp"`
	Role              string    `json:"role"`
}
