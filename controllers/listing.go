package controllers

import (
	"fionawaters/YarnSwap/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// GetListings function to retrieve listings from firebase realtime database
func GetListings(c *gin.Context) {
	ctx, client, _ := InitialiseFirebaseApp()

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
		l.ID = r.Key()
		//add new struct to array
		data[i] = l
	}

	//log.Default().Println("data = ", data)

	c.IndentedJSON(http.StatusOK, data)
}

// GetListingById function to retrieve a listing by ID
func GetListingById(listingId string) models.Listing {
	ctx, client, _ := InitialiseFirebaseApp()

	ref := client.NewRef("listings")
	var listing models.Listing
	err := ref.Child(listingId).Get(ctx, &listing)
	if err != nil {
		log.Fatalln("Error getting listing:", err)
	}

	return listing
}

// AddListing function to add a listing
func AddListing(c *gin.Context) {
	ctx, client, _ := InitialiseFirebaseApp()
	ref := client.NewRef("listings")

	var newListing models.Listing
	if err := c.BindJSON(&newListing); err != nil {
		return
	}

	userRef := client.NewRef("users")
	var user models.User
	err := userRef.Child(newListing.UserId).Get(ctx, &user)
	if err != nil {
		log.Printf("error getting user")
	}

	// if the listing has an ID (i.e. it already exists) update it
	if newListing.ID != "" {
		var id = newListing.ID
		newListing.ID = ""
		err := ref.Update(ctx, map[string]interface{}{id: newListing})
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	} else {
		// create a new listing
		newListing.Timestamp = time.Now()
		// set the listing owners username
		newListing.UserName = user.UserName
		if newListing.Swappable == true {
			newListing.Status = "Awaiting approval"
		} else {
			newListing.Status = "Available"
		}
		// increase the users amount of listings added variable by 1
		user.AmtListingsAdded++
		log.Printf("timestamp %v", newListing.Timestamp)
		_, err := ref.Push(ctx, newListing)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
		err = userRef.Update(ctx, map[string]interface{}{newListing.UserId: user})
		if err != nil {
			log.Printf("error updating user")
		}
	}
	c.IndentedJSON(http.StatusCreated, newListing)
}

func DeleteUserListings(userId string) error {
	ctx, client, _ := InitialiseFirebaseApp()

	listingsRef := client.NewRef("listings")

	//retrieve the listings in order of the keys
	results, err := listingsRef.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	log.Printf("results %v:", results)

	//loop over the results and individually marshal into Listing struct
	for _, r := range results {
		var l models.Listing
		if err := r.Unmarshal(&l); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		// get listings that match this userId
		if l.UserId == userId {
			l.ID = r.Key()
			// delete the listings
			err := listingsRef.Child(l.ID).Delete(ctx)
			if err != nil {
				log.Fatalln("Error deleting user listings")
			}
		}
	}
	return nil
}

// GetBrands function to get brands from firebase realtime database to populate dropdown in add listing form
func GetBrands(c *gin.Context) {
	ctx, client, _ := InitialiseFirebaseApp()

	//Create Ref for brands
	ref := client.NewRef("brands")
	//retrieve the listings in order of the keys
	results, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	//create an array the same length as the number of results
	data := make([]models.Brand, len(results))

	//loop over the results and individually marshal into Listing struct
	for i, r := range results {
		var b models.Brand
		if e := r.Unmarshal(&b); e != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		//add new struct to array
		data[i] = b
	}
	//log.Default().Println("data = ", data)

	c.IndentedJSON(http.StatusOK, data)
}

// GetWeights function to get weights from firebase realtime database to populate dropdown in add listing form
func GetWeights(c *gin.Context) {
	ctx, client, _ := InitialiseFirebaseApp()

	//Create Ref for weights
	ref := client.NewRef("weights")
	//retrieve the listings in order of the keys
	results, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	//create an array the same length as the number of results
	data := make([]models.Weight, len(results))

	//loop over the results and individually marshal into Listing struct
	for i, r := range results {
		var w models.Weight
		if e := r.Unmarshal(&w); e != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		//add new struct to array
		data[i] = w
	}
	//	log.Default().Println("data = ", data)

	c.IndentedJSON(http.StatusOK, data)
}

// GetFibreContents function to get fibre contents from firebase realtime database to populate dropdown in add listing form
func GetFibreContents(c *gin.Context) {
	ctx, client, _ := InitialiseFirebaseApp()

	//Create Ref for fibre contents
	ref := client.NewRef("fibres")
	//retrieve the listings in order of the keys
	results, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	//create an array the same length as the number of results
	data := make([]models.FibreContent, len(results))

	//loop over the results and individually marshal into Listing struct
	for i, r := range results {
		var f models.FibreContent
		if e := r.Unmarshal(&f); e != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		//add new struct to array
		data[i] = f
	}
	//	log.Default().Println("data = ", data)

	c.IndentedJSON(http.StatusOK, data)
}

// GetListingStatuses function to get listing statuses from firebase realtime database to populate dropdown for filter functionality
func GetListingStatuses(c *gin.Context) {
	ctx, client, _ := InitialiseFirebaseApp()

	//Create Ref for fibre contents
	ref := client.NewRef("listing-status")
	//retrieve the listings in order of the keys
	results, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	//create an array the same length as the number of results
	data := make([]models.ListingStatus, len(results))

	//loop over the results and individually marshal into Listing struct
	for i, r := range results {
		var f models.ListingStatus
		if e := r.Unmarshal(&f); e != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		//add new struct to array
		data[i] = f
	}
	//	log.Default().Println("data = ", data)

	c.IndentedJSON(http.StatusOK, data)
}
