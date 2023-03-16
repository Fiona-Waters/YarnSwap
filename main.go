package main

import (
	controllers "fionawaters/YarnSwap/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

// function to check ID token before allowing access
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

	router.GET("/listings", controllers.GetListings)
	router.POST("/listings", authMiddleware, controllers.AddListing)
	router.GET("/brands", controllers.GetBrands)
	router.GET("/weights", controllers.GetWeights)
	router.GET("/fibres", controllers.GetFibreContents)
	router.GET("/listing-statuses", controllers.GetListingStatuses)
	router.POST("/swaps", authMiddleware, controllers.AddSwap)
	router.GET("/swaps", controllers.GetSwaps)
	router.POST("/users", authMiddleware, controllers.AddUserDetails)
	router.GET("/users", controllers.GetUsers)
	router.GET("/user/:id", authMiddleware, controllers.GetUserProfile)
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		return
	}

}

//TODO wishlist functions
