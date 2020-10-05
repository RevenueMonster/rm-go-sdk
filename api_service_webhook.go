package sdk

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
