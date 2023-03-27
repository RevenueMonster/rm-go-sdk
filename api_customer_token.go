package sdk

import (
	"errors"
	"strings"
	"time"
)

type CustomerToken struct {
	ID        string    `json:"id"`
	Label     string    `json:"label"`
	Provider  string    `json:"provider"`
	Token     string    `json:"token"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Address   struct {
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		PostCode     string `json:"postCode"`
		City         string `json:"city"`
		State        string `json:"state"`
		Country      string `json:"country"`
	} `json:"address"`
	Card struct {
		Name        string `json:"name"`
		Method      string `json:"method"`
		Fundingtype string `json:"fundingtype"`
		LastFourNo  string `json:"lastFourNo"`
		ExpMonth    int    `json:"expMonth"`
		ExpYear     int    `json:"ExpYear"`
		IsCvcCheck  bool   `json:"isCvcCheck"`
	} `json:"card"`
}

type RequestGetPaymentCheckoutCustomerToken struct {
	CustomerID string
}

type ResponseGetPaymentCheckoutCustomerToken struct {
	Item []CustomerToken `json:"item"`
	Code string          `json:"code"`
	Err  *Error          `json:"error"`
}

// GetPaymentCheckoutCustomerToken :
func (c Client) GetPaymentCheckoutCustomerToken(request RequestGetPaymentCheckoutCustomerToken) (*ResponseGetPaymentCheckoutCustomerToken, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetPaymentCheckoutCustomerToken.method
	requestURL := c.prepareAPIURL(pathAPIGetPaymentCheckoutCustomerToken)
	requestURL = strings.ReplaceAll(requestURL, "{customer_id}", request.CustomerID)

	response := new(ResponseGetPaymentCheckoutCustomerToken)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

type RequestDeletePaymentCheckoutCustomerToken struct {
	CustomerID string
	Token      string
}
type ResponseDeletePaymentCheckoutCustomerToken struct {
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// DeletePaymentCheckoutCustomerToken :
func (c Client) DeletePaymentCheckoutCustomerToken(request RequestDeletePaymentCheckoutCustomerToken) error {
	if c.err != nil {
		return c.err
	}

	method := pathAPIDeletePaymentCheckoutCustomerToken.method
	requestURL := c.prepareAPIURL(pathAPIDeletePaymentCheckoutCustomerToken)
	requestURL = strings.ReplaceAll(requestURL, "{customer_id}", request.CustomerID)

	response := new(ResponseGetPaymentCheckoutCustomerToken)
	if err := c.httpAPI(method, requestURL, map[string]string{"token": request.Token}, response); err != nil {
		return err
	}

	if response.Err != nil {
		return errors.New(response.Err.Message)
	}

	return nil
}
