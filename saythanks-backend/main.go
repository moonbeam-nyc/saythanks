package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var userID string

func main() {
	// Get USPS user ID from environment variable
	userID := os.Getenv("USPS_USER_ID")
	if userID == "" {
		fmt.Println("USPS_USER_ID environment variable is required")
		return
	}

	// Create and run the Gin server
	router := gin.Default()
	router.POST("/api/address/validate", validateAddress)
	router.GET("/api/recipients", getRecipients)
	router.Run(":8080")
}

func validateAddress(c *gin.Context) {
	var request struct {
		Address string `json:"address"`
		ZipCode string `json:"zip_code"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"API": "Verify",
			"XML": `<AddressValidateRequest USERID="` + userID + `">
                        <Revision>1</Revision>
						<Address>
							<Address1></Address1>
							<Address2>` + request.Address + `</Address2>
							<City></City>
							<State></State>
							<Zip5>` + request.ZipCode + `</Zip5>
							<Zip4></Zip4>
						</Address>
					</AddressValidateRequest>`,
		}).
		Get("https://secure.shippingapis.com/ShippingAPI.dll")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp.String()})
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
