package sdk

import "time"

// PaymentTransactionQRCode :
type PaymentTransactionQRCode struct {
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
	Amount          uint64   `json:"amount"`
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
}

// PaymentTransaction :
type PaymentTransaction struct {
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
	TransactionID string `json:"transactionId"`
	ReferenceID   string `json:"referenceId"`
	Order         struct {
		ID             string `json:"id"`
		Title          string `json:"title"`
		Detail         string `json:"detail"`
		AdditionalData string `json:"additionalData"`
		Amount         int    `json:"amount"`
	} `json:"order"`
	CurrencyType  string `json:"currencyType"`
	BalanceAmount int    `json:"balanceAmount"`
	Platform      string `json:"platform"`
	Method        string `json:"method"`
	Error         struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	TransactionAt time.Time `json:"transactionAt"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	Region        string    `json:"region"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// Merchant :
type Merchant struct {
	ID                 string    `json:"id"`
	CompanyName        string    `json:"companyName"`
	CompanyType        string    `json:"companyType"`
	CompanyLogoURL     string    `json:"companyLogoUrl"`
	RegistrationNumber string    `json:"registrationNumber"`
	BusinessCategory   string    `json:"businessCategory"`
	EstablishedAt      time.Time `json:"establishedAt"`
	CountryCode        string    `json:"countryCode"`
	PhoneNumber        string    `json:"phoneNumber"`
	AddressLine1       string    `json:"addressLine1"`
	AddressLine2       string    `json:"addressLine2"`
	Postcode           string    `json:"postcode"`
	City               string    `json:"city"`
	State              string    `json:"state"`
	Country            string    `json:"country"`
	InvoiceAddress     struct {
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		Postcode     string `json:"postcode"`
		City         string `json:"city"`
		State        string `json:"state"`
		Country      string `json:"country"`
	} `json:"invoiceAddress"`
	IsActive         bool      `json:"isActive"`
	Status           string    `json:"status"`
	IsMasterMerchant bool      `json:"isMasterMerchant"`
	MasterMerchantID string    `json:"masterMerchantId"`
	IsPartner        bool      `json:"isPartner"`
	PartnerID        string    `json:"partnerId"`
	GstNo            string    `json:"gstNo"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

// Store :
type Store struct {
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
}
