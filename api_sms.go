package sdk

// SmsType :
type SmsType string

// SMS -> Type :
const (
	SmsTypeTAC SmsType = "TAC"
)

// RequestSendSms :
type RequestSendSms struct {
	CountryCode string  `json:"countryCode"`
	PhoneNumber string  `json:"phoneNumber"`
	Message     string  `json:"message"`
	Type        SmsType `json:"type"`
	Gateway     string  `json:"gateway"`
}

// ResponseSendSms :
type ResponseSendSms struct {
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// SendSms :
func (c Client) SendSms(request RequestSendSms) (*ResponseSendSms, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathSendSms.method
	requestURL := c.prepareAPIURL(pathSendSms)

	response := new(ResponseSendSms)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
