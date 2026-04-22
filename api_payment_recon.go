package sdk

import (
	"errors"
	"time"
)

type RequestReconcilePayment struct {
	TransactionType string   `json:"transactionType"`
	Date            string   `json:"date"`
	Methods         []string `json:"methods"`
	Regions         []string `json:"regions"`
	Cursor          string   `json:"cursor,omitempty"`
}

type ResponseReconcilePaymentItem struct {
	TransactionAt    time.Time `json:"transactionAt"`
	MerchantID       string    `json:"merchantId"`
	MerchantName     string    `json:"merchantName"`
	StoreID          string    `json:"storeId"`
	StoreName        string    `json:"storeName"`
	Region           string    `json:"region"`
	Method           string    `json:"method"`
	TransactionType  string    `json:"transactionType"`
	Type             string    `json:"type"`
	TransactionID    string    `json:"transactionId"`
	OrderID          string    `json:"orderId"`
	CurrencyType     string    `json:"currencyType"`
	GrossAmount      string    `json:"grossAmount"`
	MDR              string    `json:"mdr"`
	ServiceFee       string    `json:"serviceFee"`
	SettlementAmount string    `json:"settlementAmount"`
}

type ResponseReconcilePayment struct {
	Items []ResponseReconcilePaymentItem `json:"items"`
	Meta  struct {
		Cursor string `json:"cursor"`
	} `json:"meta"`
	Err *Error `json:"error,omitempty"`
}

func (c Client) ReconcilePayment(request RequestReconcilePayment) (*ResponseReconcilePayment, error) {
	if c.err != nil {
		return nil, c.err
	}

	switch request.TransactionType {
	case "PAYMENT", "REFUND":
	default:
		return nil, errors.New("invalid transaction type")
	}

	method := pathAPIPaymentReconciliationURL.method
	requestURL := c.prepareAPIURL(pathAPIPaymentReconciliationURL)

	response := new(ResponseReconcilePayment)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}
