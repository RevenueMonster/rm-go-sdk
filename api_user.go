package sdk

import (
	"errors"
	"fmt"
)

// ResponseUserItem :
type ResponseUserItem struct {
	Item User   `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// GetUserProfile :
func (c Client) GetUserProfile() (*ResponseUserItem, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetUserURL.method
	requestURL := c.prepareAPIURL(pathAPIGetUserURL)

	response := new(ResponseUserItem)
	if err := c.httpAPI(method, fmt.Sprintf("%s", requestURL), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
