package sdk

import (
	"errors"
	"fmt"
)

// ResponseGetPaymentSubscriptionStatus :
type ResponseGetPaymentSubscriptionStatus struct {
	Item struct {
		Online  map[string]string `json:"online"`
		Offline map[string]string `json:"offline"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// GetPaymentSubscriptionStatus :
func (c Client) GetPaymentSubscriptionStatus() (*ResponseGetPaymentSubscriptionStatus, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetPaymentSubscriptionStatus.method
	requestURL := c.prepareAPIURL(pathAPIGetPaymentSubscriptionStatus)
	response := new(ResponseGetPaymentSubscriptionStatus)
	if err := c.httpAPI(method, fmt.Sprintf("%s", requestURL), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
