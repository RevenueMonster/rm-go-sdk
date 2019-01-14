package sdk

import (
	"errors"
	"fmt"
)

// ResponseMerchantItem :
type ResponseMerchantItem struct {
	Item Merchant `json:"item"`
	Code string   `json:"code"`
	Err  *Error   `json:"error"`
}

// GetMerchantProfile :
func (c Client) GetMerchantProfile() (*ResponseMerchantItem, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetMerchantURL.method
	requestURL := c.prepareAPIURL(pathAPIGetMerchantURL)

	response := new(ResponseMerchantItem)
	if err := c.httpAPI(method, fmt.Sprintf("%s", requestURL), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
