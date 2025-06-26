package views

import (
	"encoding/json"
	"errors"
)

func PaymentService(body string) (bool, error){
	status := false
	type paymentRequestS struct {
		OrderId           string `json:"orderId"`
		MerchantPaymentId string `json:"merchantPaymentId"`
		MerchantSignature string `json:"merchantSignature"`
	}
	paymentRequest := paymentRequestS{}
	er := json.Unmarshal([]byte(body), &paymentRequest)
	if er != nil {
		return status, er
	}
	if paymentRequest.OrderId == "" || paymentRequest.MerchantPaymentId == "" || paymentRequest.MerchantSignature == "" {
		return status, errors.New("Error Occured")
	}
	

	return status,er
}