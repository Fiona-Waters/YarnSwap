package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/gin-contrib/cors"
	"google.golang.org/api/option"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Listing represents data about a YarnSwap listing
type Listing struct {
	ID           string `json:"id"`
	Brand        string `json:"brand"`
	Colourway    string `json:"colourway"`
	Weight       string `json:"weight"`
	FibreContent string `json:"fibreContent"`
}

// listings to seed the database
var listings = []Listing{
	{ID: "1", Brand: "Green Elephant Yarn", Colourway: "Turquoise", Weight: "DK", FibreContent: "100% Wool"},
	{ID: "2", Brand: "Malabrigo", Colourway: "Night Sky", Weight: "fingering", FibreContent: "100% Alpaca"},
	{ID: "3", Brand: "Drops", Colourway: "Marine Blue", Weight: "Aran", FibreContent: "50% Cotton, 50% Wool"},
}

// function to retrieve listings
func getListings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, listings)
}

//function to retrieve listing by id
func getListingById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range listings {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "listing not found"})
}

// function to add a listing
func addListing(c *gin.Context) {
	var newListing Listing
	if err := c.BindJSON(&newListing); err != nil {
		return
	}
	listings = append(listings, newListing)
	c.IndentedJSON(http.StatusCreated, newListing)
}

func main() {
	initialiseFirebaseApp()

	router := gin.Default()

	router.GET("/listings", getListings)
	router.GET("/listings/:id", getListingById)
	router.POST("/listings", addListing)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	router.Use(cors.New(config))
	router.Run("localhost:8080")
}

// function initialising firebase app and database and posting 2 listings.
func initialiseFirebaseApp() *firebase.App {
	ctx := context.Background()
	//var nilMap map[string]interface{}
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
	fmt.Println(data)

	listingsRef := ref.Child("userListings")
	err = listingsRef.Set(ctx, map[string]*Listing{
		"1": {Brand: "Green Elephant Yarn", Colourway: "Turquoise", Weight: "DK", FibreContent: "100% Wool"},
		"2": {Brand: "Malabrigo", Colourway: "Night Sky", Weight: "fingering", FibreContent: "100% Alpaca"},
	})
	if err != nil {
		log.Fatalln("Error setting value", err)
	}

	return app
}
