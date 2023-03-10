package main

import (
	controllers "fionawaters/YarnSwap/controllers"
	"fionawaters/YarnSwap/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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
		listing := getListingById(s.ListingID)
		//add new struct to array
		swapListing := models.SwapListing{Swap: s, Listing: listing}
		data[i] = swapListing

	}
	//log.Default().Println("data = ", data)
	c.IndentedJSON(http.StatusOK, data)
}

func getListingById(listingId string) models.Listing {
	ctx, client, _ := controllers.InitialiseFirebaseApp()

	ref := client.NewRef("listings")
	var listing models.Listing
	err := ref.Child(listingId).Get(ctx, &listing)
	if err != nil {
		log.Fatalln("Error getting listing:", err)
	}

	return listing
}

func getUsers(c *gin.Context) {
	ctx, client, _ := controllers.InitialiseFirebaseApp()

	//Create Ref for users
	ref := client.NewRef("users")
	//retrieve the users in order of the keys
	results, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	//create an array the same length as the number of results
	data := make([]models.User, len(results))

	//loop over the results and individually marshal into Listing struct
	for i, r := range results {
		var u models.User

		if err := r.Unmarshal(&u); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		u.ID = r.Key()
		//add new struct to array
		data[i] = u
	}

	log.Default().Println("data getUsersssss = ", data)

	c.IndentedJSON(http.StatusOK, data)
}

func getUserProfile(c *gin.Context) {
	id := c.Param("id")
	ctx, client, _ := controllers.InitialiseFirebaseApp()

	ref := client.NewRef("users")
	var user models.User

	err := ref.Child(id).Get(ctx, &user)
	if err != nil {
		c.AbortWithStatus(404)
	}

	if user.UserName == "" {
		c.AbortWithStatus(404)
	} else {
		c.IndentedJSON(http.StatusOK, user)
	}

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
		log.Printf("timestamp %v", newListing.Timestamp)
		_, err := ref.Push(ctx, newListing)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
		// increase the users amount of listings added variable by 1
		user.AmtListingsAdded++
		err = userRef.Update(ctx, map[string]interface{}{newListing.UserId: user})
		if err != nil {
			log.Printf("error updating user")
		}
	}
	c.IndentedJSON(http.StatusCreated, newListing)
}

func addUserDetails(c *gin.Context) {
	ctx, client, _ := controllers.InitialiseFirebaseApp()

	ref := client.NewRef("users")

	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	// if the user has an id - update the existing record
	if newUser.ID != "" {
		var id = newUser.ID
		newUser.ID = ""
		if newUser.AccountStatus == "Archived" {
			newUser.ArchiveTimestamp = time.Now()
			listingsRef := client.NewRef("listings")
			results, err := listingsRef.OrderByKey().GetOrdered(ctx)
			if err != nil {
				log.Fatalln("Error querying database:", err)
			}
			for _, v := range results {
				var l models.Listing
				if err := v.Unmarshal(&l); err != nil {
					log.Fatalln("Error unmarshaling result", err)
				}
				if l.UserId == id {
					l.Status = "Archived"

					err = listingsRef.Update(ctx, map[string]interface{}{v.Key(): l})
					{
						if err != nil {
							log.Printf("error updating listings %v", err)
						}
					}
				}
			}
		}
		if newUser.AccountStatus == "Active" {
			listingsRef := client.NewRef("listings")
			results, err := listingsRef.OrderByKey().GetOrdered(ctx)
			if err != nil {
				log.Fatalln("Error querying database:", err)
			}

			for _, v := range results {
				var l models.Listing
				if err := v.Unmarshal(&l); err != nil {
					log.Fatalln("Error unmarshaling result", err)
				}

				if l.UserId == id {
					l.Status = "Available"
					err = listingsRef.Update(ctx, map[string]interface{}{v.Key(): l})
					if err != nil {
						log.Printf("error updating listings %v", err)
					}
				}
			}
		}
		err := ref.Update(ctx, map[string]interface{}{id: newUser})
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	} else {
		// else create a new record
		newUser.CreationTimestamp = time.Now()
		newUser.Role = "user"
		log.Printf("time %v", time.Now())
		_, err := ref.Push(ctx, newUser)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	}
	c.IndentedJSON(http.StatusCreated, newUser)
}

// TODO how to implement this?
func deleteUser(c *gin.Context) {
	ctx, client, app := controllers.InitialiseFirebaseApp()
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing auth client: %v\n", err)
	}

	ref := client.NewRef("users")

	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	if newUser.ID != "" {
		var id = newUser.ID
		newUser.ID = ""
		if newUser.AccountStatus == "Archived" && newUser.ArchiveTimestamp.Sub(time.Now()) >= 720 {
			log.Printf("timestamp %v", newUser.ArchiveTimestamp.Sub(time.Now()))
			// TODO delete this users listings & swaps?
			newUser.AccountStatus = "Deleted"
		}
		err := ref.Update(ctx, map[string]interface{}{id: newUser})
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
		// check if it has been 30 days since the user has been archived - if it has delete the user
		// also delete their listings

		err = authClient.DeleteUser(ctx, id)
		if err != nil {
			log.Printf("error deleting user: %v", err)
		}
		log.Printf("successfully deleted user: %s ", id)
	}

}

// function to add a swap
func addSwap(c *gin.Context) {
	ctx, client, _ := controllers.InitialiseFirebaseApp()
	ref := client.NewRef("swaps")
	var newSwap models.Swap
	if err := c.BindJSON(&newSwap); err != nil {
		log.Printf("error binding: %v\n", err)
		return
	}

	// if the swap has an ID (i.e. it already exists) update it
	if newSwap.ID != "" {
		var id = newSwap.ID
		newSwap.ID = ""
		// if a swap has been completed add 1 'AmtSwapsCompleted' to each user involved in the swap
		if newSwap.SwapStatus == "Complete" {
			userRef := client.NewRef("users")
			// increase the users amount of listings added variable by 1
			var swapper models.User
			err := userRef.Child(newSwap.SwapperUserID).Get(ctx, &swapper)
			if err != nil {
				log.Printf("error getting user: swapper %v", err)
			}
			swapper.AmtSwapsCompleted++
			err2 := userRef.Update(ctx, map[string]interface{}{newSwap.SwapperUserID: swapper})
			if err != nil {
				log.Printf("error updating user: swapper %v", err2)
			}
			var swappee models.User
			err3 := userRef.Child(newSwap.SwappeeUserID).Get(ctx, &swappee)
			if err != nil {
				log.Printf("error getting user: swappee %v", err3)
			}
			swappee.AmtSwapsCompleted++
			// if swap is completed take 1 token from swappee
			swappee.RemainingTokens--
			err4 := userRef.Update(ctx, map[string]interface{}{newSwap.SwappeeUserID: swappee})
			if err != nil {
				log.Printf("error updating user: swappee %v", err4)
			}

		}
		err := ref.Update(ctx, map[string]interface{}{id: newSwap})
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	} else {
		// create a new swap
		_, err := ref.Push(ctx, newSwap)
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
	}

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
	ctx, client, _ := controllers.InitialiseFirebaseApp()
	controllers.PopulateFirebase(ctx, client)

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
	router.GET("/swaps", getSwaps)
	router.POST("/users", authMiddleware, addUserDetails)
	//router.GET("/users", getUsers)
	router.GET("/user/:id", getUserProfile)
	router.Run("0.0.0.0:8080")

}

//TODO wishlist functions
