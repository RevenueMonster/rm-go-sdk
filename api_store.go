package sdk

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// RequestGetStores :
type RequestGetStores struct {
	Cursor        string    `json:"cursor"`
	LastUpdatedBy time.Time `json:"lastUpdatedBy"`
}

// ResponseStoreItems :
type ResponseStoreItems struct {
	Items []Store `json:"items"`
	Code  string  `json:"code"`
	Meta  struct {
		Count  int    `json:"count"`
		Cursor string `json:"cursor"`
	} `json:"meta"`
	Err *Error `json:"error"`
}

// GetStores :
func (c Client) GetStores(request RequestGetStores) (*ResponseStoreItems, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetStoresURL.method
	requestURL := c.prepareAPIURL(pathAPIGetStoresURL)

	rawURL, err := url.Parse(requestURL)
	if err != nil {
		return nil, err
	}

	parameters := url.Values{}
	parameters.Add("cursor", request.Cursor)
	if !request.LastUpdatedBy.IsZero() {
		parameters.Add("filter", "{\"updatedAt\": {\"$gte\": \""+request.LastUpdatedBy.Format(time.RFC3339)+"\"}}")
	}
	rawURL.RawQuery = parameters.Encode()

	response := new(ResponseStoreItems)
	if err := c.httpAPI(method, rawURL.String(), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// ResponseStoreItem :
type ResponseStoreItem struct {
	Item Store  `json:"item"`
	Code string `json:"code"`
	Meta struct {
		Count  int    `json:"count"`
		Cursor string `json:"cursor"`
	} `json:"meta"`
	Err *Error `json:"error"`
}

// GetStoreByID :
func (c Client) GetStoreByID(storeID string) (*ResponseStoreItem, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetStoreByIDURL.method
	requestURL := c.prepareAPIURL(pathAPIGetStoreByIDURL)

	response := new(ResponseStoreItem)
	if err := c.httpAPI(method, fmt.Sprintf("%s/%s", requestURL, storeID), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
