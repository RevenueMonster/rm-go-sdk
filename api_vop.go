package sdk

import (
	"errors"
	"strings"
)

// RequestCreateVOPSubscription :
type RequestCreateVOPSubscription struct {
	TokenizedCustomerID string `json:"-"`
	NotifyURL           string `json:"notifyUrl"`
}

// RequestDeleteVOPSubscription :
type RequestDeleteVOPSubscription struct {
	TokenizedCustomerID string `json:"-"`
}

// ResponseCreateVOPSubscription :
type ResponseVOPSubscription struct {
	Item struct {
		ID              string `json:"id"`
		NotifyURL       string `json:"notifyUrl"`
		CreatedDateTime string `json:"createdAt"`
		UpdatedDateTime string `json:"updatedAt"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// CreateVOPSubscription :
func (c Client) CreateVOPSubscription(request RequestCreateVOPSubscription) (*ResponseVOPSubscription, error) {

	if c.err != nil {
		return nil, c.err
	}

	method := pathCreateVOPSubscription.method
	requestURL := c.prepareAPIURL(pathCreateVOPSubscription)
	requestURL = strings.ReplaceAll(requestURL, "{customer_id}", request.TokenizedCustomerID)

	response := new(ResponseVOPSubscription)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// DeleteVOPSubscription :
func (c Client) DeleteVOPSubscription(request RequestDeleteVOPSubscription) (*ResponseVOPSubscription, error) {

	if c.err != nil {
		return nil, c.err
	}

	method := pathDeleteVOPSubscription.method
	requestURL := c.prepareAPIURL(pathDeleteVOPSubscription)
	requestURL = strings.ReplaceAll(requestURL, "{customer_id}", request.TokenizedCustomerID)

	response := new(ResponseVOPSubscription)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
