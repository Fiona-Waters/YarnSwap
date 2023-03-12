package models

import "time"

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
