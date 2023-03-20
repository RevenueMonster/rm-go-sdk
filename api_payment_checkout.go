package sdk

import (
	"errors"
	"fmt"
	"time"
)

// RequestCreatePaymentCheckoutByMethod :
type RequestCreatePaymentCheckoutByMethod struct {
	Method     string                            `json:"method"`
	CheckoutID string                            `json:"checkoutId"`
	Type       string                            `json:"type"`
	Card       *RequestCreatePaymentCheckoutCard `json:"card"`
}

type RequestCreatePaymentCheckoutCard struct {
	IsToken     bool
	IsSave      bool
	No          string
	CVC         string
	Name        string
	Month       uint32
	Year        uint32
	CountryCode string
}

// ResponseCreatePaymentCheckoutByMethod :
type ResponseCreatePaymentCheckoutByMethod struct {
	Item struct {
		URL    string `json:"url"`
		Data   string `json:"data"`
		QRCode struct {
			Data        string `json:"data"`
			Base64Image string `json:"base64Image"`
		} `json:"qrcode"`
		Type string `json:"type"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// CreatePaymentCheckoutByMethod :
func (c Client) CreatePaymentCheckoutByMethod(request RequestCreatePaymentCheckoutByMethod) (*ResponseCreatePaymentCheckoutByMethod, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPICreatePaymentCheckoutByMethodURL.method
	requestURL := c.prepareAPIURL(pathAPICreatePaymentCheckoutByMethodURL)

	response := new(ResponseCreatePaymentCheckoutByMethod)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// ResponseGetOnlineTransactionByCheckoutID :
type ResponseGetOnlineTransactionByCheckoutID struct {
	Item struct {
		ID    string `json:"id"`
		Order struct {
			ID             string `json:"id"`
			Title          string `json:"title"`
			Detail         string `json:"detail"`
			AdditionalData string `json:"additionalData"`
			CurrencyType   string `json:"currencyType"`
			Amount         uint64 `json:"amount"`
		} `json:"order"`
		Type            string    `json:"type"`
		TransactionID   string    `json:"transactionId"`
		Platform        string    `json:"platform"`
		Method          []string  `json:"method"`
		RedirectURL     string    `json:"redirectUrl"`
		NotifyURL       string    `json:"notifyUrl"`
		StartDateTime   time.Time `json:"startAt"`
		EndDateTime     time.Time `json:"endAt"`
		Status          string    `json:"status"`
		CreatedDateTime time.Time `json:"createdAt"`
		UpdatedDateTime time.Time `json:"updatedAt"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// GetOnlineTransactionByCheckoutID :
func (c Client) GetOnlineTransactionByCheckoutID(checkoutID string) (*ResponseGetOnlineTransactionByCheckoutID, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetOnlineTransactionByCheckoutIDURL.method
	requestURL := c.prepareAPIURL(pathAPIGetOnlineTransactionByCheckoutIDURL)

	response := new(ResponseGetOnlineTransactionByCheckoutID)
	if err := c.httpAPI(method, fmt.Sprintf("%s?checkoutId=%s", requestURL, checkoutID), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
