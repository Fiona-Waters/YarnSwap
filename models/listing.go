package models

import (
	"time"
)

// Listing represents data about a YarnSwap listing

type Listing struct {
	ID             string    `json:"id,omitempty"`
	UserId         string    `json:"userId"`
	UserName       string    `json:"userName"`
	Brand          string    `json:"brand"`
	Colourway      string    `json:"colourway"`
	Meterage       int       `json:"meterage"`
	Weight         string    `json:"weight"`
	FibreContent   string    `json:"fibreContent"`
	UnitWeight     int       `json:"unitWeight"`
	DyeLot         string    `json:"dyeLot,omitempty"`
	Swappable      bool      `json:"swappable"`
	Cost           float64   `json:"cost,omitempty"`
	OriginalCount  int       `json:"originalCount"`
	RemainingCount int       `json:"remainingCount"`
	Timestamp      time.Time `json:"timestamp"`
	ImageUrl       string    `json:"image"`
	Status         string    `json:"status"`
	ListingNote    string    `json:"listingNote"`
	//	Status         *ListingStatus `json:"status"`
}

type Weight struct {
	WeightName    string `json:"weightName"`
	WeightAltName string `json:"weightAltName"`
	Enabled       bool   `json:"enabled"`
	SortOrder     int    `json:"sortOrder"` // is this the correct type?
}

type FibreContent struct {
	FibreId   string `json:"fibreId"`
	FibreName string `json:"fibreName"`
	Enabled   bool   `json:"enabled"`
	SortOrder int    `json:"sortOrder"`
}

type Brand struct {
	BrandId   string `json:"brandId"`
	BrandName string `json:"brandName"`
	Enabled   bool   `json:"enabled"`
	SortOrder int    `json:"sortOrder"`
}

//type ListingStatus struct {
//	StatusId   string `json:"statusId"`
//	StatusName string `json:"statusName"`
//	Enabled    bool   `json:"enabled"`
//	SortOrder  int    `json:"sortOrder"`
//}

type Swap struct {
	ID             string `json:"id,omitempty"`
	SwapName       string `json:"swapName"`
	SwapperUserID  string `json:"swapperUserId"`
	SwappeeUserID  string `json:"swappeeUserId"`
	ListingID      string `json:"listingId"`
	SwapStatus     string `json:"swapStatus"`
	ChatChannelUrl string `json:"chatChannelUrl"`
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

type Wishlist struct {
	UserWishlist *[]Listing `json:"userWishlist"`
	Project      *Project   `json:"projects"`
}

type Project struct {
	ProjectName   string     `json:"projectName"`
	ProjectID     string     `json:"projectId"`
	LinkToPattern string     `json:"linkToPattern"`
	DateAdded     time.Time  `json:"dateAdded"`
	Listings      *[]Listing `json:"listings"`
}
