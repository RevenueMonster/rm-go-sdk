package sdk

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// VoucherVoidResponse :
type VoucherVoidResponse struct {
	Item struct {
		Key                string      `json:"key"`
		Label              string      `json:"label"`
		RedemptionRuleKey  interface{} `json:"redemptionRuleKey"`
		VoucherBatchKey    string      `json:"voucherBatchKey"`
		Type               string      `json:"type"`
		Amount             int         `json:"amount"`
		DiscountRate       int         `json:"discountRate"`
		MinimumSpendAmount int         `json:"minimumSpendAmount"`
		Origin             string      `json:"origin"`
		ImageURL           string      `json:"imageUrl"`
		MemberProfile      interface{} `json:"memberProfile"`
		RedemptionRule     interface{} `json:"redemptionRule"`
		AssignedAt         time.Time   `json:"assignedAt"`
		Payload            interface{} `json:"payload"`
		QrURL              string      `json:"qrUrl"`
		Code               string      `json:"code"`
		IsShipping         bool        `json:"isShipping"`
		Address            interface{} `json:"address"`
		Expiry             struct {
			Type      string    `json:"type"`
			Day       int       `json:"day"`
			ExpiredAt time.Time `json:"expiredAt"`
		} `json:"expiry"`
		UsedAt         time.Time `json:"usedAt"`
		RedeemedAt     time.Time `json:"redeemedAt"`
		IsDeviceRedeem bool      `json:"isDeviceRedeem"`
		Status         string    `json:"status"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// GetVoucherByCodeResponse :
type GetVoucherByCodeResponse struct {
	Item Voucher `json:"item"`
	Code string  `json:"code"`
	Err  *Error  `json:"error"`
}

// RequestGetVoucherBatchByKey :
type RequestGetVoucherBatchByKey struct {
	Key    string `json:"key"`
	Cursor string `json:"cursor"`
}

// GetVoucherBatchByKeyResponse :
type GetVoucherBatchByKeyResponse struct {
	Items []Voucher `json:"items"`
	Code  string    `json:"code"`
	Meta  struct {
		Count  int    `json:"count"`
		Cursor string `json:"cursor"`
	} `json:"meta"`
	Err *Error `json:"error"`
}

// RequestGetVoucherBatches :
type RequestGetVoucherBatches struct {
	Status          string `json:"status"`
	BalanceQuantity int    `json:"balanceQuantity"`
	Cursor          string `json:"cursor"`
}

// GetVoucherBatchesResponse :
type GetVoucherBatchesResponse struct {
	Items []VoucherBatch `json:"items"`
	Code  string         `json:"code"`
	Meta  struct {
		Count  int    `json:"count"`
		Cursor string `json:"cursor"`
	} `json:"meta"`
	Err *Error `json:"error"`
}

// VoucherVoid :
func (c Client) VoucherVoid(code string, usedDateTime time.Time) (*VoucherVoidResponse, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIVoucherVoidURL.method
	requestURL := c.prepareAPIURL(pathAPIVoucherVoidURL)

	response := new(VoucherVoidResponse)
	if err := c.httpAPI(method, fmt.Sprintf("%s/%s/void", requestURL, code), struct {
		UsedDateTime time.Time `json:"usedAt"`
	}{
		UsedDateTime: usedDateTime,
	}, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetVoucherByCode :
func (c Client) GetVoucherByCode(code string) (*GetVoucherByCodeResponse, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetVoucherByCodeURL.method
	requestURL := c.prepareAPIURL(pathAPIGetVoucherByCodeURL)

	response := new(GetVoucherByCodeResponse)
	if err := c.httpAPI(method, fmt.Sprintf("%s/%s", requestURL, code), nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetVoucherBatches :
func (c Client) GetVoucherBatches(request RequestGetVoucherBatches) (*GetVoucherBatchesResponse, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetVoucherBatchesURL.method
	requestURL := c.prepareAPIURL(pathAPIGetVoucherBatchesURL)

	rawURL, err := url.Parse(requestURL)
	if err != nil {
		return nil, err
	}

	parameters := url.Values{}
	parameters.Add("cursor", request.Cursor)
	if request.Status != "" {
		parameters.Add("filter", "{\"status\": \""+request.Status+"\"}")
	}
	if request.BalanceQuantity > 0 {
		parameters.Add("filter", "{\"balanceQuantity\": {\"$gte\": \""+fmt.Sprintf("%d", request.BalanceQuantity)+"\"}}")
	}
	rawURL.RawQuery = parameters.Encode()

	response := new(GetVoucherBatchesResponse)
	if err := c.httpAPI(method, rawURL.String(), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// GetVoucherBatchByKey :
func (c Client) GetVoucherBatchByKey(request RequestGetVoucherBatchByKey) (*GetVoucherBatchByKeyResponse, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetVoucherBatchByKeyURL.method
	requestURL := c.prepareAPIURL(pathAPIGetVoucherBatchByKeyURL)

	rawURL, err := url.Parse(fmt.Sprintf("%s/%s/vouchers", requestURL, request.Key))
	if err != nil {
		return nil, err
	}

	parameters := url.Values{}
	parameters.Add("cursor", request.Cursor)

	rawURL.RawQuery = parameters.Encode()

	response := new(GetVoucherBatchByKeyResponse)
	if err := c.httpAPI(method, rawURL.String(), nil, response); err != nil {
		return nil, err
	}

	return response, nil
}
