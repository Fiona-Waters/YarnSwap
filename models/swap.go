package models

import (
	"time"
)

// Swap struct represents detail about a swap
type Swap struct {
	ID             string    `json:"id,omitempty"`
	SwapName       string    `json:"swapName"`
	SwapperUserID  string    `json:"swapperUserId"`
	SwappeeUserID  string    `json:"swappeeUserId"`
	ListingID      string    `json:"listingId"`
	SwapStatus     string    `json:"swapStatus"`
	ChatChannelUrl string    `json:"chatChannelUrl"`
	SwapNote       string    `json:"swapNote"`
	Timestamp      time.Time `json:"timestamp"`
}

// SwapListing an object containing a swap object and a listing object
type SwapListing struct {
	Swap    Swap    `json:"swap"`
	Listing Listing `json:"listing"`
}

//
//type SwapStatus struct {
//	StatusId   string `json:"statusId"`
//	StatusName string `json:"statusName"`
//	Enabled    bool   `json:"enabled"`
//	SortOrder  int    `json:"sortOrder"`
//}
