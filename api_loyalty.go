package sdk

import (
	"errors"
	"strings"
	"time"
)

// RequestCheckMemberExist :
type RequestCheckMemberExist struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
}

// ResponseCheckMemberExist :
type ResponseCheckMemberExist struct {
	Item struct {
		Exist bool `json:"exist"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

func (c Client) CheckMemberExist(request RequestCheckMemberExist) (*ResponseCheckMemberExist, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathCheckMemberExist.method
	requestURL := c.prepareAPIURL(pathCheckMemberExist)
	response := new(ResponseCheckMemberExist)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}
	return response, nil
}

// requestRegisterLoyaltyMember :

type RequestRegisterLoyaltyMember struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	Email           string             `json:"email"`
	NRIC            string             `json:"nric"`
	Passport        string             `json:"passport"`
	Address         Address            `json:"address,omitempty"`
	Gender          string             `json:"gender"`
	State           string             `json:"state"`
	ReferralCode    string             `json:"referralCode"`
	BirthDate       time.Time          `json:"birthDate"`
	LoyaltyPoint    uint64             `json:"loyaltyPoint"`
	Credit          uint64             `json:"credit"`
	CountryCode     string             `json:"countryCode"`
	PhoneNumber     string             `json:"phoneNumber"`
	ProfileImageURL string             `json:"profileImageURL"`
	HasPinCode      bool               `json:"hasPinCode"`
	MemberTier      *LoyaltyMemberTier `json:"memberTier"`
	Status          string             `json:"status"`
	CreatedDateTime time.Time          `json:"createdDateTime"`
}

type Address struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	Postcode     string `json:"postcode"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}

type LoyaltyMemberTier struct {
	Label        string                     `json:"label"`
	MinimumPoint uint64                     `json:"minimumPoint"`
	Description  []string                   `json:"description"`
	Discount     *LoyaltyMemberTierDiscount `json:"discount"`
	CreatedAt    time.Time                  `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time                  `bson:"updatedAt" json:"updatedAt"`
}

const (
	LoyaltyMemberTierDiscountTypeCash     = "CASH"
	LoyaltyMemberTierDiscountTypeDiscount = "DISCOUNT"
)

type LoyaltyMemberTierDiscount struct {
	Type               string
	Amount             uint64
	MinimumSpendAmount uint64
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
		return response, errors.New(response.Err.Message)
	}
	return response, nil
}

type RequestGetLoyaltyMemberByID struct {
	MemberID string `json:"memberId"`
}

type ResponseGetLoyaltyMemberByID struct {
	Item RequestRegisterLoyaltyMember `json:"item"`
	Code string                       `json:"code"`
	Err  *Error                       `json:"error"`
}

type MerchantMember struct {
	Key             string    `json:"-"`
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	NRIC            string    `json:"nric"`
	Passport        string    `json:"passport"`
	Address         Address   `json:"address"`
	Gender          string    `json:"gender"`
	State           string    `json:"state"`
	ReferralCode    string    `json:"referralCode"`
	BirthDate       time.Time `json:"birthDate"`
	LoyaltyPoint    uint64    `json:"loyaltyPoint"`
	Credit          uint64    `json:"credit"`
	CountryCode     string    `json:"countryCode"`
	PhoneNumber     string    `json:"phoneNumber"`
	ProfileImageURL string    `json:"profileImageURL"`
	HasPinCode      bool      `json:"hasPinCode"`
	MemberTier      *string   `json:"memberTier"`
	Status          string    `json:"status"`
	CreatedDateTime time.Time `json:"createdDateTime"`
}

func (c Client) GetLoyaltyMemberByID(request RequestGetLoyaltyMemberByID) (*ResponseGetLoyaltyMemberByID, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathGetLoyaltyMemberByID.method
	requestURL := c.prepareAPIURL(pathGetLoyaltyMemberByID)
	requestURL = strings.ReplaceAll(requestURL, "{member_id}", request.MemberID)
	response := new(ResponseGetLoyaltyMemberByID)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return nil, errors.New(response.Err.Message)
	}

	return response, nil
}

// RequestGetLoyaltyMember :
type RequestGetLoyaltyMember struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
}

// ResponseGetLoyaltyMember :
type ResponseGetLoyaltyMember struct {
	Item MerchantMember `json:"item"`
	Code string         `json:"code"`
	Err  *Error         `json:"error"`
}

func (c Client) GetLoyaltyMember(request RequestGetLoyaltyMember) (*ResponseGetLoyaltyMember, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathGetLoyaltyMember.method
	requestURL := c.prepareAPIURL(pathGetLoyaltyMember)
	requestURL = strings.ReplaceAll(requestURL, "{country_code}", request.CountryCode)
	requestURL = strings.ReplaceAll(requestURL, "{phone_number}", request.PhoneNumber)
	response := new(ResponseGetLoyaltyMember)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return nil, errors.New(response.Err.Message)
	}

	return response, nil
}

type RequestLoyaltyCreditMemberTopUpOnline struct {
	MemberID    string `json:"memberId" validate:"required"`
	TopUpAmount int    `json:"topUpAmount" validate:"required"`
	RedirectURL string `json:"redirectURL" validate:"required"`
}

type ResponseLoyaltyCreditMemberTopUpOnline struct {
	Code string      `json:"code"`
	Err  *Error      `json:"error"`
	Item *PaymentUrl `json:"item"`
}

type PaymentUrl struct {
	PaymentUrl string `json:"paymentUrl"`
}

func (c Client) LoyaltyCreditMemberTopUpOnline(request RequestLoyaltyCreditMemberTopUpOnline) (*ResponseLoyaltyCreditMemberTopUpOnline, error) {

	if c.err != nil {
		return nil, c.err
	}

	method := pathLoyaltyCreditMemberTopUpOnline.method
	requestURL := c.prepareAPIURL(pathLoyaltyCreditMemberTopUpOnline)
	requestURL = strings.ReplaceAll(requestURL, "{member_id}", request.MemberID)

	response := new(ResponseLoyaltyCreditMemberTopUpOnline)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return nil, errors.New(response.Err.Message)
	}

	return response, nil
}

type RequestSpendBalance struct {
	AuthCode string `json:"authCode" validate:"required"`
	StoreID  int64  `json:"storeId,string" validate:"required"`
	Order    Order  `json:"order" validate:"required"`
}

type Order struct {
	ID             string `json:"id" validate:"required"`
	Title          string `json:"title" validate:"required"`
	Detail         string `json:"detail"`
	AdditionalData string `json:"additionalData,omitempty"`
	Amount         uint64 `json:"amount"`
}

type MemberProfile struct {
	ID                  string             `json:"id" goloquent:"-"`
	Key                 *string            `json:"key" goloquent:"__key__"`
	MerchantKey         *string            `json:"-"`
	Name                string             `json:"name"`
	Email               string             `json:"email"`
	NRIC                string             `json:"nric"`
	Passport            string             `json:"passport"`
	ReferralCode        string             `json:"referralCode"`
	BirthDate           time.Time          `json:"-"`
	BirthDateAt         string             `json:"birthDate" goloquent:"-"`
	Gender              string             `json:"gender"`
	State               string             `json:"state"`
	Address             Address            `json:"address" goloquent:",flatten"`
	ShippingAddress     []byte             `json:"-"`
	MemberTier          *LoyaltyMemberTier `json:"memberTier" goloquent:"-"`
	TotalLoyaltyPoint   uint64             `json:"totalLoyaltyPoint"`
	HasPinCode          bool               `json:"hasPinCode"`
	PinCodeSalt         []byte             `json:"-"`
	PinCodeHash         []byte             `json:"-"`
	LoyaltyPointBalance uint64             `json:"loyaltyPointBalance" goloquent:"-"`
	SpendingPoint       uint64             `json:"spendingPoint"`
	CreditBalance       uint64             `json:"creditBalance"`
	Status              string             `json:"status"`
	Payload             []byte             `json:"-"`
	CreatedAt           time.Time          `json:"createdAt"`
	UpdatedAt           time.Time          `json:"updatedAt"`
}

type LoyaltyTransaction struct {
	ID            string         `json:"id" goloquent:"-"`
	Key           *string        `json:"-" goloquent:"__key__"`
	MerchantKey   *string        `json:"-"`
	StoreKey      *string        `json:"-"`
	MemberKey     *string        `json:"-"`
	Store         *Store         `json:"store,omitempty" goloquent:"-"`
	MemberProfile *MemberProfile `json:"memberProfile,omitempty" goloquent:"-"`
	TransactionID string         `json:"transactionId"`
	Method        string         `json:"method" goloquent:"-"`
	Type          string         `json:"type"`
	Order         struct {
		ID             string `json:"id" qson:"Order.ID"`
		Title          string `json:"title"`
		Detail         string `json:"detail" goloquent:",longtext"`
		AdditionalData string `json:"additionalData,omitempty"`
		Amount         uint64 `json:"amount"`
	} `json:"order" goloquent:",flatten"`
	TerminalID          string    `json:"terminalId"`
	CurrencyType        string    `json:"currencyType"`
	TransactionDateTime time.Time `json:"transactionAt"`
	Status              string    `json:"status"`
	Payload             []byte    `json:"-"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type ResponseSpendBalance struct {
	Item *LoyaltyTransaction `json:"item"`
	Code string              `json:"code"`
	Err  *Error              `json:"error"`
}

func (c Client) SpendBalance(request RequestSpendBalance) (*ResponseSpendBalance, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathSpendBalance.method
	requestURL := c.prepareAPIURL(pathSpendBalance)
	response := new(ResponseSpendBalance)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return nil, errors.New(response.Err.Message)
	}

	return response, nil
}

type RequestRefundBalance struct {
	BalanceID    string `json:"balanceId" validate:"required"`
	Amount       uint64 `json:"amount" validate:"requiredif=IsFullRefund:false"`
	IsFullRefund bool   `json:"isFullRefund"`
}

type ResponseRefundBalance struct {
	Item *LoyaltyTransaction `json:"item"`
	Code string              `json:"code"`
	Err  *Error              `json:"error"`
}

func (c Client) RefundBalance(request RequestRefundBalance) (*ResponseRefundBalance, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathRefundBalance.method
	requestURL := c.prepareAPIURL(pathRefundBalance)
	response := new(ResponseRefundBalance)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return nil, errors.New(response.Err.Message)
	}

	return response, nil
}
