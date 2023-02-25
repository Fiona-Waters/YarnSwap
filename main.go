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
	ctx, client, _ := controllers.InitialiseFirebaseApp()

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

func getSwaps(c *gin.Context) {
	ctx, client, _ := controllers.InitialiseFirebaseApp()

	//Create Ref for swaps
	ref := client.NewRef("swaps")
	//retrieve the swaps in order of the keys
	results, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	//create an array the same length as the number of results
	data := make([]models.SwapListing, len(results))

	//loop over the results and individually marshal into Swap struct
	for i, r := range results {
		var s models.Swap

		if err := r.Unmarshal(&s); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		s.ID = r.Key()
		listing := getListingById(c, s.ListingID)
		//add new struct to array
		if listing.ID == s.ListingID {
			swapListing := models.SwapListing{Swap: s, Listing: listing}
			log.Printf("swapListing %v", swapListing)
			data[i] = swapListing
		}
	}
	log.Default().Println("data = ", data)
	c.IndentedJSON(http.StatusOK, data)
}

//TODO create a func to get Listing by ID
func getListingById(c *gin.Context, listingId string) models.Listing {
	ctx, client, _ := controllers.InitialiseFirebaseApp()

	ref := client.NewRef("listings")
	//retrieve the listings in order of the keys
	results, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	//create an array the same length as the number of results
	data := make([]models.Listing, len(results))
	var l models.Listing
	//loop over the results and individually marshal into Listing struct
	for i, r := range results {
		//var l models.Listing
		if err := r.Unmarshal(&l); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		data[i] = l
	}
	c.IndentedJSON(http.StatusOK, data)
	return l
}

func getBrands(c *gin.Context) {
	ctx, client, _ := controllers.InitialiseFirebaseApp()

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

func getWeights(c *gin.Context) {
	ctx, client, _ := controllers.InitialiseFirebaseApp()

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

func getFibreContents(c *gin.Context) {
	ctx, client, _ := controllers.InitialiseFirebaseApp()

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

// function to add a listing
func addListing(c *gin.Context) {
	ctx, client, _ := controllers.InitialiseFirebaseApp()
	ref := client.NewRef("listings")

	var newListing models.Listing
	if err := c.BindJSON(&newListing); err != nil {
		return
	}
	if newListing.Swappable == true {
		newListing.Status = "Available"
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
		_, err := ref.Push(ctx, newListing)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	}

	c.IndentedJSON(http.StatusCreated, newListing)
}

// function to add a swap
func addSwap(c *gin.Context) {
	ctx, client, _ := controllers.InitialiseFirebaseApp()
	ref := client.NewRef("swaps")
	log.Println("hello")
	var newSwap models.Swap
	if err := c.BindJSON(&newSwap); err != nil {
		log.Printf("error binding: %v\n", err)
		return
	}
	log.Printf("newswap %v", newSwap)

	// if the swap has an ID (i.e. it already exists) update it
	//if newSwap.SwapID != "" {
	//	err := ref.Update(ctx, map[string]interface{}{newSwap.SwapID: newSwap})
	//	if err != nil {
	//		log.Fatalln("Error setting value:", err)
	//	}
	//} else {
	// create a new swap
	_, err := ref.Push(ctx, newSwap)
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	//}

	c.IndentedJSON(http.StatusCreated, newSwap)
}

func authMiddleware(c *gin.Context) {
	ctx, _, app := controllers.InitialiseFirebaseApp()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Printf("error getting Auth client: %v\n", err)
	}
	idToken := c.Request.Header.Get("X-ID-TOKEN")
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		c.AbortWithStatus(401)
	}
	if token != nil {
		c.Next()
	} else {
		c.AbortWithStatus(401)
		log.Printf("user not authorised: %v\n", err)
	}
}

func main() {
	controllers.InitialiseFirebaseApp()
	controllers.PopulateFirebase()

	router := gin.Default()
	//router.Use(cors.Default())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "X-ID-TOKEN")
	router.Use(cors.New(corsConfig))

	router.GET("/listings", getListings)
	router.POST("/listings", authMiddleware, addListing)
	router.GET("/brands", getBrands)
	router.GET("/weights", getWeights)
	router.GET("/fibres", getFibreContents)
	router.POST("/swaps", authMiddleware, addSwap)
	router.GET("swaps", getSwaps)
	router.Run("0.0.0.0:8080")

}

//TODO addSwap and getSwaps functions
//TODO wishlist functions
