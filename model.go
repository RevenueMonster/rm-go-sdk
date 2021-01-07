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
	Payee         struct {
		UserID    string `json:"userId"`
		SubUserID string `json:"subUserId"`
	} `json:"payee"`
	Error struct {
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
	BrandName          string    `json:"brandName"`
	CompanyName        string    `json:"companyName"`
	CompanyType        string    `json:"companyType"`
	CompanyLogoURL     string    `json:"companyLogoUrl"`
	RegistrationNumber string    `json:"registrationNumber"`
	BusinessCategory   string    `json:"businessCategory"`
	BusinessScope      string    `json:"businessScope"`
	SourceOfFunds      string    `json:"sourceOfFunds"`
	CustomerOrigin     string    `json:"customerOrigin"`
	WebsiteURL         string    `json:"websiteUrl"`
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

// MerchantSubscription :
type MerchantSubscription struct {
	ID          uint64    `json:"id"`
	GracePeriod uint64    `json:"gracePeriod"`
	ExpiryAt    time.Time `json:"expiryAt"`
	TerminateAt time.Time `json:"terminateAt"`
	Status      string    `json:"status"`
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

// User :
type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	CountryCode string    `json:"countryCode"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	AvatarURL   string    `json:"avatarUrl"`
	Status      string    `json:"status"`
	StoreID     string    `json:"storeId"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Voucher :
type Voucher struct {
	Key                string    `json:"key"`
	Label              string    `json:"label"`
	VoucherBatchKey    string    `json:"voucherBatchKey"`
	Type               string    `json:"type"`
	Amount             int       `json:"amount"`
	DiscountRate       int       `json:"discountRate"`
	MinimumSpendAmount int       `json:"minimumSpendAmount"`
	Origin             string    `json:"origin"`
	ImageURL           string    `json:"imageUrl"`
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
}

// VoucherBatch :
type VoucherBatch struct {
	ID                 string `json:"id"`
	Key                string `json:"key"`
	Label              string `json:"label"`
	Type               string `json:"type"`
	Amount             int    `json:"amount"`
	DiscountRate       int    `json:"discountRate"`
	MinimumSpendAmount int    `json:"minimumSpendAmount"`
	ImageURL           string `json:"imageUrl"`
	Quantity           int    `json:"quantity"`
	BalanceQuantity    int    `json:"balanceQuantity"`
	UsedQuantity       int    `json:"usedQuantity"`
	Status             string `json:"status"`
	Expiry             struct {
		Type      string    `json:"type"`
		Day       int       `json:"day"`
		ExpiredAt time.Time `json:"expiredAt"`
	} `json:"expiry"`
	Origin         string    `json:"origin"`
	IsShipping     bool      `json:"isShipping"`
	IsStatic       bool      `json:"isStatic"`
	StaticCode     string    `json:"staticCode"`
	IsDeviceRedeem bool      `json:"isDeviceRedeem"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
