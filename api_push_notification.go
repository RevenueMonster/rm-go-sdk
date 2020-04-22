package sdk

import "fmt"

// RequestPushNotificationToStore :
type RequestPushNotificationToStore struct {
	StoreID string `json:"-"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Event   struct {
		Type        string `json:"type"`
		ReferenceID string `json:"referenceId"`
		Data        string `json:"data"`
	} `json:"event"`
}

// ResponsePushNotificationToStore :
type ResponsePushNotificationToStore struct {
	Code string `json:"code"`
}

// PushNotificationToStore :
func (c Client) PushNotificationToStore(request RequestPushNotificationToStore) (*ResponsePushNotificationToStore, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIPushNotificationToStoreURL.method
	requestURL := c.prepareAPIURL(pathAPIPushNotificationToStoreURL)

	response := new(ResponsePushNotificationToStore)
	if err := c.httpAPI(method, fmt.Sprintf("%s/%s", requestURL, request.StoreID), request, response); err != nil {
		return nil, err
	}

	return response, nil
}
