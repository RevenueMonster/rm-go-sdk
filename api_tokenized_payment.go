package sdk

import (
	"errors"
	"strings"
)

// --- models

type TokenizedCustomer struct {
	ID                 string `json:"id"`
	MerchantID         string `json:"merchantId"`
	StoreID            string `json:"storeId"`
	Label              string `json:"label"`
	Email              string `json:"email"`
	Name               string `json:"name"`
	CountryCode        string `json:"countryCode"`
	PhoneNumber        string `json:"phoneNumber"`
	ProductName        string `json:"productName"`
	ProductDescription string `json:"productDescription"`
	CreatedDateTime    string `json:"createdAt"`
	UpdatedDateTime    string `json:"updatedAt"`
	ClientKey          string `json:"clientKey"`
	RedirectURL        string `json:"redirectUrl"`
	NotifyURL          string `json:"notifyUrl"`
	PaymentURL         string `json:"paymentUrl,omitempty"`
	IsActive           bool   `json:"isActive"`
}

// --- sdk

type RequestCreateTokenizedPaymentCustomer struct {
	StoreID            string `json:"storeID"`
	Email              string `json:"email"`
	RedirectURL        string `json:"redirectUrl"`
	NotifyURL          string `json:"notifyUrl"`
	ProductName        string `json:"productName"`
	ProductDescription string `json:"productDescription"`
	Name               string `json:"name"`
	CountryCode        string `json:"countryCode"`
	PhoneNumber        string `json:"phoneNumber"`
}

type ResponseCreateTokenizedPaymentCustomer struct {
	Item *TokenizedCustomer `json:"item"`
	Code string             `json:"code"`
	Err  *Error             `json:"error"`
}

// CreateTokenizedPaymentCustomer :
func (c Client) CreateTokenizedPaymentCustomer(request RequestCreateTokenizedPaymentCustomer) (*ResponseCreateTokenizedPaymentCustomer, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPICreateTokenizedPaymentCustomer.method
	requestURL := c.prepareAPIURL(pathAPICreateTokenizedPaymentCustomer)
	response := new(ResponseCreateTokenizedPaymentCustomer)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

type RequestGetTokenizedPaymentCustomerByID struct {
	TokenizedCustomerID string `json:"customerId"`
}

type ResponseGetTokenizedPaymentCustomerByID struct {
	Item *TokenizedCustomer `json:"item"`
	Code string             `json:"code"`
	Err  *Error             `json:"error"`
}

// ToggleTokenizedPaymentCustomerStatus :
func (c Client) GetTokenizedPaymentCustomerByID(request RequestGetTokenizedPaymentCustomerByID) (*ResponseGetTokenizedPaymentCustomerByID, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetTokenizedPaymentCustomerByID.method
	requestURL := c.prepareAPIURL(pathAPIGetTokenizedPaymentCustomerByID)
	requestURL = strings.ReplaceAll(requestURL, "{customer_id}", request.TokenizedCustomerID)

	response := new(ResponseGetTokenizedPaymentCustomerByID)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

type RequestToggleTokenizedPaymentCustomerStatus struct {
	TokenizedCustomerID string
}

type ResponseToggleTokenizedPaymentCustomerStatus struct {
	Item *TokenizedCustomer `json:"item"`
	Code string             `json:"code"`
	Err  *Error             `json:"error"`
}

// ToggleTokenizedPaymentCustomerStatus :
func (c Client) ToggleTokenizedPaymentCustomerStatus(request RequestToggleTokenizedPaymentCustomerStatus) (*ResponseToggleTokenizedPaymentCustomerStatus, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIToggleTokenizedPaymentCustomerStatus.method
	requestURL := c.prepareAPIURL(pathAPIToggleTokenizedPaymentCustomerStatus)
	requestURL = strings.ReplaceAll(requestURL, "{customer_id}", request.TokenizedCustomerID)

	response := new(ResponseToggleTokenizedPaymentCustomerStatus)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

type RequestCreateOrderWithTokenizedCustomer struct {
	TokenizedCustomerID string `json:"-"`
	Currency            string `json:"currency"`
	Amount              int64  `json:"amount"`
	Title               string `json:"title"`
	Description         string `json:"description"`
}

type ResponseCreateOrderWithTokenizedCustomer struct {
	Item *PaymentTransaction `json:"item"`
	Code string              `json:"code"`
	Err  *Error              `json:"error"`
}

// ToggleTokenizedPaymentCustomerStatus :
func (c Client) CreateOrderWithTokenizedCustomer(request RequestCreateOrderWithTokenizedCustomer) (*ResponseCreateOrderWithTokenizedCustomer, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPICreateOrderWithTokenizedCustomer.method
	requestURL := c.prepareAPIURL(pathAPICreateOrderWithTokenizedCustomer)
	requestURL = strings.ReplaceAll(requestURL, "{customer_id}", request.TokenizedCustomerID)

	response := new(ResponseCreateOrderWithTokenizedCustomer)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

type RequestGetTokenizedCustomerOrders struct {
	TokenizedCustomerID string
	Cursor              string
}
type ResponseGetTokenizedCustomerOrders struct {
	Items []*TokenizedPaymentOrder `json:"items"`
	Code  string                   `json:"code"`
	Meta  *struct {
		Cursor string `json:"cursor"`
	} `json:"meta"`
	Err *Error `json:"error"`
}

func (c Client) GetTokenizedCustomerOrders(request RequestGetTokenizedCustomerOrders) (*ResponseGetTokenizedCustomerOrders, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetTokenizedCustomerOrders.method
	requestURL := c.prepareAPIURL(pathAPIGetTokenizedCustomerOrders)
	requestURL = strings.ReplaceAll(requestURL, "{customer_id}", request.TokenizedCustomerID)

	if request.Cursor != "" {
		requestURL += "?cursor=" + request.Cursor
	}

	response := new(ResponseGetTokenizedCustomerOrders)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
