package sdk

import (
	"errors"
	"fmt"
)

// NotificationEvent :
type NotificationEvent struct {
	Type        string `json:"type"`
	ReferenceID string `json:"referenceId"`
	Data        string `json:"data"`
}

// Notification :
type Notification struct {
	Title             string            `json:"title"`
	Description       string            `json:"description"`
	NotificationEvent NotificationEvent `json:"event"`
}

// SendEventRequest :
type SendEventRequest struct {
	ReferenceID  string       `json:"referenceId"`
	Platform     string       `json:"platform"`
	Channel      string       `json:"channel"`
	Action       string       `json:"action"`
	Notification Notification `json:"notification"`
	Body         string       `json:"body"`
}

// SendEventRequestResponse :
type SendEventRequestResponse struct {
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// SendEvent :
func (c Client) SendEvent(req SendEventRequest) (*SendEventRequestResponse, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIPostSendEventURL.method
	requestURL := c.prepareAPIURL(pathAPIPostSendEventURL)

	response := new(SendEventRequestResponse)
	if err := c.httpAPI(method, fmt.Sprintf("%s", requestURL), req, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
