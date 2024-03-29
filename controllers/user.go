package controllers

import (
	"fionawaters/YarnSwap/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// GetUsers function to get users from firebase
func GetUsers(c *gin.Context) {
	ctx, client, _ := InitialiseFirebaseApp()

	//Create Ref for users
	ref := client.NewRef("users")
	//retrieve the users in order of the keys
	results, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	//create an array the same length as the number of results
	data := make([]models.User, len(results))

	//loop over the results and individually marshal into user struct
	for i, r := range results {
		var u models.User

		if err := r.Unmarshal(&u); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		u.ID = r.Key()
		//add new struct to array
		data[i] = u
	}

	//log.Default().Println("data getUsers = ", data)

	c.IndentedJSON(http.StatusOK, data)
}

// GetUserProfile function to get an individual user profile from firebase
func GetUserProfile(c *gin.Context) {
	id := c.Param("id")
	ctx, client, _ := InitialiseFirebaseApp()

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

func DeleteTestUserProfile(userId string) error {
	ctx, client, _ := InitialiseFirebaseApp()

	userRef := client.NewRef("users")
	err := userRef.Child(userId).Delete(ctx)
	if err != nil {
		log.Fatalln("error deleting test user profile, please delete manually")
	}
	return nil
}

// AddUserDetails function to add/update user details
func AddUserDetails(c *gin.Context) {
	ctx, client, _ := InitialiseFirebaseApp()

	ref := client.NewRef("users")

	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	var id = newUser.ID
	newUser.ID = ""
	uniqueUsername := isUsernameUnique(id, newUser.UserName)
	log.Printf("uniqueusername %v", uniqueUsername)

	existingUser := GetUserById(id)
	if uniqueUsername {
		newUser.CreationTimestamp = time.Now()
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
		if newUser.AccountStatus == "Active" && existingUser.AccountStatus != "Active" {
			newUser.Role = "user"
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
		//log.Printf("timestamp difference %v", newUser.ArchiveTimestamp.Sub(time.Now()))
		err := ref.Update(ctx, map[string]interface{}{id: newUser})
		if err != nil {
			log.Fatalln("Error setting value:", err)
		}
		c.IndentedJSON(http.StatusCreated, newUser)
	} else {
		c.AbortWithStatus(400)
	}
}

//TODO create a cron job that utilises DeleteUser function

// DeleteUser function to delete user
// including user listings, user swaps + firebase auth record
func DeleteUser(userId string) error {
	ctx, client, app := InitialiseFirebaseApp()
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing auth client: %v\n", err)
	}

	ref := client.NewRef("users")

	var newUser models.User
	err = ref.Child(userId).Get(ctx, &newUser)
	if err != nil {
		log.Fatalln("Error getting user: ", err)
	}
	// check if it has been 30 days(in nanoseconds) since the user has been archived - mark the user record as deleted
	if newUser.AccountStatus == "Archived" && newUser.ArchiveTimestamp.Sub(time.Now()) >= time.Duration(2.592e+15) {
		log.Printf("timestamp %v", newUser.ArchiveTimestamp.Sub(time.Now()))
		newUser.AccountStatus = "Deleted"
		err := DeleteUserListings(userId)
		if err != nil {
			log.Fatalln("Error deleting user listings ", err)
		}
		err = DeleteUserSwaps(userId)
		if err != nil {
			log.Fatalln("Error deleting user swaps ", err)
		}
	}
	if newUser.AccountStatus == "Active" {
		newUser.ArchiveTimestamp = time.Time{}
	}
	err = ref.Update(ctx, map[string]interface{}{userId: newUser})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	//delete user from firebase auth
	err = authClient.DeleteUser(ctx, userId)
	if err != nil {
		log.Printf("error deleting user: %v", err)
	}
	log.Printf("successfully deleted user: %s ", userId)

	return nil
}

func isUsernameUnique(id string, userName string) bool {
	//get all user profiles
	ctx, client, _ := InitialiseFirebaseApp()

	//Create Ref for users
	ref := client.NewRef("users")
	//retrieve the users in order of the keys
	results, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	// loop through them and check the username against the profile usernames
	var res = true
	for _, r := range results {
		var u models.User
		if err := r.Unmarshal(&u); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		u.ID = r.Key()
		// if suggested username matches existing username return false
		if u.UserName == userName && u.ID != id {
			res = false
			break
		}
	}
	return res
}

func GetUserById(userId string) models.User {
	ctx, client, _ := InitialiseFirebaseApp()

	ref := client.NewRef("users")
	var user models.User
	err := ref.Child(userId).Get(ctx, &user)
	if err != nil {
		log.Fatalln("Error getting user:", err)
	}

	return user
}
