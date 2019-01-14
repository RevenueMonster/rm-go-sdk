package sdk

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	api "github.com/myussufz/fasthttp-api"
)

var cacheToken map[string]Token

// Client :
type Client struct {
	ID        string
	Secret    string
	IsSandbox bool
	// IsCache      bool
	OAuthVersion string
	APIVersion   string
	SignType     string
	PublicKey    []byte
	PrivateKey   []byte
	oauthURL     string
	apiURL       string
	AccessToken  string
	err          error
}

// Token :
type Token struct {
	ExpiredDateTime time.Time `json:"expiredAt"`
	AccessToken     string    `json:"accessToken"`
	TokenType       string    `json:"tokenType"`
	ExpiresIn       int       `json:"expiresIn"`
	RefreshToken    string    `json:"refreshToken"`
	Err             *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// NewClient : Init a client
func NewClient(c Client) Client {
	// set default if empty
	c.setDefault()

	// get the auth url
	c.oauthURL = fmt.Sprintf("%s/%s", c.getOAuthURL(), c.OAuthVersion)

	// get api url
	c.apiURL = fmt.Sprintf("%s/%s", c.getAPIURL(), c.APIVersion)

	return c
}

func (c *Client) httpAPI(method, requestURL string, request, response interface{}) error {
	if reflect.ValueOf(response).Kind() != reflect.Ptr {
		return errors.New("response struct should be pointer")
	}

	// generate signature
	signature, err := c.generateSignature(request, method, requestURL)
	if err != nil {
		return err
	}

	return api.Fetch(requestURL, api.Option{
		Method: method,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", c.AccessToken),
			"Content-Type":  "application/json",
			"X-Nonce-Str":   signature.NonceStr,
			"X-Signature":   signature.Sign,
			"X-Timestamp":   signature.Timestamp,
		},
		Body: request,
	}).ToJSON(response)
}

func (c *Client) prepareOAuthURL(p path) string {
	return fmt.Sprintf("%s%s", c.oauthURL, p.urlPath)
}

func (c *Client) prepareAPIURL(p path) string {
	return fmt.Sprintf("%s%s", c.apiURL, p.urlPath)
}

func (c *Client) setDefault() {
	if c.OAuthVersion == "" {
		c.OAuthVersion = defaultOAuthVersion
	}

	if c.APIVersion == "" {
		c.APIVersion = defaultAPIVersion
	}

	if c.SignType == "" {
		c.SignType = defaultSignType
	}
}

func (c *Client) getOAuthURL() string {
	url := fmt.Sprintf("%s", productionOAuthURL)

	if c.IsSandbox {
		url = fmt.Sprintf("%s", sandboxOAuthURL)
	}

	return url
}

// private function
func (c *Client) getAPIURL() string {
	url := fmt.Sprintf("%s", productionAPIURL)

	if c.IsSandbox {
		url = fmt.Sprintf("%s", sandboxAPIURL)
	}

	return url
}
