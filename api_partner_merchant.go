package sdk

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// CreatePartnerMerchant
type RequestCreatePartnerMerchant struct {
	CompanyName string `json:"companyName"`
	Email       string `json:"email"`
	CountryCode string `json:"countryCode"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
}

type ResponseCreatePartnerMerchant struct {
	Code string                            `json:"code"`
	Item ResponseCreatePartnerMerchantItem `json:"item"`
	Err  *Error                            `json:"error"`
}

type ResponseCreatePartnerMerchantItem struct {
	Merchant `json:"merchant"`
	Store    `json:"store"`
	User     `json:"user"`
}

func (c Client) CreatePartnerMerchant(i RequestCreatePartnerMerchant) (*ResponseCreatePartnerMerchant, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPICreatePartnerMerchantURL.method
	requestURL := c.prepareAPIURL(pathAPICreatePartnerMerchantURL)

	response := new(ResponseCreatePartnerMerchant)
	if err := c.httpAPI(method, fmt.Sprintf("%s", requestURL), i, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// GetPartnerMerchants
type RequestGetPartnerMerchants struct{}

type ResponseGetPartnerMerchants struct {
	Code  string     `json:"code"`
	Items []Merchant `json:"items"`
	Meta  struct {
		Cursor string `json:"cursor"`
	} `json:"meta"`
	Err *Error `json:"error"`
}

func (c Client) GetPartnerMerchants(i RequestGetPartnerMerchants) (*ResponseGetPartnerMerchants, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetPartnerMerchantsURL.method
	requestURL := c.prepareAPIURL(pathAPIGetPartnerMerchantsURL)

	response := new(ResponseGetPartnerMerchants)
	if err := c.httpAPI(method, fmt.Sprintf("%s", requestURL), i, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// GetPartnerMerchantByID
type RequestGetPartnerMerchantByID struct {
	MerchantID string
}

type ResponseGetPartnerMerchantByID struct {
	Code string                             `json:"code"`
	Item ResponseGetPartnerMerchantByIDItem `json:"item"`
	Err  *Error                             `json:"error"`
}

type ResponseGetPartnerMerchantByIDItem struct {
	Merchant           `json:"merchant"`
	MerchantSettlement `json:"settlement"`
}

func (c Client) GetPartnerMerchantByID(i RequestGetPartnerMerchantByID) (*ResponseGetPartnerMerchantByID, error) {
	if c.err != nil {
		return nil, c.err
	}

	if i.MerchantID == "" {
		return nil, errors.New("Merchant ID is required")
	}

	method := pathAPIGetPartnerMerchantByIDURL.method
	requestURL := c.prepareAPIURL(pathAPIGetPartnerMerchantByIDURL)
	requestURL = strings.ReplaceAll(requestURL, "{merchant_id}", i.MerchantID)

	response := new(ResponseGetPartnerMerchantByID)
	if err := c.httpAPI(method, fmt.Sprintf("%s", requestURL), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// UpdatePartnerMerchant
type RequestUpdatePartnerMerchant struct {
	MerchantID string

	BrandName                 string              `json:"brandName"`
	WebsiteURL                string              `json:"websiteUrl"`
	RegistrationNumber        string              `json:"registrationNumber"`
	EstablishedDateTime       *time.Time          `json:"establishedAt"`
	AddressLine1              string              `json:"addressLine1"`
	AddressLine2              string              `json:"addressLine2"`
	PostCode                  string              `json:"postCode,omitempty" validate:"omitempty,numeric,min=3,max=10"`
	City                      string              `json:"city"`
	State                     string              `json:"state"`
	Country                   string              `json:"country"`
	CountryCode               string              `json:"countryCode,omitempty" validate:"omitempty,countrycode"`
	PhoneNumber               string              `json:"phoneNumber,omitempty" validate:"omitempty,numeric,min=7,max=12"`
	MerchantName              string              `json:"companyName"`
	CompanyType               CompanyType         `json:"companyType"`
	BusinessCategory          string              `json:"businessCategory"`
	AverageTicketSize         uint64              `json:"averageTicketSize"`
	AverageTurnoverPerMonth   uint64              `json:"averageTurnoverPerMonth"`
	BusinessScope             string              `json:"businessScope"`
	InvoiceAddress            *InvoiceAddress     `json:"invoiceAddress"`
	Document                  Document            `json:"document"`
	InspectList               []InspectList       `json:"inspectList,omitempty" validate:"omitempty,gte=1,dive"`
	BankAccountType           BankAccountType     `json:"bankAccountType,omitempty"`
	BankAccountHolderName     string              `json:"bankAccountHolderName"`
	BankAccountNo             string              `json:"bankAccountNo"`
	BankCode                  string              `json:"bankCode"`
	Latitude                  float64             `json:"latitude"`
	Longitude                 float64             `json:"longitude"`
	PaymentSubscription       PaymentSubscription `json:"paymentSubscription,omitempty" validate:"omitempty"`
	TerminalOfflineEWallet    bool                `json:"terminalOfflineEWallet"`
	TerminalOfflineCreditCard bool                `json:"terminalOfflineCreditCard"`
	OnlineCreditCard          bool                `json:"onlineCreditCard"`
	VopAPI                    bool                `json:"vopApi"`
}

type ResponseUpdatePartnerMerchant struct {
	Code string             `json:"code"`
	Err  *Error             `json:"error"`
	Item MerchantSettlement `json:"item"`
}

func (c Client) UpdatePartnerMerchant(i RequestUpdatePartnerMerchant) (*ResponseUpdatePartnerMerchant, error) {
	if c.err != nil {
		return nil, c.err
	}

	if i.MerchantID == "" {
		return nil, errors.New("Merchant ID is required")
	}

	method := pathAPIUpdatePartnerMerchantURL.method
	requestURL := c.prepareAPIURL(pathAPIUpdatePartnerMerchantURL)
	requestURL = strings.ReplaceAll(requestURL, "{merchant_id}", i.MerchantID)

	response := new(ResponseUpdatePartnerMerchant)
	if err := c.httpAPI(method, fmt.Sprintf("%s", requestURL), i, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// SubmitPartnerMerchant
type RequestSubmitPartnerMerchant struct {
	MerchantID string
}

type ResponseSubmitPartnerMerchant struct {
	Code string             `json:"code"`
	Err  *Error             `json:"error"`
	Item MerchantSettlement `json:"item"`
}

func (c Client) SubmitPartnerMerchant(i RequestSubmitPartnerMerchant) (*ResponseSubmitPartnerMerchant, error) {
	if c.err != nil {
		return nil, c.err
	}

	if i.MerchantID == "" {
		return nil, errors.New("Merchant ID is required")
	}

	method := pathAPISubmitPartnerMerchantURL.method
	requestURL := c.prepareAPIURL(pathAPISubmitPartnerMerchantURL)
	requestURL = strings.ReplaceAll(requestURL, "{merchant_id}", i.MerchantID)

	response := new(ResponseSubmitPartnerMerchant)
	if err := c.httpAPI(method, fmt.Sprintf("%s", requestURL), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
