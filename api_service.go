package sdk

import "encoding/json"

// RequestServiceWebhook :
type RequestServiceWebhook struct {
	Function string `json:"function"`
	Header   struct {
		ClientID         string   `json:"clientId"`
		NotificationURLs []string `json:"notificationUrls"`
	} `json:"header"`
	Data interface{} `json:"data"`
}

// ResponseServiceWebhook :
type ResponseServiceWebhook struct {
	Code string `json:"code"`
}

// RequestService :
type RequestService struct {
	Service  string      `json:"service"`
	Version  string      `json:"version"`
	Function string      `json:"function"`
	Request  interface{} `json:"request"`
}

// ResponseService :
type ResponseService struct {
	Code  string          `json:"code"`
	Item  json.RawMessage `json:"item"`
	Items json.RawMessage `json:"items"`
	Err   *Error          `json:"error"`
}

// ServiceWebhook :
func (c Client) ServiceWebhook(request RequestServiceWebhook) (*ResponseServiceWebhook, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIServiceWebhookURL.method
	requestURL := c.prepareAPIURL(pathAPIServiceWebhookURL)

	response := new(ResponseServiceWebhook)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}

// RequestService :
func (c Client) RequestService(request RequestService) (*ResponseService, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIService.method
	requestURL := c.prepareAPIURL(pathAPIService)

	response := new(ResponseService)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
