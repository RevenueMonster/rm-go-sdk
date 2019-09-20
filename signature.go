package sdk

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/clbanning/mxj"
)

// Signature :
type Signature struct {
	SignType  string
	Sign      string
	NonceStr  string
	Timestamp string
	Data      string
}

var signTypeMapper = map[string]crypto.Hash{
	SHA256: crypto.SHA256,
}

// GenerateDataBase64 :
func GenerateDataBase64(jsonData []byte) (string, error) {
	bodyBuffer := new(bytes.Buffer)

	if jsonData != nil {
		m, err := mxj.NewMapJson(jsonData)
		sortJSON, err := m.Json()
		if err = json.Compact(bodyBuffer, sortJSON); err != nil {
			return "", err
		}
	}

	bodyStr := strings.TrimSpace(bodyBuffer.String())
	bodyByte := bodyBuffer.Bytes()

	base64Body := ""

	switch bodyStr {
	case "{}":
	case "null":
	default:
		base64Body = base64.StdEncoding.EncodeToString(bodyByte)
	}

	return base64Body, nil
}

func (c *Client) generateSignature(data interface{}, method, requestURL string) (*Signature, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	bodyBuffer := new(bytes.Buffer)

	if data != nil {
		m, err := mxj.NewMapJson(jsonData)
		sortJSON, err := m.Json()
		if err = json.Compact(bodyBuffer, sortJSON); err != nil {
			return nil, err
		}
	}

	bodyStr := strings.TrimSpace(bodyBuffer.String())
	bodyByte := bodyBuffer.Bytes()

	base64Body := ""

	switch bodyStr {
	case "{}":
	case "null":
	default:
		base64Body = base64.StdEncoding.EncodeToString(bodyByte)
	}

	return Generate(
		strconv.FormatInt(int64(time.Now().UTC().Unix()), 10),
		randomString(32),
		method,
		requestURL,
		c.SignType,
		base64Body,
		c.PrivateKey,
	)
}

// Generate :
func Generate(timestamp, nonceStr, method, requestURL, signType, body string, privateKey []byte) (*Signature, error) {
	method = strings.ToLower(method)
	h, isExist := signTypeMapper[signType]
	if !isExist {
		return nil, errors.New("sign type not found")
	}

	data, err := formData(timestamp, nonceStr, method, requestURL, signType, body)
	if err != nil {
		return nil, err
	}

	signature, err := rsa2Hash(h, []byte(data), privateKey)
	if err != nil {
		return nil, err
	}

	return &Signature{
		SignType:  signType,
		Sign:      fmt.Sprintf("%s %s", signType, signature),
		NonceStr:  nonceStr,
		Timestamp: timestamp,
		Data:      data,
	}, nil
}

// Validate :
func Validate(timestamp, nonceStr, method, requestURL, signature, base64Body string, publicKey []byte) bool {
	method = strings.ToLower(method)
	signArr := strings.Split(signature, " ")

	if len(signArr) != 2 {
		return false
	}

	signType := signArr[0]
	signature = signArr[1]

	h, isExist := signTypeMapper[signType]
	if !isExist {
		return false
	}

	data, err := formData(timestamp, nonceStr, method, requestURL, signType, base64Body)
	if err != nil {
		return false
	}

	if err := verifyPKCS1v15(h, publicKey, []byte(data), signature); err != nil {
		return false
	}

	return true
}

type signForm struct {
	Data       string `json:"data,omitempty"`
	Method     string `json:"method"`
	NonceStr   string `json:"nonceStr"`
	RequestURL string `json:"requestUrl"`
	SignType   string `json:"signType"`
	Timestamp  string `json:"timestamp"`
}

func formData(timestamp, nonceStr, method, requestURL, signType, body string) (string, error) {
	s := signForm{
		Data:       body,
		Timestamp:  timestamp,
		NonceStr:   nonceStr,
		Method:     method,
		RequestURL: requestURL,
		SignType:   signType,
	}

	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	data, err := sortJSONByKey(jsonData)
	if err != nil {
		return "", err
	}

	return data, nil
}

func rsa2Hash(h crypto.Hash, data, privateKey []byte) (string, error) {
	signature, err := signPKCS1v15(h, data, privateKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func signPKCS1v15(hash crypto.Hash, src, key []byte) ([]byte, error) {
	var h = hash.New()
	if _, err := h.Write(src); err != nil {
		return nil, err
	}

	hashed := h.Sum(nil)

	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error")
	}

	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, pk, hash, hashed)
}

func verifyPKCS1v15(ch crypto.Hash, publicKey []byte, data []byte, signature string) error {
	var hashed []byte

	switch ch {
	case crypto.SHA256:
		var h = sha256.New()
		h.Write(data)
		hashed = h.Sum(nil)
	}

	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}

	var pubInterface interface{}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	var pub = pubInterface.(*rsa.PublicKey)

	signByte, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed, signByte)
}

func sortJSONByKey(b []byte) (string, error) {
	m, err := mxj.NewMapJson(b)
	if err != nil {
		return "", err
	}

	m = m.Old()
	arr := make([]string, 0)
	for k, v := range m {
		if strings.TrimSpace(strings.ToLower(k)) == "sign" {
			continue
		}

		str := ""
		switch vi := v.(type) {
		case map[string]interface{}:
			continue
		case string:
			if vi == "" || vi == "{}" {
				continue
			}
			str = vi
		default:
			str = fmt.Sprintf("%v", vi)
		}
		arr = append(arr, fmt.Sprintf("%s=%s", k, str))
	}

	sort.Sort(sort.StringSlice(arr))

	return strings.Join(arr, "&"), nil
}
