package sdk

import (
	"fmt"
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
	Item struct {
		Key                string    `json:"key"`
		Label              string    `json:"label"`
		VoucherBatchKey    string    `json:"voucherBathKey"`
		Type               string    `json:"type"`
		Amount             int       `json:"amount"`
		DiscountRate       int       `json:"discountRate"`
		MinimumSpendAmount int       `json:"minimumSpendAmount"`
		Origin             string    `json:"origin"`
		ImageURL           string    `json:"imageUrl"`
		MemberProfile      string    `json:"memberProfile"`
		AssignedAt         time.Time `json:"assignedAt"`
		QrURL              string    `json:"qrUrl"`
		Code               string    `json:"code"`
		IsShipping         bool      `json:"isShipping"`
		Address            string    `json:"address"`
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
