package sdk

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	api "github.com/myussufz/fasthttp-api"
)

// ResponseGetTokenInfoByAuthorizationCode :
type ResponseGetTokenInfoByAuthorizationCode struct {
	Active   bool     `json:"active"`
	ClientID string   `json:"clientId"`
	Exp      uint64   `json:"exp"`
	Scopes   []string `json:"scopes"`
	Username string   `json:"username"`
}

// GetAuthorizationCodeURL :
func (c Client) GetAuthorizationCodeURL(state, redirectURL string, scopes ...string) string {
	return fmt.Sprintf("%s/authorize?responseType=code&clientId=%s&redirectUri=%s&scope=%s&state=%s", c.getOAuthURL(), c.ID, redirectURL, strings.Join(scopes, ","), state)
}

// GetAccessTokenByClientCredentials :
func (c Client) GetAccessTokenByClientCredentials() (*Token, error) {
	tk := new(Token)

	method := pathOAuthPathTokenURL.method
	requestURL := c.prepareOAuthURL(pathOAuthPathTokenURL)

	if err := api.Fetch(requestURL, api.Option{
		Method: method,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.ID, c.Secret)))),
		},
		Body: map[string]interface{}{
			"grantType": "client_credentials",
		},
	}).ToJSON(&tk); err != nil || tk.Err != nil {
		switch {
		case tk.Err != nil:
			return nil, errors.New(tk.Err.Message)

		default:
			return nil, err

		}
	}

	tk.ExpiredDateTime = time.Now().UTC().Add(time.Second * time.Duration(tk.ExpiresIn-60))

	return tk, nil
}

// GetAccessTokenByRefreshToken :
func (c Client) GetAccessTokenByRefreshToken(refreshToken string) (*Token, error) {
	tk := new(Token)

	method := pathOAuthPathTokenURL.method
	requestURL := c.prepareOAuthURL(pathOAuthPathTokenURL)

	if err := api.Fetch(requestURL, api.Option{
		Method: method,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.ID, c.Secret)))),
		},
		Body: map[string]interface{}{
			"grantType":    "refresh_token",
			"refreshToken": refreshToken,
		},
	}).ToJSON(&tk); err != nil || tk.Err != nil {
		switch {
		case tk.Err != nil:
			return nil, errors.New(tk.Err.Message)

		default:
			return nil, err
		}
	}

	tk.ExpiredDateTime = time.Now().UTC().Add(time.Second * time.Duration(tk.ExpiresIn-60))

	return tk, nil
}

// GetAccessTokenByAuthorizationCode :
func (c Client) GetAccessTokenByAuthorizationCode(code, redirectURL string) (*Token, error) {
	tk := new(Token)

	method := pathOAuthPathTokenURL.method
	requestURL := c.prepareOAuthURL(pathOAuthPathTokenURL)

	if err := api.Fetch(requestURL, api.Option{
		Method: method,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.ID, c.Secret)))),
		},
		Body: map[string]interface{}{
			"grantType":   "authorization_code",
			"code":        code,
			"redirectUri": redirectURL,
		},
	}).ToJSON(tk); err != nil || tk.Err != nil {
		switch {
		case tk.Err != nil:
			return nil, errors.New(tk.Err.Message)

		default:
			return nil, err

		}
	}

	tk.ExpiredDateTime = time.Now().UTC().Add(time.Second * time.Duration(tk.ExpiresIn-60))

	return tk, nil
}

// GetTokenInfoByAuthorizationCode :
func (c Client) GetTokenInfoByAuthorizationCode(code string) (*ResponseGetTokenInfoByAuthorizationCode, error) {
	response := new(ResponseGetTokenInfoByAuthorizationCode)

	method := pathOAuthPathTokenInfoURL.method
	requestURL := c.prepareOAuthURL(pathOAuthPathTokenInfoURL)

	if err := api.Fetch(requestURL, api.Option{
		Method: method,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.ID, c.Secret)))),
		},
		Body: map[string]interface{}{
			"token": code,
		},
	}).ToJSON(response); err != nil {
		return nil, err
	}

	return response, nil
}
