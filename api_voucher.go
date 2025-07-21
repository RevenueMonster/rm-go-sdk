package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// VoucherVoidResponse :
type VoucherVoidResponse struct {
	Item Voucher `json:"item"`
	Code string  `json:"code"`
	Err  *Error  `json:"error"`
}

// VoucherReinstateResponse :
type VoucherReinstateResponse struct {
	Item Voucher `json:"item"`
	Code string  `json:"code"`
	Err  *Error  `json:"error"`
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
	Status   string `json:"status"`
	IsStatic bool   `json:"isStatic"`
	Cursor   string `json:"cursor"`
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

// VoucherReinstate :
func (c Client) VoucherReinstate(code, reason, pin string) (*VoucherReinstateResponse, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIVoucherReinstateURL.method
	requestURL := c.prepareAPIURL(pathAPIVoucherReinstateURL)

	if pin == "" {
		return nil, errors.New("pin is required for voucher reinstatement")
	}

	response := new(VoucherReinstateResponse)
	if err := c.httpAPI(method, fmt.Sprintf("%s/%s/reinstate", requestURL, code), struct {
		Reason string `json:"reason,omitempty"`
		Pin    string `json:"pin"`
	}{
		Reason: reason,
		Pin:    pin,
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

	code = strings.Replace(code, "#", "%23", -1)

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

	type Filter struct {
		Status   string `json:"status,omitempty"`
		IsStatic bool   `json:"isStatic,omitempty"`
	}

	filter := new(Filter)
	if request.Status != "" {
		filter.Status = request.Status
	}

	if request.IsStatic {
		filter.IsStatic = true
	}

	data, _ := json.Marshal(filter)
	parameters.Add("filter", string(data))
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
