package sdk

import (
	"errors"
	"fmt"
	"time"
)

// RequestCreatePaymentTransactionQRCode :
type RequestCreatePaymentTransactionQRCode struct {
	Amount       uint64   `json:"amount"`
	CurrencyType string   `json:"currencyType"`
	Method       []string `json:"method"`
	Expiry       struct {
		Type            string    `json:"type"`
		ExpiredDateTime time.Time `json:"expiredAt"`
	} `json:"expiry"`
	Order struct {
		Title          string `json:"title"`
		Detail         string `json:"detail"`
		AdditionalData string `json:"additionalData"`
	} `json:"order"`
	RedirectURL     string `json:"redirectUrl"`
	Type            string `json:"type"`
	StoreID         string `json:"storeId"`
	IsPreFillAmount bool   `json:"isPreFillAmount"`
}

// ResponsePaymentTransactionQRCodeItem :
type ResponsePaymentTransactionQRCodeItem struct {
	Item PaymentTransactionQRCode `json:"item"`
	Code string                   `json:"code"`
	Err  *Error                   `json:"error"`
}

// CreatePaymentTransactionQRCode :
func (c Client) CreatePaymentTransactionQRCode(request RequestCreatePaymentTransactionQRCode) (*ResponsePaymentTransactionQRCodeItem, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPICreatePaymentTransactionQRURL.method
	requestURL := c.prepareAPIURL(pathAPICreatePaymentTransactionQRURL)

	response := new(ResponsePaymentTransactionQRCodeItem)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// ResponseGetPaymentTransactionQRCode :
type ResponseGetPaymentTransactionQRCode struct {
	Item *struct {
		Store struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			AddressLine1 string `json:"addressLine1"`
			AddressLine2 string `json:"addressLine2"`
			PostCode     string `json:"postCode"`
			City         string `json:"city"`
			State        string `json:"state"`
			Country      string `json:"country"`
			CountryCode  string `json:"countryCode"`
			PhoneNumber  string `json:"phoneNumber"`
			GeoLocation  struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"geoLocation"`
			Status    string    `json:"status"`
			CreatedAt time.Time `json:"createdAt"`
			UpdatedAt time.Time `json:"updatedAt"`
		} `json:"store"`
		Type            string   `json:"type"`
		IsPreFillAmount bool     `json:"isPreFillAmount"`
		CurrencyType    string   `json:"currencyType"`
		Amount          int      `json:"amount"`
		Platform        string   `json:"platform"`
		Method          []string `json:"method"`
		Expiry          struct {
			Type      string    `json:"type"`
			Day       int       `json:"day"`
			ExpiredAt time.Time `json:"expiredAt"`
		} `json:"expiry"`
		Code        string `json:"code"`
		Status      string `json:"status"`
		QrCodeURL   string `json:"qrCodeUrl"`
		RedirectURL string `json:"redirectUrl"`
		Order       struct {
			Title  string `json:"title"`
			Detail string `json:"detail"`
		} `json:"order"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// GetPaymentTransactionQRByCode :
func (c Client) GetPaymentTransactionQRByCode(code string) (*ResponseGetPaymentTransactionQRCode, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetPaymentTransactionQRURL.method
	requestURL := c.prepareAPIURL(pathAPIGetPaymentTransactionQRURL)

	response := new(ResponseGetPaymentTransactionQRCode)
	if err := c.httpAPI(method, fmt.Sprintf("%s/%s", requestURL, code), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
