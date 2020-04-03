package sdk

import "fmt"

// RequestCreateDNSRecord :
type RequestCreateDNSRecord struct {
	RootDomain    string `json:"rootDomain"`
	SubDomainName string `json:"subDomainName"`
	IP            string `json:"ip"`
	Kubernetes    struct {
		Cluster   string `json:"cluster"`
		Service   string `json:"service"`
		Namespace string `json:"namespace"`
		Port      int    `json:"port"`
	} `json:"kubernetes"`
}

// ResponseCreateDNSRecord :
type ResponseCreateDNSRecord struct {
	Item struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		IP    string `json:"ip"`
		Proxy bool   `json:"proxy"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// CreateDNSRecord :
func (c Client) CreateDNSRecord(request RequestCreateDNSRecord) (*ResponseCreateDNSRecord, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathCreateDNSRecordURL.method
	requestURL := c.prepareAPIURL(pathCreateDNSRecordURL)

	response := new(ResponseCreateDNSRecord)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}

// RequestGetDNSRecords :
type RequestGetDNSRecords struct {
	RootDomain    string `json:"rootDomain"`
	SubDomainName string `json:"subDomainName"`
}

// ResponseGetDNSRecords :
type ResponseGetDNSRecords struct {
	Items []struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		IP    string `json:"ip"`
		Proxy bool   `json:"proxy"`
	} `json:"items"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// GetDNSRecords :
func (c Client) GetDNSRecords(rootDomain, subDomainName string) (*ResponseGetDNSRecords, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathGetDNSRecordsURL.method
	requestURL := c.prepareAPIURL(pathGetDNSRecordsURL) + fmt.Sprintf("?rootDomain=%s&subDomainName=%s", rootDomain, subDomainName)

	response := new(ResponseGetDNSRecords)
	if err := c.httpAPI(method, requestURL, nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// RequestDeleteDNSRecord :
type RequestDeleteDNSRecord struct {
	RootDomain    string `json:"rootDomain"`
	SubDomainName string `json:"subDomainName"`
	Kubernetes    struct {
		Cluster   string `json:"cluster"`
		Service   string `json:"service"`
		Namespace string `json:"namespace"`
		Port      int    `json:"port"`
	} `json:"kubernetes"`
}

// ResponseDeleteDNSRecord :
type ResponseDeleteDNSRecord struct {
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// DeleteDNSRecord :
func (c Client) DeleteDNSRecord(request RequestDeleteDNSRecord) (*ResponseDeleteDNSRecord, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathDeleteDNSRecordURL.method
	requestURL := c.prepareAPIURL(pathDeleteDNSRecordURL)

	response := new(ResponseDeleteDNSRecord)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
