package sdk

import (
	"errors"
)

type RequestEkycMykadDirect struct {
	Image     string `json:"query_image_content"`
	OCREngine string `json:"ocr_engine"`
}

type RequestEkycMykad struct {
	Image     string `json:"query_image_content"`
	NotifyUrl string `json:"notify_url"`
}

type RequestEkycFaceCompare struct {
	Image1 string `json:"query_image_content_1"`
	Image2 string `json:"query_image_content_2"`
}

type RequestEkycLiveness struct {
	Image          string `json:"query_image_content"`
	MykadRequestID string `json:"mykad_request_id"`
}

type RequestGetMykadResult struct {
	ID string `json:"id"`
}

type RequestGetEkycResult struct {
	ID       string   `json:"id"`
	Includes []string `json:"includes"`
}

type ResponseEkycMykadDirect struct {
	Code string `json:"code"`
	Item struct {
		ID      string `json:"id"`
		Status  string `json:"status"`
		Action  string `json:"action"`
		Objects struct {
			Duration float64             `json:"duration"`
			Data     []BoundingBoxResult `json:"data"`
		} `json:"objects,omitempty"`
		Results      []MykadResult `json:"results"`
		Duration     float64       `json:"duration"`
		VisImagePath string        `json:"visualizeImagePath"`
		CreatedAt    string        `json:"createdAt,omitempty"`
		UpdatedAt    string        `json:"updatedAt,omitempty"`
	} `json:"item"`
	Err *Error `json:"error"`
}

type ResponseEkycMykad struct {
	Code string `json:"code"`
	Item struct {
		RequestID string `json:"requestID"`
		Status    string `json:"status"`
	} `json:"item"`
	Err *Error `json:"error"`
}

type ResponseEkycFaceCompare struct {
	Code string            `json:"code"`
	Item FaceCompareResult `json:"item"`
	Err  *Error            `json:"error"`
}

type ResponseEkycLiveness struct {
	Code string            `json:"code"`
	Item FaceCompareResult `json:"item"`
	Err  *Error            `json:"error"`
}

type ResponseGetMykadResult struct {
	Code string `json:"code"`
	Item struct {
		ID      string `json:"id"`
		Status  string `json:"status"`
		Action  string `json:"action"`
		Objects struct {
			Duration float64             `json:"duration"`
			Data     []BoundingBoxResult `json:"data"`
		} `json:"objects,omitempty"`
		Results      []MykadResult `json:"results"`
		Duration     float64       `json:"duration"`
		VisImagePath string        `json:"visualizeImagePath"`
		CreatedAt    string        `json:"createdAt,omitempty"`
		UpdatedAt    string        `json:"updatedAt,omitempty"`
	} `json:"item"`
	Err *Error `json:"error"`
}

type ResponseGetEkycResult struct {
	Code string        `json:"code"`
	Item GetEkycResult `json:"item"`
	Err  *Error        `json:"error"`
}

// EkycMyKad :
func (c Client) EkycMyKadDirect(request RequestEkycMykadDirect) (*ResponseEkycMykadDirect, error) {
	if c.err != nil {
		return nil, c.err
	}

	response := new(ResponseEkycMykadDirect)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, requestURL, RequestService{
		Service:  "ekyc",
		Version:  "v1",
		Function: "mykad",
		Request:  request,
	}, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// EkycIDMyKad :
func (c Client) EkycMyKad(request RequestEkycMykad) (*ResponseEkycMykad, error) {
	if c.err != nil {
		return nil, c.err
	}

	response := new(ResponseEkycMykad)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, requestURL, RequestService{
		Service:  "ekyc",
		Version:  "v1",
		Function: "id-mykad",
		Request:  request,
	}, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// EkycFaceCompare :
func (c Client) EkycFaceCompare(request RequestEkycFaceCompare) (*ResponseEkycFaceCompare, error) {
	if c.err != nil {
		return nil, c.err
	}

	response := new(ResponseEkycFaceCompare)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, requestURL, RequestService{
		Service:  "ekyc",
		Version:  "v1",
		Function: "face-compare",
		Request:  request,
	}, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}
	return response, nil
}

// EkycLiveness :
func (c Client) EkycLiveness(request RequestEkycLiveness) (*ResponseEkycLiveness, error) {
	if c.err != nil {
		return nil, c.err
	}

	response := new(ResponseEkycLiveness)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, requestURL, RequestService{
		Service:  "ekyc",
		Version:  "v1",
		Function: "face-compare-with-mykad",
		Request:  request,
	}, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}
	return response, nil
}

// GetMykadResult :
func (c Client) GetMykadResult(request RequestGetMykadResult) (*ResponseGetMykadResult, error) {
	if c.err != nil {
		return nil, c.err
	}

	response := new(ResponseGetMykadResult)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, requestURL, RequestService{
		Service:  "ekyc",
		Version:  "v1",
		Function: "get-mykad-result",
		Request:  request,
	}, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}
	return response, nil
}

// GetEkycResult :
func (c Client) GetEkycResult(request RequestGetEkycResult) (*ResponseGetEkycResult, error) {
	if c.err != nil {
		return nil, c.err
	}

	response := new(ResponseGetEkycResult)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, requestURL, RequestService{
		Service:  "ekyc",
		Version:  "v1",
		Function: "get-ekyc-result",
		Request:  request,
	}, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}
	return response, nil
}
