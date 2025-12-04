package sdk

import (
	"errors"
	"strings"
)

// ResponseMerchantSettlement :
type ResponseMerchantSettlement struct {
	Item MerchantSettlement `json:"item"`
	Code string             `json:"code"`
	Err  *Error             `json:"error"`
}

// ResponseMerchantSubscriptions :
type ResponseMerchantSettlements struct {
	Items []MerchantSettlement `json:"items"`
	Code  string               `json:"code"`
	Meta  struct {
		Cursor string `json:"cursor"`
	} `json:"meta"`
	Err *Error `json:"error"`
}

// GetMerchantSettlementList :
func (c Client) GetMerchantSettlementList(cursor string) (*ResponseMerchantSettlements, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetSettlementAccountList.method
	requestURL := c.prepareAPIURL(pathAPIGetSettlementAccountList)
	if cursor != "" {
		requestURL += "?cursor=" + cursor
	}

	response := new(ResponseMerchantSettlements)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// GetMerchantSettlement :
func (c Client) GetMerchantSettlement() (*ResponseMerchantSettlement, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetSettlementAccount.method
	requestURL := c.prepareAPIURL(pathAPIGetSettlementAccount)

	response := new(ResponseMerchantSettlement)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// GetMerchantSettlementById :
func (c Client) GetMerchantSettlementById(id string) (*ResponseMerchantSettlement, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetSettlementAccountById.method
	requestURL := c.prepareAPIURL(pathAPIGetSettlementAccountById)
	requestURL = strings.ReplaceAll(requestURL, "{settlement_account_id}", id)

	response := new(ResponseMerchantSettlement)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
