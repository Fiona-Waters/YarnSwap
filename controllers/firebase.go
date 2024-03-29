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
		BrandId:   "0",
		BrandName: "Other",
		Enabled:   true,
		SortOrder: 1,
	},
	{
		BrandId:   "1",
		BrandName: "Bergere de France",
		Enabled:   true,
		SortOrder: 2,
	},
	{
		BrandId:   "2",
		BrandName: "Brooklyn Tweed",
		Enabled:   true,
		SortOrder: 3,
	}, {
		BrandId:   "3",
		BrandName: "Cygnet Yarns",
		Enabled:   true,
		SortOrder: 4,
	}, {
		BrandId:   "4",
		BrandName: "Debbie Bliss",
		Enabled:   true,
		SortOrder: 5,
	}, {
		BrandId:   "5",
		BrandName: "Donegal Yarns",
		Enabled:   true,
		SortOrder: 6,
	}, {
		BrandId:   "6",
		BrandName: "Drops Design",
		Enabled:   true,
		SortOrder: 7,
	}, {
		BrandId:   "7",
		BrandName: "Green Elephant Yarn",
		Enabled:   true,
		SortOrder: 8,
	}, {
		BrandId:   "8",
		BrandName: "Hayfield",
		Enabled:   true,
		SortOrder: 9,
	}, {
		BrandId:   "9",
		BrandName: "James C Brett",
		Enabled:   true,
		SortOrder: 10,
	}, {
		BrandId:   "10",
		BrandName: "King Cole",
		Enabled:   true,
		SortOrder: 11,
	}, {
		BrandId:   "11",
		BrandName: "Madelinetosh",
		Enabled:   true,
		SortOrder: 12,
	}, {
		BrandId:   "12",
		BrandName: "Malabrigo Yarn",
		Enabled:   true,
		SortOrder: 13,
	}, {
		BrandId:   "13",
		BrandName: "Manos del Uruguay",
		Enabled:   true,
		SortOrder: 14,
	}, {
		BrandId:   "14",
		BrandName: "Red Heart",
		Enabled:   true,
		SortOrder: 15,
	}, {
		BrandId:   "15",
		BrandName: "Rowan",
		Enabled:   true,
		SortOrder: 16,
	}, {
		BrandId:   "16",
		BrandName: "Schoppell Wolle",
		Enabled:   true,
		SortOrder: 17,
	}, {
		BrandId:   "17",
		BrandName: "Sirdar",
		Enabled:   true,
		SortOrder: 18,
	}, {
		BrandId:   "18",
		BrandName: "Stylecraft",
		Enabled:   true,
		SortOrder: 19,
	}, {
		BrandId:   "19",
		BrandName: "Sublime",
		Enabled:   true,
		SortOrder: 20,
	}, {
		BrandId:   "20",
		BrandName: "Yarn Vibes",
		Enabled:   true,
		SortOrder: 21,
	},
}

var Weights = []models.Weight{
	{
		WeightId:      "0",
		WeightName:    "Other",
		WeightAltName: "Other",
		Enabled:       true,
		SortOrder:     1,
	},
	{
		WeightId:      "1",
		WeightName:    "Lace",
		WeightAltName: "2ply",
		Enabled:       true,
		SortOrder:     2,
	},
	{
		WeightId:      "2",
		WeightName:    "Sock",
		WeightAltName: "4ply",
		Enabled:       true,
		SortOrder:     3,
	},
	{
		WeightId:      "3",
		WeightName:    "Sport",
		WeightAltName: "5ply",
		Enabled:       true,
		SortOrder:     4,
	},
	{
		WeightId:      "4",
		WeightName:    "DK",
		WeightAltName: "8ply",
		Enabled:       true,
		SortOrder:     5,
	},
	{
		WeightId:      "5",
		WeightName:    "Aran or Worsted",
		WeightAltName: "10ply",
		Enabled:       true,
		SortOrder:     6,
	},
	{
		WeightId:      "6",
		WeightName:    "Chunky",
		WeightAltName: "14-20ply",
		Enabled:       true,
		SortOrder:     7,
	},
	{
		WeightId:      "7",
		WeightName:    "Super Bulky",
		WeightAltName: "14ply",
		Enabled:       true,
		SortOrder:     8,
	},
	{
		WeightId:      "8",
		WeightName:    "Jumbo or Roving",
		WeightAltName: "14ply",
		Enabled:       true,
		SortOrder:     9,
	},
}

var Fibres = []models.FibreContent{
	{
		FibreId:   "0",
		FibreName: "Other",
		Enabled:   true,
		SortOrder: 1,
	},
	{
		FibreId:   "1",
		FibreName: "Wool",
		Enabled:   true,
		SortOrder: 2,
	},
	{
		FibreId:   "2",
		FibreName: "Alpaca",
		Enabled:   true,
		SortOrder: 3,
	}, {
		FibreId:   "3",
		FibreName: "Cashmere",
		Enabled:   true,
		SortOrder: 4,
	}, {
		FibreId:   "4",
		FibreName: "Mohair",
		Enabled:   true,
		SortOrder: 5,
	}, {
		FibreId:   "5",
		FibreName: "Cotton",
		Enabled:   true,
		SortOrder: 6,
	}, {
		FibreId:   "6",
		FibreName: "Linen",
		Enabled:   true,
		SortOrder: 7,
	}, {
		FibreId:   "7",
		FibreName: "Bamboo",
		Enabled:   true,
		SortOrder: 8,
	}, {
		FibreId:   "8",
		FibreName: "Silk",
		Enabled:   true,
		SortOrder: 9,
	}, {
		FibreId:   "9",
		FibreName: "Nylon",
		Enabled:   true,
		SortOrder: 10,
	}, {
		FibreId:   "10",
		FibreName: "Acrylic",
		Enabled:   true,
		SortOrder: 11,
	},
}

var ListingStatuses = []models.ListingStatus{
	{
		StatusId:   "0",
		StatusName: "Awaiting approval",
		Enabled:    true,
		SortOrder:  1,
	},
	{
		StatusId:   "1",
		StatusName: "Available",
		Enabled:    true,
		SortOrder:  2,
	},
	{
		StatusId:   "2",
		StatusName: "Declined",
		Enabled:    true,
		SortOrder:  3,
	},
	{
		StatusId:   "3",
		StatusName: "Swap requested",
		Enabled:    true,
		SortOrder:  4,
	},
	{
		StatusId:   "4",
		StatusName: "Archived",
		Enabled:    true,
		SortOrder:  5,
	},
}

// InitialiseFirebaseApp function initialising firebase app and database.
func InitialiseFirebaseApp() (context.Context, *db.Client, *firebase.App) {
	ctx := context.Background()

	conf := &firebase.Config{
		AuthOverride: nil,
		DatabaseURL:  "https://yarnswap-52dbd-default-rtdb.europe-west1.firebasedatabase.app",
	}

	opt := option.WithCredentialsFile("credentials/yarnswap-firebase.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initialising database client:", err)
	}

	return ctx, client, app
}

func addBrandsToFirebase(ctx context.Context, client *db.Client) {
	ref := client.NewRef("brands")

	for _, v := range Brands {
		brandRef := ref.Child(v.BrandId)
		err := brandRef.Set(ctx, v)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	}
}

func addWeightsToFirebase(ctx context.Context, client *db.Client) {
	ref := client.NewRef("weights")

	for _, v := range Weights {
		weightRef := ref.Child(v.WeightId)
		err := weightRef.Set(ctx, v)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	}
}

func addFibresToFirebase(ctx context.Context, client *db.Client) {
	ref := client.NewRef("fibres")

	for _, v := range Fibres {
		fibreRef := ref.Child(v.FibreId)
		err := fibreRef.Set(ctx, v)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	}
}

func addListingStatusesToFirebase(ctx context.Context, client *db.Client) {
	ref := client.NewRef("listing-status")

	for _, v := range ListingStatuses {
		statusRef := ref.Child(v.StatusId)
		err := statusRef.Set(ctx, v)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	}
}

func PopulateFirebase(ctx context.Context, client *db.Client) {
	addBrandsToFirebase(ctx, client)
	addWeightsToFirebase(ctx, client)
	addFibresToFirebase(ctx, client)
	addListingStatusesToFirebase(ctx, client)
}
