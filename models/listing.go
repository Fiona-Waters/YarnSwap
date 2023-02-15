package models

import "time"

// Listing represents data about a YarnSwap listing

type Listing struct {
	ID             string         `json:"id"`
	UserId         string         `json:"userId"`
	Brand          *Brand         `json:"brand"`
	Colourway      string         `json:"colourway"`
	Meterage       string         `json:"meterage"`
	Weight         *Weight        `json:"settings"`
	FibreContent   *FibreContent  `json:"fibreContent"`
	UnitWeight     int            `json:"unitWeight"`
	DyeLot         string         `json:"dyeLot"`
	Swappable      bool           `json:"swappable"`
	Cost           float64        `json:"cost"`
	OriginalCount  int            `json:"originalCount"`
	RemainingCount int            `json:"remainingCount"`
	Timestamp      time.Time      `json:"timestamp"`
	Image          *listingImage  `json:"image"`
	Status         *listingStatus `json:"status"`
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

type listingImage struct {
	ImageId string `json:"imageId"`
}

type listingStatus struct {
	StatusId   string `json:"statusId"`
	StatusName string `json:"statusName"`
	Enabled    bool   `json:"enabled"`
	SortOrder  int    `json:"sortOrder"`
}