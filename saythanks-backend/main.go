package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var (
	token     string
	tokenLock sync.Mutex
	expiry    time.Time
)

func main() {
	// Initialize the router
	router := gin.Default()

	// Set up CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/api/address/validate", validateAddress)
	router.GET("/api/recipients", getRecipients)
	router.Run(":8080")
}

func getAccessToken() (string, error) {
	tokenLock.Lock()
	defer tokenLock.Unlock()

	if token != "" && time.Now().Before(expiry) {
		return token, nil
	}

	clientID := os.Getenv("USPS_CLIENT_ID")
	clientSecret := os.Getenv("USPS_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		return "", fmt.Errorf("USPS_CLIENT_ID or USPS_CLIENT_SECRET environment variable not set")
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"client_id":     clientID,
			"client_secret": clientSecret,
			"grant_type":    "client_credentials",
		}).
		Post("https://api.usps.com/oauth2/v3/token")

	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", err
	}

	if accessToken, ok := result["access_token"].(string); ok {
		token = accessToken
		expiresIn := int64(result["expires_in"].(float64))
		expiry = time.Now().Add(time.Duration(expiresIn) * time.Second)
		return token, nil
	}

	return "", fmt.Errorf("failed to get access token")
}

func validateAddress(c *gin.Context) {
	var request struct {
		Address string `json:"address"`
		City    string `json:"city"`
		State   string `json:"state"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := getAccessToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := resty.New()
	resp, err := client.R().
		SetAuthToken(token).
		SetQueryParams(map[string]string{
			"streetAddress": request.Address,
			"city":          request.City,
			"state":         request.State,
		}).
		Get("https://api.usps.com/addresses/v3/address")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal([]byte(resp.String()), &responseData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responseData})
}

func getRecipients(c *gin.Context) {
	address := c.Query("address")
	zipCode := c.Query("zip_code")

	if address == "" || zipCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address and zip_code query parameters are required"})
		return
	}

	// Implement your logic to get recipients based on address and zip code
	recipients := []string{"Recipient 1", "Recipient 2"} // Placeholder

	c.JSON(http.StatusOK, gin.H{"recipients": recipients})
}
