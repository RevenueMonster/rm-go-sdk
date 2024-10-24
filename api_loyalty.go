package sdk

import (
	"errors"
	"time"
)

// requestRegisterLoyaltyMember :
type RequestRegisterLoyaltyMember struct {
	Name        string    `json:"name" validate:"required"`
	CountryCode string    `json:"countryCode" validate:"required,numeric,min=1,max=4"`
	PhoneNumber string    `json:"phoneNumber" validate:"required,numeric,min=7,max=12"`
	Email       string    `json:"email" validate:"omitempty,email"`
	NRIC        string    `json:"nric,omitempty"`
	Passport    string    `json:"passport,omitempty"`
	BirthDate   time.Time `json:"birthDate,omitempty"`
	Gender      string    `json:"gender" validate:"omitempty,eq=MALE|eq=FEMALE"`
	State       string    `json:"state,omitempty"`
	Address     struct {
		AddressLine1 string `json:"addressLine1,omitempty"`
		AddressLine2 string `json:"addressLine2,omitempty"`
		Postcode     string `json:"postcode,omitempty" validate:"omitempty,numeric"`
		City         string `json:"city,omitempty"`
		State        string `json:"state,omitempty"`
		Country      string `json:"country,omitempty"`
	} `json:"address,omitempty"`
	Point        uint64 `json:"loyaltyPoint"`
	Status       string `json:"status"`
	Error        string `json:"error"`
	ReferralCode string `json:"referralCode,omitempty"`
}

// responseRegisterLoyaltyMember :
type ResponseRegisterLoyaltyMember struct {
	Item RequestRegisterLoyaltyMember `json:"item"`
	Code string                       `json:"code"`
	Err  *Error                       `json:"error"`
}

func (c Client) RegisterLoyaltyMember(request RequestRegisterLoyaltyMember) (*ResponseRegisterLoyaltyMember, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathRegisterLoyaltyMember.method
	requestURL := c.prepareAPIURL(pathRegisterLoyaltyMember)

	response := new(ResponseRegisterLoyaltyMember)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return nil, errors.New(response.Err.Message)
	}

	return nil, nil
}
