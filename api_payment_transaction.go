package sdk

import (
	"errors"
	"fmt"
	"net/url"
)

// RequestRefundPaymentTransaction :
type RequestRefundPaymentTransaction struct {
	TransactionID string `json:"transactionId"`
	Refund        struct {
		Type         string `json:"type"`
		CurrencyType string `json:"currencyType"`
		Amount       uint64 `json:"amount"`
	} `json:"refund"`
	Reason string `json:"reason"`
}

// ResponsePaymentTransactionItem :
type ResponsePaymentTransactionItem struct {
	Item PaymentTransaction `json:"item"`
	Code string             `json:"code"`
	Err  *Error             `json:"error"`
}

// ResponsePaymentTransactionItems :
type ResponsePaymentTransactionItems struct {
	Items []PaymentTransaction `json:"items"`
	Code  string               `json:"code"`
	Err   *Error               `json:"error"`
}

// GetPaymentTransactionByID :
func (c Client) GetPaymentTransactionByID(transactionID string) (*ResponsePaymentTransactionItem, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetPaymentTransactionByIDURL.method
	requestURL := c.prepareAPIURL(pathAPIGetPaymentTransactionByIDURL)

	response := new(ResponsePaymentTransactionItem)
	if err := c.httpAPI(method, fmt.Sprintf("%s/%s", requestURL, transactionID), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// RefundPaymentTransaction :
func (c Client) RefundPaymentTransaction(request RequestRefundPaymentTransaction) (*ResponsePaymentTransactionItem, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIRefundPaymentURL.method
	requestURL := c.prepareAPIURL(pathAPIRefundPaymentURL)

	response := new(ResponsePaymentTransactionItem)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// ReversedTransaction :
func (c Client) ReversedTransaction(orderID string) (*ResponsePaymentTransactionItem, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIReversedPaymentURL.method
	requestURL := c.prepareAPIURL(pathAPIReversedPaymentURL)

	response := new(ResponsePaymentTransactionItem)
	if err := c.httpAPI(method, requestURL, struct {
		OrderID string `json:"orderId"`
	}{
		OrderID: orderID,
	}, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// GetSuccessPaymentTransactionsByQRCode :
func (c Client) GetSuccessPaymentTransactionsByQRCode(code string) (*ResponsePaymentTransactionItems, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetPaymentTransactionQRURL.method
	requestURL := c.prepareAPIURL(pathAPIGetPaymentTransactionQRURL)

	endpoint := fmt.Sprintf("%s/%s/transactions", requestURL, code)
	rawURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	parameters := url.Values{}
	parameters.Add("filter", "{\"status\": \"SUCCESS\"}")
	rawURL.RawQuery = parameters.Encode()

	response := new(ResponsePaymentTransactionItems)
	if err := c.httpAPI(method, rawURL.String(), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// GetPaymentTransactionByOrderID :
func (c Client) GetPaymentTransactionByOrderID(orderID string) (*ResponsePaymentTransactionItem, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetPaymentTransactionByOrderIDURL.method
	requestURL := c.prepareAPIURL(pathAPIGetPaymentTransactionByOrderIDURL)

	response := new(ResponsePaymentTransactionItem)
	if err := c.httpAPI(method, fmt.Sprintf("%s/%s", requestURL, orderID), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
