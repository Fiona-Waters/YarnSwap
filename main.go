package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
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

// function to retrieve listings from firebase realtime database
func getListings(c *gin.Context) {
	ctx, client := initialiseFirebaseApp()

	//Create Ref for listings
	ref := client.NewRef("listings")
	//retrieve the listings in order of the keys
	results, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	//create an array the same length as the number of results
	data := make([]Listing, len(results))

	//loop over the results and individually marshal into Listing struct
	for i, r := range results {
		var l Listing
		if err := r.Unmarshal(&l); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		//add new struct to array
		data[i] = l
	}

	log.Default().Println("data = ", data)

	c.IndentedJSON(http.StatusOK, data)
}

//function to retrieve listing by id
func getListingById(c *gin.Context) {
	id := c.Param("id")
	//loop over listings to find one with requested id
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
	router.Use(cors.Default())
	router.GET("/listings", getListings)
	router.GET("/listings/:id", getListingById)
	router.POST("/listings", addListing)

	router.Run("0.0.0.0:8080")

}

// function initialising firebase app and database and posting 2 listings.
func initialiseFirebaseApp() (context.Context, *db.Client) {
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
	fmt.Println(data)

	return ctx, client
}
