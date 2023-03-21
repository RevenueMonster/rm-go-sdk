package sdk

import (
	"errors"
	"fmt"
)

// ResponseMerchantItem :
type ResponseFpxBankList struct {
	Item map[string]ResponseFpxBank `json:"item"`
	Code string                     `json:"code"`
	Err  *Error                     `json:"error"`
}

type ResponseFpxBank struct {
	Code     string `json:"code"`
	IsOnline bool   `json:"isOnline"`
	Name     string `json:"name"`
}

// GetMerchantProfile :
func (c Client) GetFpxBankList() (*ResponseFpxBankList, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetFpxBankList.method
	requestURL := c.prepareAPIURL(pathAPIGetFpxBankList)

	response := new(ResponseFpxBankList)
	if err := c.httpAPI(method, fmt.Sprintf("%s", requestURL), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
