package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type InvoiceStatus string

const (
	InvoiceStatusPaid    = "PAID"
	InvoiceStatusPending = "PENDING"
	InvoiceStatusExpired = "EXPIRED"
	InvoiceStatusDeleted = "DELETED"
)

type Invoice struct {
	ID                 string `json:"id"`
	ReferenceID        string `json:"referenceId"`
	RecurringInvoiceID string `json:"recurringInvoiceId"`
	InvoiceGroupID     string `json:"invoiceGroupId"`

	MerchantID      string `json:"merchantId"`
	MerchantName    string `json:"merchantName"`
	MerchantLogoURL string `json:"merchantLogoUrl"`
	MerchantEmail   string `json:"merchantEmail"`

	StoreID string `json:"storeId"`

	PayerEmail       string `json:"payerEmail"`
	PayerName        string `json:"payerName"`
	PayerPhoneNumber string `json:"payerPhoneNumber"`
	PayerCountryCode string `json:"payerCountryCode"`

	Currency       string `json:"currency"`
	Amount         uint64 `json:"amount"`
	PaidAmount     uint64 `json:"paidAmount"`
	FeesPaidAmount uint64 `json:"feesPaidAmount"`
	NetPaidAmount  uint64 `json:"netPaidAmount"`

	InvoiceStatus InvoiceStatus `json:"invoiceStatus"`
	TransactionID string        `json:"transactionId"`

	InvoiceURL  string `json:"invoiceUrl"`
	RedirectURL string `json:"redirectUrl"`
	TncURL      string `json:"tncUrl"`

	Remarks         string `json:"remarks"`
	ShouldSendEmail bool   `json:"shouldSendEmail"`
	EmailSubject    string `json:"emailSubject"`
	Description     string `json:"description"`

	ExpiryAt        time.Time  `json:"expiryAt"`
	PaidAt          *time.Time `json:"paidAt"`
	CreatedDateTime time.Time  `json:"createdDateTime"`
	UpdatedDateTime time.Time  `json:"updatedDateTime"`
}

func (i *Invoice) UnmarshalJSON(data []byte) error {
	type Alias Invoice
	aux := &struct {
		PaidAt string `json:"paidAt"`
		*Alias
	}{Alias: (*Alias)(i)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	if aux.PaidAt != "" {
		t, err := time.Parse(time.RFC3339, aux.PaidAt)
		if err != nil {
			return err
		}
		i.PaidAt = &t
	}
	return nil
}

// GetInvoices

type RequestGetInvoices struct {
	StoreID       string
	Status        []InvoiceStatus
	Limit         int
	Cursor        string
	StartDateTime time.Time
	EndDateTime   time.Time
}

type ResponseGetInvoices struct {
	Cursor string    `json:"cursor"`
	Items  []Invoice `json:"items"`
	Err    *Error    `json:"error,omitempty"`
}

func (c Client) GetInvoices(i RequestGetInvoices) (*ResponseGetInvoices, error) {
	if c.err != nil {
		return nil, c.err
	}

	requestURL := c.prepareAPIURL(pathAPIGetInvoicesURL)
	rawURL, err := url.Parse(requestURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if i.StoreID != "" {
		params.Set("storeId", i.StoreID)
	}
	for _, s := range i.Status {
		params.Add("status", string(s))
	}
	if i.Limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", i.Limit))
	}
	if i.Cursor != "" {
		params.Set("cursor", i.Cursor)
	}
	if !i.StartDateTime.IsZero() {
		params.Set("start", i.StartDateTime.Format(time.RFC3339))
	}
	if !i.EndDateTime.IsZero() {
		params.Set("end", i.EndDateTime.Format(time.RFC3339))
	}
	rawURL.RawQuery = params.Encode()

	response := new(ResponseGetInvoices)
	if err := c.httpAPI(pathAPIGetInvoicesURL.method, rawURL.String(), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// CreateInvoice

type RequestCreateInvoice struct {
	ReferenceID      string    `json:"referenceId" validate:"required"`
	StoreID          string    `json:"storeId" validate:"required"`
	MerchantEmail    string    `json:"merchantEmail"`
	PayerEmail       string    `json:"payerEmail" validate:"required"`
	PayerName        string    `json:"payerName"`
	PayerPhoneNumber string    `json:"payerPhoneNumber"`
	PayerCountryCode string    `json:"payerCountryCode"`
	EmailSubject     string    `json:"emailSubject" validate:"required"`
	ShouldSendEmail  bool      `json:"shouldSendEmail"` // Optional, default: Yes
	ExpiryDateTime   time.Time `json:"expiryDateTime" validate:"required"`
	Amount           uint64    `json:"amount" validate:"required"`
	Currency         string    `json:"currency" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	RedirectURL      string    `json:"redirectUrl"`
}

type ResponseCreateInvoice struct {
	Code string  `json:"code"`
	Item Invoice `json:"item"`
	Err  *Error  `json:"error,omitempty"`
}

func (c Client) CreateInvoice(i RequestCreateInvoice) (*ResponseCreateInvoice, error) {
	if c.err != nil {
		return nil, c.err
	}

	err := validate.Struct(i)
	if err != nil {
		return nil, err
	}

	method := pathAPICreateInvoiceURL.method
	requestURL := c.prepareAPIURL(pathAPICreateInvoiceURL)

	response := new(ResponseCreateInvoice)
	if err := c.httpAPI(method, requestURL, i, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// GetInvoiceByID

type RequestGetInvoiceByID struct {
	ID string `json:"id" validate:"required"`
}

type ResponseGetInvoiceByID struct {
	Code string  `json:"code"`
	Item Invoice `json:"item"`
	Err  *Error  `json:"error,omitempty"`
}

func (c Client) GetInvoiceByID(i RequestGetInvoiceByID) (*ResponseGetInvoiceByID, error) {
	if c.err != nil {
		return nil, c.err
	}

	if err := validate.Struct(i); err != nil {
		return nil, err
	}

	method := pathAPIGetInvoiceByIDURL.method
	requestURL := c.prepareAPIURL(pathAPIGetInvoiceByIDURL)
	requestURL = strings.ReplaceAll(requestURL, "{invoice_id}", i.ID)

	response := new(ResponseGetInvoiceByID)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// CancelInvoiceByID

type RequestCancelInvoiceByID struct {
	ID string `json:"id" validate:"required"`
}

type ResponseCancelInvoiceByID struct {
	Code string  `json:"code"`
	Item Invoice `json:"item"`
	Err  *Error  `json:"error,omitempty"`
}

func (c Client) CancelInvoiceByID(i RequestCancelInvoiceByID) (*ResponseCancelInvoiceByID, error) {
	if c.err != nil {
		return nil, c.err
	}

	if err := validate.Struct(i); err != nil {
		return nil, err
	}

	method := pathAPICancelInvoiceByIDURL.method
	requestURL := c.prepareAPIURL(pathAPICancelInvoiceByIDURL)
	requestURL = strings.ReplaceAll(requestURL, "{invoice_id}", i.ID)

	response := new(ResponseCancelInvoiceByID)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// EditInvoice: at the moment multiple fields are required for edit invoice, RM will enhance this in the future to allow partial update

type RequestEditInvoice struct {
	ID               string    `json:"id" validate:"required"`
	ReferenceID      string    `json:"referenceId" validate:"required"`
	StoreID          string    `json:"storeId" validate:"required"`
	ShouldSendEmail  bool      `json:"shouldSendEmail"`
	EmailSubject     string    `json:"emailSubject" validate:"required"`
	PayerEmail       string    `json:"payerEmail" validate:"required"`
	PayerName        string    `json:"payerName"`
	PayerPhoneNumber string    `json:"payerPhoneNumber"`
	PayerCountryCode string    `json:"payerCountryCode"`
	Amount           uint64    `json:"amount" validate:"required"`
	Currency         string    `json:"currency" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	RedirectURL      string    `json:"redirectUrl"`
	Remarks          string    `json:"remarks"`
	ExpiryDateTime   time.Time `json:"expiryDateTime" validate:"required"`
}

type ResponseEditInvoice struct {
	Code string  `json:"code"`
	Item Invoice `json:"item"`
	Err  *Error  `json:"error,omitempty"`
}

func (c Client) EditInvoice(i RequestEditInvoice) (*ResponseEditInvoice, error) {
	if c.err != nil {
		return nil, c.err
	}

	err := validate.Struct(i)
	if err != nil {
		return nil, err
	}

	method := pathAPIEditInvoiceURL.method
	requestURL := c.prepareAPIURL(pathAPIEditInvoiceURL)
	requestURL = strings.ReplaceAll(requestURL, "{invoice_id}", i.ID)

	response := new(ResponseEditInvoice)
	if err := c.httpAPI(method, requestURL, i, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
