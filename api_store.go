package sdk

import (
	"errors"
	"fmt"
)

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
func (c Client) GetStores(cursor string) (*ResponseStoreItems, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetStoresURL.method
	requestURL := c.prepareAPIURL(pathAPIGetStoresURL)

	response := new(ResponseStoreItems)
	if err := c.httpAPI(method, fmt.Sprintf("%s?cursor=%s", requestURL, cursor), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
