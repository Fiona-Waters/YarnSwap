package controllers

import (
	"fionawaters/YarnSwap/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetSwaps function to get swaps
func GetSwaps(c *gin.Context) {
	ctx, client, _ := InitialiseFirebaseApp()

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
		listing := GetListingById(s.ListingID)
		//add new struct to array
		swapListing := models.SwapListing{Swap: s, Listing: listing}
		data[i] = swapListing

	}
	//log.Default().Println("data = ", data)
	c.IndentedJSON(http.StatusOK, data)
}

// AddSwap function to add a swap
func AddSwap(c *gin.Context) {
	ctx, client, _ := InitialiseFirebaseApp()
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

func DeleteUserSwaps(userId string) error {
	ctx, client, _ := InitialiseFirebaseApp()

	swapsRef := client.NewRef("swaps")

	//retrieve the listings in order of the keys
	results, err := swapsRef.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	//loop over the results and individually marshal into Listing struct
	for _, r := range results {
		var s models.Swap
		if err := r.Unmarshal(&s); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		// get listings that match this userId
		if s.SwapperUserID == userId {
			s.ID = r.Key()
			err := swapsRef.Child(s.ID).Delete(ctx)
			if err != nil {
				log.Fatalln("Error deleting user swaps")
			}
		}
	}
	return nil
}
