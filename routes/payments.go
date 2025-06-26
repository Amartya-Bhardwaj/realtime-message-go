package routes

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const razorpayBaseURL = "https://api.razorpay.com/v1/orders"

type RazorpayOrderRequest struct {
	Amount   int64             `json:"amount"`          // Amount in smallest currency unit (e.g., paise for INR)
	Currency string            `json:"currency"`        // E.g., "INR"
	Receipt  string            `json:"receipt"`         // Unique identifier for the order
	Notes    map[string]string `json:"notes,omitempty"` // Optional: Add additional info
}

type RazorpayOrderResponse struct {
	ID        string `json:"id"`
	Entity    string `json:"entity"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	Status    string `json:"status"`
	Receipt   string `json:"receipt"`
	CreatedAt int64  `json:"created_at"`
}

func CreateOrder(c *gin.Context) {
	apiKey := os.Getenv("razorpayKeyID")
	apiSecret := os.Getenv("razorpayKeySecret")
	razorpayOrderRequest := RazorpayOrderRequest{}
	if err := c.ShouldBindJSON(&razorpayOrderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if apiKey == "" || apiSecret == "" {
		log.Println("Razorpay API key and secret must be set as environment variables")
	}

	// Prepare request payload
	orderRequest := RazorpayOrderRequest{
		Amount:   razorpayOrderRequest.Amount,
		Currency: razorpayOrderRequest.Currency,
		Receipt:  razorpayOrderRequest.Receipt,
		Notes:    razorpayOrderRequest.Notes,
	}
	requestBody, err := json.Marshal(orderRequest)
	if err != nil {
		log.Println("failed to marshal request: ", err)
	}

	// Create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("POST", razorpayBaseURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("failed to create request: ", err)
	}

	// Add headers
	req.Header.Add("Content-Type", "application/json")
	auth := base64.StdEncoding.EncodeToString([]byte(apiKey + ":" + apiSecret))
	req.Header.Add("Authorization", "Basic "+auth)

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("request failed: ", err)
	}
	defer resp.Body.Close()

	// Read and parse response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("non-200 response from Razorpay: ", string(body))
		c.JSON(resp.StatusCode, string(body))
	}

	var orderResponse RazorpayOrderResponse
	if err := json.Unmarshal(body, &orderResponse); err != nil {
		log.Println("failed to parse response: ", err)
	}

	c.JSON(http.StatusOK, orderResponse)
}

func VerifyPayment(c *gin.Context) {

}

func OrderWebhookEvent(c *gin.Context) {
	var webHookEvent map[string]interface{}
	err := c.BindJSON(&webHookEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(webHookEvent)
}
