package main

import (
	"bytes"
	"encoding/json"
	"fionawaters/YarnSwap/controllers"
	"fionawaters/YarnSwap/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func SetUpTestRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetListings(t *testing.T) {
	router := SetUpTestRouter()

	router.GET("/listings", controllers.GetListings)
	req, _ := http.NewRequest("GET", "/listings", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var listings []models.Listing
	json.Unmarshal(w.Body.Bytes(), &listings)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, listings)

}

func TestAddListing(t *testing.T) {
	router := SetUpTestRouter()
	router.POST("/listings", controllers.AddListing)
	listingId := "test-listing-1"
	listing := models.Listing{
		ID:             listingId,
		UserId:         "test-userid",
		UserName:       "test-username",
		Brand:          "Green Elephant Yarn",
		Colourway:      "Standing Stone",
		Meterage:       425,
		Weight:         "sock",
		FibreContent:   "Wool",
		UnitWeight:     100,
		DyeLot:         "",
		Swappable:      true,
		Cost:           0,
		OriginalCount:  1,
		RemainingCount: 0,
		Timestamp:      time.Time{},
		ImageUrl:       "",
		Status:         "Available",
		ListingNote:    "",
	}
	jsonValue, _ := json.Marshal(listing)
	req, _ := http.NewRequest("POST", "/listings", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetBrands(t *testing.T) {
	router := SetUpTestRouter()

	router.GET("/brands", controllers.GetBrands)
	req, _ := http.NewRequest("GET", "/brands", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var brands []models.Brand
	json.Unmarshal(w.Body.Bytes(), &brands)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, brands)

}

func TestGetWeights(t *testing.T) {
	router := SetUpTestRouter()

	router.GET("/weights", controllers.GetWeights)
	req, _ := http.NewRequest("GET", "/weights", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var weights []models.Weight
	json.Unmarshal(w.Body.Bytes(), &weights)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, weights)

}

func TestGetFibres(t *testing.T) {
	router := SetUpTestRouter()

	router.GET("/fibres", controllers.GetFibreContents)
	req, _ := http.NewRequest("GET", "/fibres", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var fibres []models.FibreContent
	json.Unmarshal(w.Body.Bytes(), &fibres)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, fibres)

}

func TestGetListingStatuses(t *testing.T) {
	router := SetUpTestRouter()

	router.GET("/listing-statuses", controllers.GetListingStatuses)
	req, _ := http.NewRequest("GET", "/listing-statuses", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var listingStatuses []models.ListingStatus
	json.Unmarshal(w.Body.Bytes(), &listingStatuses)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, listingStatuses)

}

func TestAddSwap(t *testing.T) {
	router := SetUpTestRouter()
	router.POST("/swaps", controllers.AddSwap)
	swapId := "test-swapId"
	swap := models.Swap{
		ID:             swapId,
		SwapName:       "green elephant red yarn swap chat",
		SwapperUserID:  "test-swapperuserid",
		SwappeeUserID:  "test-swappeeuserid",
		ListingID:      "acb123",
		SwapStatus:     "swap requested",
		ChatChannelUrl: "",
		SwapNote:       "",
	}
	jsonValue, _ := json.Marshal(swap)
	req, _ := http.NewRequest("POST", "/swaps", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetSwaps(t *testing.T) {
	router := SetUpTestRouter()

	router.GET("/swaps", controllers.GetSwaps)
	req, _ := http.NewRequest("GET", "/swaps", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var swaps []models.Swap
	json.Unmarshal(w.Body.Bytes(), &swaps)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, swaps)

}

func TestAddUserDetails(t *testing.T) {
	router := SetUpTestRouter()
	router.POST("/users", controllers.AddUserDetails)
	userId := "test-user-1"
	user := models.User{
		ID:                userId,
		UserName:          "test-username",
		RemainingTokens:   0,
		AccountStatus:     "Active",
		AmtListingsAdded:  0,
		AmtSwapsCompleted: 0,
		CreationTimestamp: time.Time{},
		ArchiveTimestamp:  time.Time{},
		Role:              "user",
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetUsers(t *testing.T) {
	router := SetUpTestRouter()

	router.GET("/users", controllers.GetUsers)
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)

}

func TestGetUserProfile(t *testing.T) {
	router := SetUpTestRouter()

	router.GET("/user/:id", controllers.GetUserProfile)
	user := models.User{
		ID:                "test-user-1",
		UserName:          "test-username",
		RemainingTokens:   0,
		AccountStatus:     "Active",
		AmtListingsAdded:  0,
		AmtSwapsCompleted: 0,
		CreationTimestamp: time.Time{},
		ArchiveTimestamp:  time.Time{},
		Role:              "user",
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("GET", "/user/"+user.ID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, user)

}
