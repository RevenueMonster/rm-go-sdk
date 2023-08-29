package sdk

import "errors"

// RequestCreatePaymentQuickPay :
type RequestCreatePaymentQuickPay struct {
	StoreID   string `json:"storeId"`
	AuthCode  string `json:"authCode"`
	IPAddress string `json:"ipAddress"`
	Order     struct {
		ID             string `json:"id"`
		Title          string `json:"title"`
		CurrencyType   string `json:"currencyType"`
		Amount         int64  `json:"amount"`
		Detail         string `json:"detail"`
		AdditionalData string `json:"additionalData"`
	} `json:"order"`
	Voucher *RequestCreatePaymentQuickPayVoucher `json:"voucher,omitempty"`
}

// RequestCreatePaymentQuickPayVoucher :
type RequestCreatePaymentQuickPayVoucher struct {
	Code string `json:"code"`
}

// ResponseCreatePaymentQuickPay :
type ResponseCreatePaymentQuickPay struct {
	Item PaymentTransaction `json:"item"`
	Code string             `json:"code"`
	Err  *Error             `json:"error"`
}

// CreatePaymentQuickPay :
func (c Client) CreatePaymentQuickPay(request RequestCreatePaymentQuickPay) (*ResponseCreatePaymentQuickPay, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPICreatePaymentQuickPay.method
	requestURL := c.prepareAPIURL(pathAPICreatePaymentQuickPay)

	response := new(ResponseCreatePaymentQuickPay)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

type RequestCreateTerminalPaymentType string

const (
	RequestCreateTerminalPaymentTypeWallet = "E-WALLET"
	RequestCreateTerminalPaymentTypeCard   = "CARD"
)

type RequestCreateTerminalPaymentCameraType string

const (
	RequestCreateTerminalPaymentCameraTypeFront = "FRONT"
	RequestCreateTerminalPaymentCameraTypeBack  = "BACK"
)

// RequestCreateTerminalPayment :
type RequestCreateTerminalPayment struct {
	TerminalID  string                                 `json:"terminalId"`
	Type        RequestCreateTerminalPaymentType       `json:"type"`
	CameraType  RequestCreateTerminalPaymentCameraType `json:"cameraType"`
	ReceiptType int                                    `json:"receiptType"`
	Order       struct {
		ID             string `json:"id"`
		Title          string `json:"title"`
		CurrencyType   string `json:"currencyType"`
		Amount         int64  `json:"amount"`
		Detail         string `json:"detail"`
		AdditionalData string `json:"additionalData"`
	} `json:"order"`
}

// ResponseCreatePaymentQuickPay :
type ResponseCreateTerminalPayment struct {
	Item PaymentTransaction `json:"item"`
	Code string             `json:"code"`
	Err  *Error             `json:"error"`
}

// CreateTerminalPayment :
func (c Client) CreateTerminalPayment(request RequestCreateTerminalPayment) (*ResponseCreateTerminalPayment, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPICreateTerminalPayment.method
	requestURL := c.prepareAPIURL(pathAPICreateTerminalPayment)

	response := new(ResponseCreateTerminalPayment)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
