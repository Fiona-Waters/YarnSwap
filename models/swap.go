package models

type Swap struct {
	ID             string `json:"id,omitempty"`
	SwapName       string `json:"swapName"`
	SwapperUserID  string `json:"swapperUserId"`
	SwappeeUserID  string `json:"swappeeUserId"`
	ListingID      string `json:"listingId"`
	SwapStatus     string `json:"swapStatus"`
	ChatChannelUrl string `json:"chatChannelUrl"`
	SwapNote       string `json:"swapNote"`
}

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
