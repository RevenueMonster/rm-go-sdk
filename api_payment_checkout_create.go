package sdk

import (
	"errors"
	"fmt"
)

// RequestExtraInfoExtraFee :
type RequestExtraInfoExtraFee struct {
	Type        string `json:"type"`
	ReferenceID string `json:"referenceId"`
	Amount      uint64 `json:"amount"`
	Source      string `json:"source"`
}

// RequestAmountBreakdown :
type RequestAmountBreakdown struct {
	Type              string `json:"type"`
	IsDiscountAllowed bool   `json:"isDiscountAllowed"`
	Label             string `json:"label"`
	Amount            uint64 `json:"amount"`
}

// RequestInHousePromo :
type RequestInHousePromo struct {
	Label    string `json:"label"`
	UniqueID string `json:"uniqueId"`
	Source   string `json:"source"`
	Amount   uint64 `json:"amount"`
}

// RequestCustomer :
type RequestCustomer struct {
	UserID      string `json:"userId"`
	Email       string `json:"email" validate:"omitempty,email"`
	CountryCode string `json:"countryCode"`
	PhoneNumber string `json:"phoneNumber"`
}

// RequestCreatePaymentCheckout :
type RequestCreatePaymentCheckout struct {
	Order struct {
		ID              string                   `json:"id"`
		Title           string                   `json:"title"`
		Detail          string                   `json:"detail"`
		AdditionalData  string                   `json:"additionalData"`
		CurrencyType    string                   `json:"currencyType"`
		Amount          uint64                   `json:"amount"`
		AmountBreakdown []RequestAmountBreakdown `json:"amountBreakdown"`
	} `json:"order"`
	Method        []string `json:"method"`
	ExcludeMethod []string `json:"excludeMethod"`
	Type          string   `json:"type"`
	StoreID       string   `json:"storeId"`
	RedirectURL   string   `json:"redirectUrl"`
	NotifyURL     string   `json:"notifyUrl"`
	LayoutVersion string   `json:"layoutVersion"`
	ExtraInfo     struct {
		ExtraFee []RequestExtraInfoExtraFee `json:"extraFee"`
	} `json:"extraInfo"`
	Customer RequestCustomer `json:"customer"`
	Voucher  struct {
		Code string `json:"code"`
	} `json:"voucher"`
	Source           string                `json:"source"`
	InHousePromo     []RequestInHousePromo `json:"inHousePromo"`
	ExpiresInSeconds int64                 `json:"expiresInSeconds"`
}

// ResponseCreatePaymentCheckout :
type ResponseCreatePaymentCheckout struct {
	Item struct {
		CheckoutID string `json:"checkoutId"`
		URL        string `json:"url"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// CreatePaymentCheckout :
func (c Client) CreatePaymentCheckout(request RequestCreatePaymentCheckout) (*ResponseCreatePaymentCheckout, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPICreatePaymentCheckoutURL.method
	requestURL := c.prepareAPIURL(pathAPICreatePaymentCheckoutURL)

	response := new(ResponseCreatePaymentCheckout)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// ResponseGetQRCodeByCheckoutID :
type ResponseGetQRCodeByCheckoutID struct {
	Item struct {
		QRCodeImageBase64 string `json:"qrCodeImageBase64"`
		URL               string `json:"url"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// GetQRCodeByCheckoutID :
func (c Client) GetQRCodeByCheckoutID(checkoutID, walletMethod string) (*ResponseGetQRCodeByCheckoutID, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetQRCodeByCheckoutIDURL.method
	requestURL := c.prepareAPIURL(pathAPIGetQRCodeByCheckoutIDURL)

	response := new(ResponseGetQRCodeByCheckoutID)
	if err := c.httpAPI(method, fmt.Sprintf("%s?checkoutId=%s&method=%s", requestURL, checkoutID, walletMethod), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
