package sdk

import "fmt"

type VOPEnrollUser struct {
	ID              string `json:"id"`
	CreatedDateTime string `json:"createdAt"`
	UpdatedDateTime string `json:"updatedAt"`
}

type VOPEnrollCard struct {
	ID              string `json:"id"`
	UserID          string `json:"userId"`
	Name            string `json:"name"`
	PanLastFour     string `json:"panLastFour"`
	Type            string `json:"type"`
	CreatedDateTime string `json:"createdAt"`
	UpdatedDateTime string `json:"updatedAt"`
}

type RequestVOPEnrollUser struct {
	UserID    string `json:"userId"`
	CardPan   string `json:"cardPan"`
	Name      string `json:"name"`
	NotifyURL string `json:"notifyUrl"`
}

type RequestVOPEnrollCard struct {
	UserID  string `json:"userId"`
	CardPan string `json:"cardPan"`
	Name    string `json:"name"`
}

type ResponseVOPUserItem struct {
	Item struct {
		VOPEnrollCard VOPEnrollCard `json:"card"`
		VOPEnrollUser VOPEnrollUser `json:"user"`
	} `json:"item"`
	Code string `json:"code"`
}

type RequestVOPUnenrollUser struct {
	UserID string `json:"userId"`
}

type ResponseVOPCardItem struct {
	Item VOPEnrollCard `json:"item"`
	Code string        `json:"code"`
}

type RequestVOPUnenrollCard struct {
	UserID string `json:"userId"`
	CardID string `json:"cardId"`
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
func (c Client) VOPEnrollUser(request RequestVOPEnrollUser) (*ResponseVOPUserItem, error) {

	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIVOPEnrollUserURL.method
	requestURL := c.prepareAPIURL(pathAPIVOPEnrollUserURL)

	fmt.Println(requestURL)

	response := new(ResponseVOPUserItem)
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

// VOPEnrollCard :
func (c Client) VOPEnrollCard(request RequestVOPEnrollCard) (*ResponseVOPCardItem, error) {

	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIVOPEnrollCardURL.method
	requestURL := c.prepareAPIURL(pathAPIVOPEnrollCardURL)

	response := new(ResponseVOPCardItem)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}

// VOPUnenrollCard :
func (c Client) VOPUnenrollCard(request RequestVOPUnenrollCard) (*ResponseVOPItem, error) {

	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIVOPUnenrollCardURL.method
	requestURL := c.prepareAPIURL(pathAPIVOPUnenrollCardURL)

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
