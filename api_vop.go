package sdk

import "fmt"

type RequestVOPEnrollUser struct {
	UserID    string `json:"userId"`
	CardPan   string `json:"cardPan"`
	Name      string `json:"name"`
	NotifyURL string `json:"notifyUrl"`
}

type ResponseVOPUserItem struct {
	Item VOPEnrollUser `json:"item"`
	Code string        `json:"code"`
}
type VOPEnrollUser struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	PanLastFour     string `json:"panLastFour"`
	Type            string `json:"type"`
	CreatedDateTime string `json:"createdAt"`
	UpdatedDateTime string `json:"updatedAt"`
}

type RequestVOPUnenrollUser struct {
	UserID string `json:"userId"`
}

type RequestVOPWebhook struct {
	UserID               string `json:"userId"`
	CardID               string `json:"cardId"`
	TransactionID        string `json:"transactionId"`
	TransactionDateTime  string `json:"transactionAt"`
	Amount               uint64 `json:"amount"`
	Currency             string `json:"currency"`
	MerchantCategoryCode string `json:"merchantCategoryCode"`
}

type ResponseVOPItem struct {
	Code string `json:"code"`
}

// VOPEnrollUser :
func (c Client) VOPEnrollUser(request RequestVOPEnrollUser) (*ResponseVOPItem, error) {

	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIVOPEnrollUserURL.method
	requestURL := c.prepareAPIURL(pathAPIVOPEnrollUserURL)

	fmt.Println(requestURL)

	response := new(ResponseVOPItem)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}

// VOPUnenrollUser :
func (c Client) VOPUnenrollUser(request RequestVOPUnenrollUser) (*ResponseVOPItem, error) {

	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIVOPUnenrollUserURL.method
	requestURL := c.prepareAPIURL(pathAPIVOPUnenrollUserURL)

	response := new(ResponseVOPItem)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}

// VOPWebhook :
func (c Client) VOPWebhook(request RequestVOPWebhook) (*ResponseVOPItem, error) {

	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIVOPWebhookURL.method
	requestURL := fmt.Sprintf("%s%s", c.getAPIURL(), pathAPIVOPWebhookURL.urlPath)

	response := new(ResponseVOPItem)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
