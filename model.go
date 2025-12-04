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
	ExtraInfo     struct {
		ExtraFee []struct {
			Type            string `json:"type"`
			ReferenceId     string `json:"referenceId"`
			Amount          uint   `json:"amount"`
			IsRefundAllowed bool   `json:"isRefundAllowed"`
			FeeType         string `json:"feeType"`
			FeeValue        uint   `json:"feeValue"`
			IsIncludedMDR   bool   `json:"isIncludedMDR"`
		} `json:"extraFee"`
	} `json:"extraInfo"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TokenizedPaymentOrder struct {
	ID                  string    `json:"id"` // order ID
	PaymentURL          string    `json:"paymentUrl"`
	RecurringID         string    `json:"recurringPaymentId"`
	RecurringCustomerID string    `json:"recurringCustomerId"`
	RecurringReference  string    `json:"recurringReference"`
	MerchantID          string    `json:"merchantId"`
	StoreID             string    `json:"storeId"`
	Amount              int       `json:"amount"`
	CurrencyType        string    `json:"currency"`
	CheckoutID          string    `json:"checkoutId"`
	TransactionID       string    `json:"transactionId"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
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
	Key                *string     `json:"key"`
	Label              string      `json:"label"`
	RedemptionRuleKey  interface{} `json:"redemptionRuleKey"`
	VoucherBatchKey    *string     `json:"voucherBatchKey"`
	Type               string      `json:"type"`
	Amount             int         `json:"amount"`
	DiscountRate       int         `json:"discountRate"`
	MinimumSpendAmount int         `json:"minimumSpendAmount"`
	Origin             string      `json:"origin"`
	ImageURL           string      `json:"imageUrl"`
	MemberProfile      interface{} `json:"memberProfile"`
	RedemptionRule     interface{} `json:"redemptionRule"`
	AssignedAt         *time.Time  `json:"assignedAt"`
	Payload            interface{} `json:"payload"`
	QrURL              string      `json:"qrUrl"`
	Code               string      `json:"code"`
	IsShipping         bool        `json:"isShipping"`
	Address            string      `json:"address"`
	Expiry             struct {
		Type      string  `json:"type"`
		Day       int     `json:"day"`
		ExpiredAt *string `json:"expiredAt"`
	} `json:"expiry"`
	VoucherComboKey *string    `json:"voucherComboKey"`
	UsedAt          *time.Time `json:"usedAt"`
	RedeemedAt      *time.Time `json:"redeemedAt"`
	IsDeviceRedeem  bool       `json:"isDeviceRedeem"`
	Status          string     `json:"status"`
	CreatedAt       *time.Time `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
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

type LoyaltyMember struct {
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

// MerchantSettlement :
type MerchantSettlement struct {
	ID                      string          `json:"id" goloquent:"__key__"`
	MerchantID              string          `json:"merchantId"`
	IsDefault               bool            `json:"isDefault"`
	Name                    string          `json:"companyName"`
	CompanyType             string          `json:"companyType"`
	RegistrationNumber      string          `json:"registrationNumber"`
	BusinessCategory        string          `json:"businessCategory"`
	BusinessScope           string          `json:"businessScope" goloquent:",longtext"`
	SourceOfFund            string          `json:"sourceOfFunds" goloquent:",longtext"`
	CustomerOrigin          string          `json:"customerOrigin" goloquent:",longtext"`
	EstablishedDateTime     string          `json:"establishedAt"`
	CountryCode             string          `json:"countryCode"`
	PhoneNumber             string          `json:"phoneNumber"`
	AddressLine1            string          `json:"addressLine1"`
	AddressLine2            string          `json:"addressLine2"`
	Postcode                string          `json:"postcode"`
	City                    string          `json:"city"`
	State                   string          `json:"state"`
	Country                 string          `json:"country"`
	IsSameBusinessAddress   bool            `json:"isSameBusinessAddress"`
	InvoiceAddress          *InvoiceAddress `json:"invoiceAddress"`
	InspectList             []InspectList   `json:"inspectList"`
	Status                  string          `json:"status"`
	Document                *Document       `json:"document"`
	DocumentFile            DocumentFile    `json:"documentFile"`
	BankAccountNo           string          `json:"bankAccountNo"`
	BankAccountType         string          `json:"bankAccountType"`
	BankAccountHolderName   string          `json:"bankAccountHolderName"`
	BankName                string          `json:"bankName" goloquent:"-"`
	BankCode                string          `json:"bankCode"`
	AverageTicketSize       uint64          `json:"averageTicketSize"`
	AverageTurnoverPerMonth uint64          `json:"averageTurnoverPerMonth"`
	PaymentSubscription     string          `json:"paymentSubscription"`
	AgentRate               uint64          `json:"agentRate"`
	CreatedDateTime         string          `json:"createdAt"`
	UpdatedDateTime         string          `json:"updatedAt"`
	TaxInfo                 TaxInfo         `json:"taxInfo"`
	HasTaxInfo              bool            `json:"hasTaxInfo"`
}

type TaxInfo struct {
	Tin      string `json:"tin"`
	SstRegNo string `json:"sstRegNo"`
	TtxRegNo string `json:"ttxRegNo"`
	Msic     string `json:"msic"`
}

// InvoiceAddress :
type InvoiceAddress struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	Postcode     string `json:"postcode"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}

// InspectList :
type InspectList struct {
	FullName      string `json:"fullName"`
	Email         string `json:"email"`
	Position      string `json:"position"`
	CountryCode   string `json:"countryCode"`
	PhoneNumber   string `json:"phoneNumber"`
	Gender        string `json:"gender"`
	Birthday      string `json:"birthday"`
	Nationality   string `json:"nationality"`
	IDType        string `json:"idType"`
	IDNo          string `json:"idNo"`
	BeginDateTime string `json:"beginAt"`
	EndDateTime   string `json:"endAt"`
}

type Document struct {
	CTOSFileURL                      string `json:"ctosFileUrl"`
	OwnerICFileURL                   string `json:"ownerICFileUrl"`
	DirectorICFileURL                string `json:"directorICFileUrl"`
	ShareHolderICFileURL             string `json:"shareHolderICFileUrl"`
	BusinessRegistrationFileURL      string `json:"businessRegistrationFileUrl"`
	BankStatementFileURL             string `json:"bankStatementFileUrl"`
	MOAFileURL                       string `json:"moaFileUrl"`
	Form24FileURL                    string `json:"form24FileUrl"`
	Form49FileURL                    string `json:"form49FileUrl"`
	Section14FileURL                 string `json:"section14FileUrl"`
	Form44FileURL                    string `json:"form44FileUrl"`
	BusinessSitePhotoFileURL         string `json:"businessSitePhotoFileUrl"`
	ESSMDocumenFileURL               string `json:"essmDocumentFileUrl"`
	LetterOfConsentFileURL           string `json:"letterOfConsentFileUrl"`
	CertificateOfRegistrationFileURL string `json:"certificateOfRegistrationFileUrl"`
	MeetingMinutesFileURL            string `json:"meetingMinutesFileUrl"`
	OrganizationChartFileURL         string `json:"organizationChartFileUrl"`
}

// DocumentFile : one of the struct in Merchant
type DocumentFile struct {
	CTOSFileURL                 []string
	OwnerICFileURL              []string
	DirectorICFileURL           []string
	ShareHolderICFileURL        []string
	BusinessRegistrationFileURL []string
	BankStatementFileURL        []string
	MOAFileURL                  []string
	Form24FileURL               []string
	Form49FileURL               []string
	Section14FileURL            []string
	Form44FileURL               []string
	BusinessSitePhotoFileURL    []string
}
