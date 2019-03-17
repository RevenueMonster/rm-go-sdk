package sdk

import (
	"encoding/json"
	"errors"
)

// RequestGetRMWeChatUserOAuthURL :
type RequestGetRMWeChatUserOAuthURL struct {
	RedirectURL string `json:"redirectUrl"`
	Scope       string `json:"scope"`
}

// ResponseGetRMWeChatUserOAuthURL :
type ResponseGetRMWeChatUserOAuthURL struct {
	Item struct {
		URL string `json:"url"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// GetRMWeChatUserOAuthURL :
func (c Client) GetRMWeChatUserOAuthURL(request RequestGetRMWeChatUserOAuthURL) (*ResponseGetRMWeChatUserOAuthURL, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetRMWeChatUserOAuthURL.method
	requestURL := c.prepareAPIURL(pathAPIGetRMWeChatUserOAuthURL)

	response := new(ResponseGetRMWeChatUserOAuthURL)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// RequestGetRMWeChatUserInfoByCode :
type RequestGetRMWeChatUserInfoByCode struct {
	Code string `json:"code"`
}

// ReponseGetRMWeChatUserInfoByCode :
type ReponseGetRMWeChatUserInfoByCode struct {
	Item struct {
		UserID    string `json:"userId"`
		AvatarURL string `json:"avatarUrl"`
		Name      string `json:"name"`
		Gender    string `json:"gender"`
		City      string `json:"city"`
		State     string `json:"state"`
		Country   string `json:"country"`
		Language  string `json:"language"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// GetRMWeChatUserInfoByCode :
func (c Client) GetRMWeChatUserInfoByCode(request RequestGetRMWeChatUserInfoByCode) (*ReponseGetRMWeChatUserInfoByCode, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetRMWeChatUserInfoByCodeURL.method
	requestURL := c.prepareAPIURL(pathAPIGetRMWeChatUserInfoByCodeURL)

	response := new(ReponseGetRMWeChatUserInfoByCode)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// RequestSendWeChatPageTemplateMessage :
type RequestSendWeChatPageTemplateMessage struct {
	UserID      string `json:"userId"`
	TemplateID  string `json:"templateId"`
	URL         string `json:"url"`
	MiniProgram struct {
		AppID    string `json:"appId"`
		PagePath string `json:"pagePath"`
	} `json:"miniProgram"`
	Data json.RawMessage `json:"data"`
}

// ReponseSendWeChatPageTemplateMessage :
type ReponseSendWeChatPageTemplateMessage struct {
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// SendWeChatPageTemplateMessage :
func (c Client) SendWeChatPageTemplateMessage(request RequestSendWeChatPageTemplateMessage) (*ReponseSendWeChatPageTemplateMessage, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPISendWeChatPageTemplateMessageURL.method
	requestURL := c.prepareAPIURL(pathAPISendWeChatPageTemplateMessageURL)

	response := new(ReponseSendWeChatPageTemplateMessage)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// ResponseGetWeChatPageAccessToken :
type ResponseGetWeChatPageAccessToken struct {
	Item struct {
		AccessToken string `json:"accessToken"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// GetWeChatPageAccessToken :
func (c Client) GetWeChatPageAccessToken() (*ResponseGetWeChatPageAccessToken, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetWeChatPageAccessTokenURL.method
	requestURL := c.prepareAPIURL(pathAPIGetWeChatPageAccessTokenURL)

	response := new(ResponseGetWeChatPageAccessToken)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
