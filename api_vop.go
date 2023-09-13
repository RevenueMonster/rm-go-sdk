package sdk

import (
	"errors"
)

type RequestVOPEnrollUser struct {
	Identifier  string `json:"identifier"`
	UserID      string `json:"userId"`
	CardPan     string `json:"cardPan"`
	Name        string `json:"name"`
	RedirectURL string `json:"redirectUrl"`
}

type RequestVOPUnenrollUser struct {
	UserID string `json:"userId"`
}

type RequestVOPWebhook struct {
	UserID               string `json:"userId"`
	CardID               string `json:"cardId"`
	TransactionID        string `json:"transactionId"`
	TransactionDateTime  string `json:"transactionAt"`
	Amount               uint64 `json:"amount"`
	Currency             string `json:"currency"`
	MerchantCategoryCode string `json:"merchantCategoryCode"`
}

type ResponseVOPUserItem struct {
	Item VOPEnrollUser `json:"item"`
	Code string        `json:"code"`
}

type ResponseVOPItem struct {
	Code string `json:"code"`
}

type VOPEnrollUser struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	PanLastFour     string `json:"panLastFour"`
	Type            string `json:"type"`
	CreatedDateTime string `json:"createdAt"`
	UpdatedDateTime string `json:"updatedAt"`
}

// VOPEnrollUser :
func (c Client) VOPEnrollUser(request RequestVOPEnrollUser) (*ResponseMerchantItem, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIVOPEnrollUserURL.method
	requestURL := c.prepareAPIURL(pathAPIVOPEnrollUserURL)

	response := new(ResponseMerchantItem)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// VOPUnenrollUser :
func (c Client) VOPUnenrollUser(request RequestVOPUnenrollUser) (*ResponseMerchantItem, error) {

	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIVOPUnenrollUserURL.method
	requestURL := c.prepareAPIURL(pathAPIVOPUnenrollUserURL)

	response := new(ResponseMerchantItem)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// VOPWebhook :
func (c Client) VOPWebhook(request RequestVOPUnenrollUser) (*ResponseMerchantItem, error) {

	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIVOPWebhookURL.method
	requestURL := c.prepareAPIURL(pathAPIVOPWebhookURL)

	response := new(ResponseMerchantItem)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
