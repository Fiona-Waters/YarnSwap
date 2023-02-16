package main

import (
	controllers "fionawaters/YarnSwap/controllers"
	"fionawaters/YarnSwap/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// listings to seed the database
// var listings = []models.Listing{
//{ID: "1", Brand: "Green Elephant Yarn", Colourway: "Turquoise", Weight: "DK", FibreContent: "100% Wool"},
//{ID: "2", Brand: "Malabrigo", Colourway: "Night Sky", Weight: "fingering", FibreContent: "100% Alpaca"},
//{ID: "3", Brand: "Drops", Colourway: "Marine Blue", Weight: "Aran", FibreContent: "50% Cotton, 50% Wool"},
//}

// function to retrieve listings from firebase realtime database
func getListings(c *gin.Context) {
	ctx, client := controllers.InitialiseFirebaseApp()

	//Create Ref for listings
	ref := client.NewRef("listings")
	//retrieve the listings in order of the keys
	results, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	//create an array the same length as the number of results
	data := make([]models.Listing, len(results))

	//loop over the results and individually marshal into Listing struct
	for i, r := range results {
		var l models.Listing
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
//func getListingById(c *gin.Context) {
//	id := c.Param("id")
//	//loop over listings to find one with requested id
//	for _, a := range listings {
//		if a.ID == id {
//			c.IndentedJSON(http.StatusOK, a)
//			return
//		}
//	}
//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "listing not found"})
//}

// function to add a listing
func addListing(c *gin.Context) {
	log.Printf("HELLO")
	ctx, client := controllers.InitialiseFirebaseApp()
	ref := client.NewRef("listings")

	var newListing models.Listing
	if err := c.BindJSON(&newListing); err != nil {
		return
	}
	if newListing.Swappable == true {
		newListing.Status = &models.ListingStatus{
			StatusId:   "",
			StatusName: "available",
			Enabled:    true,
			SortOrder:  0,
		}
	}
	log.Printf("STATUS %v", newListing.Status.StatusName)
	firebaseListingId, err := ref.Push(ctx, newListing)
	log.Printf("newListing %v", newListing)
	log.Printf("firebaseListingId %v", firebaseListingId)
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}

	c.IndentedJSON(http.StatusCreated, newListing)
}

func main() {
	controllers.InitialiseFirebaseApp()
	controllers.PopulateFirebase()

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/listings", getListings)
	// router.GET("/listings/:id", getListingById)
	router.POST("/listings", addListing)

	router.Run("0.0.0.0:8080")

}
