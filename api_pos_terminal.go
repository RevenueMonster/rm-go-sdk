package sdk

type PaymentType string
type ReceiptType uint
type CameraType string

const (
	PaymentTypeCard    PaymentType = "CARD"
	PaymentTypeEWallet PaymentType = "E-WALLET"

	ReceiptTypeMerchantAndCustomer ReceiptType = 1
	ReceiptTypeCustomer            ReceiptType = 2
	ReceiptTypeNone                ReceiptType = 3

	CameraTypeFront CameraType = "FRONT"
	CameraTypeBack  CameraType = "BACK"
)

type TerminalQuickPayOrder struct {
	Amount         uint   `json:"amount"`
	CurrencyType   string `json:"currencyType"`
	ID             string `json:"id"`
	Title          string `json:"title"`
	Details        string `json:"details"`
	AdditionalData string `json:"additionalData"`
}

type RequestTerminalQuickPay struct {
	TerminalID  string                `json:"terminalId"`
	Type        PaymentType           `json:"type"`
	ReceiptType ReceiptType           `json:"receiptType"`
	CameraType  CameraType            `json:"cameraType"`
	Order       TerminalQuickPayOrder `json:"order"`
}

type ResponseTerminalQuickPay struct {
	Item PaymentTransactionQRCode `json:"item"`
	Code string                   `json:"code"`
	Err  *Error                   `json:"error"`
}

// QuickPay :
func (c Client) QuickPay(request RequestTerminalQuickPay) (*ResponseTerminalQuickPay, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIQuickPay.method
	requestURL := c.prepareAPIURL(pathAPIQuickPay)

	response := new(ResponseTerminalQuickPay)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
