package controllers

import (
	"context"
	"fionawaters/YarnSwap/models"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
	"log"
)

var Brands = []models.Brand{
	{
		BrandId:   "1",
		BrandName: "Bergere de France",
		Enabled:   true,
		SortOrder: 1,
	},
	{
		BrandId:   "2",
		BrandName: "Brooklyn Tweed",
		Enabled:   true,
		SortOrder: 2,
	}, {
		BrandId:   "3",
		BrandName: "Cygnet Yarns",
		Enabled:   true,
		SortOrder: 3,
	}, {
		BrandId:   "4",
		BrandName: "Debbie Bliss",
		Enabled:   true,
		SortOrder: 4,
	}, {
		BrandId:   "5",
		BrandName: "Donegal Yarns",
		Enabled:   true,
		SortOrder: 5,
	}, {
		BrandId:   "6",
		BrandName: "Drops Design",
		Enabled:   true,
		SortOrder: 6,
	}, {
		BrandId:   "7",
		BrandName: "Green Elephant Yarn",
		Enabled:   true,
		SortOrder: 7,
	}, {
		BrandId:   "8",
		BrandName: "Hayfield",
		Enabled:   true,
		SortOrder: 8,
	}, {
		BrandId:   "9",
		BrandName: "James C Brett",
		Enabled:   true,
		SortOrder: 9,
	}, {
		BrandId:   "10",
		BrandName: "King Cole",
		Enabled:   true,
		SortOrder: 10,
	}, {
		BrandId:   "11",
		BrandName: "Madelinetosh",
		Enabled:   true,
		SortOrder: 11,
	}, {
		BrandId:   "12",
		BrandName: "Malabrigo Yarn",
		Enabled:   true,
		SortOrder: 12,
	}, {
		BrandId:   "13",
		BrandName: "Manos del Uruguay",
		Enabled:   true,
		SortOrder: 13,
	}, {
		BrandId:   "14",
		BrandName: "Red Heart",
		Enabled:   true,
		SortOrder: 14,
	}, {
		BrandId:   "15",
		BrandName: "Rowan",
		Enabled:   true,
		SortOrder: 15,
	}, {
		BrandId:   "16",
		BrandName: "Schoppell Wolle",
		Enabled:   true,
		SortOrder: 16,
	}, {
		BrandId:   "17",
		BrandName: "Sirdar",
		Enabled:   true,
		SortOrder: 17,
	}, {
		BrandId:   "18",
		BrandName: "Stylecraft",
		Enabled:   true,
		SortOrder: 18,
	}, {
		BrandId:   "19",
		BrandName: "Sublime",
		Enabled:   true,
		SortOrder: 19,
	}, {
		BrandId:   "20",
		BrandName: "Yarn Vibes",
		Enabled:   true,
		SortOrder: 20,
	},
}

var Weights = []models.Weight{
	{
		WeightName:    "Lace",
		WeightAltName: "2ply",
		Enabled:       true,
		SortOrder:     1,
	},
	{
		WeightName:    "Sock",
		WeightAltName: "4ply",
		Enabled:       true,
		SortOrder:     2,
	},
	{
		WeightName:    "Sport",
		WeightAltName: "5ply",
		Enabled:       true,
		SortOrder:     3,
	}, {
		WeightName:    "DK",
		WeightAltName: "8ply",
		Enabled:       true,
		SortOrder:     4,
	}, {
		WeightName:    "Aran or Worsted",
		WeightAltName: "10ply",
		Enabled:       true,
		SortOrder:     5,
	}, {
		WeightName:    "Chunky",
		WeightAltName: "14-20ply",
		Enabled:       true,
		SortOrder:     6,
	}, {
		WeightName:    "Super Bulky",
		WeightAltName: "14ply",
		Enabled:       true,
		SortOrder:     7,
	}, {
		WeightName:    "Jumbo or Roving",
		WeightAltName: "14ply",
		Enabled:       true,
		SortOrder:     8,
	},
}

var Fibres = []models.FibreContent{
	{
		FibreId:   "1",
		FibreName: "Wool",
		Enabled:   true,
		SortOrder: 11,
	},
	{
		FibreId:   "2",
		FibreName: "Alpaca",
		Enabled:   true,
		SortOrder: 2,
	}, {
		FibreId:   "3",
		FibreName: "Cashmere",
		Enabled:   true,
		SortOrder: 4,
	}, {
		FibreId:   "4",
		FibreName: "Mohair",
		Enabled:   true,
		SortOrder: 7,
	}, {
		FibreId:   "5",
		FibreName: "Cotton",
		Enabled:   true,
		SortOrder: 5,
	}, {
		FibreId:   "6",
		FibreName: "Linen",
		Enabled:   true,
		SortOrder: 6,
	}, {
		FibreId:   "7",
		FibreName: "Bamboo",
		Enabled:   true,
		SortOrder: 3,
	}, {
		FibreId:   "8",
		FibreName: "Silk",
		Enabled:   true,
		SortOrder: 10,
	}, {
		FibreId:   "9",
		FibreName: "Nylon",
		Enabled:   true,
		SortOrder: 8,
	}, {
		FibreId:   "10",
		FibreName: "Acrylic",
		Enabled:   true,
		SortOrder: 1,
	},
	{
		FibreId:   "11",
		FibreName: "Other",
		Enabled:   true,
		SortOrder: 11,
	},
}

//var ListingStatuses = []models.ListingStatus{
//	{
//		StatusId:   "1",
//		StatusName: "Available",
//		Enabled:    true,
//		SortOrder:  1,
//	},
//	{
//		StatusId:   "2",
//		StatusName: "Unavailable",
//		Enabled:    true,
//		SortOrder:  2,
//	},
//	{
//		StatusId:   "3",
//		StatusName: "Inactive",
//		Enabled:    true,
//		SortOrder:  3,
//	},
//	{
//		StatusId:   "4",
//		StatusName: "Archived",
//		Enabled:    true,
//		SortOrder:  4,
//	},
//}

// InitialiseFirebaseApp function initialising firebase app and database and posting 2 listings.
func InitialiseFirebaseApp() (context.Context, *db.Client) {
	ctx := context.Background()

	conf := &firebase.Config{
		AuthOverride: nil,
		DatabaseURL:  "https://yarnswap-52dbd-default-rtdb.europe-west1.firebasedatabase.app",
	}

	opt := option.WithCredentialsFile("yarnswap-firebase.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initialising database client:", err)
	}

	ref := client.NewRef("listings")
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading data from database:", err)
	}
	//fmt.Println(data)

	return ctx, client
}

func addBrandsToFirebase() {
	ctx, client := InitialiseFirebaseApp()
	ref := client.NewRef("brands")

	for _, v := range Brands {
		brandRef := ref.Child(v.BrandId)
		err := brandRef.Set(ctx, v)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	}
}

func addWeightsToFirebase() {
	ctx, client := InitialiseFirebaseApp()
	ref := client.NewRef("weights")

	for _, v := range Weights {
		weightRef := ref.Child(v.WeightName)
		err := weightRef.Set(ctx, v)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	}
}

func addFibresToFirebase() {
	ctx, client := InitialiseFirebaseApp()
	ref := client.NewRef("fibres")

	for _, v := range Fibres {
		fibreRef := ref.Child(v.FibreId)
		err := fibreRef.Set(ctx, v)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	}
}

//func addListingStatusesToFirebase() {
//	ctx, client := InitialiseFirebaseApp()
//	ref := client.NewRef("listing-status")
//
//	for _, v := range ListingStatuses {
//		statusRef := ref.Child(v.StatusId)
//		err := statusRef.Set(ctx, v)
//		if err != nil {
//			log.Fatalln("Error setting value:", err)
//		}
//	}
//}

func PopulateFirebase() {
	addBrandsToFirebase()
	addWeightsToFirebase()
	addFibresToFirebase()
	//	addListingStatusesToFirebase()
}
