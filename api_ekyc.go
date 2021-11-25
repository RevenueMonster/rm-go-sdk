package sdk

import (
	"errors"
	"fmt"
)

type RequestEkycMykad struct {
	Image string `json:"query_image_content"`
}

type RequestEkycFaceCompare struct {
	Image1 string `json:"query_image_content_1"`
	Image2 string `json:"query_image_content_2"`
}

type RequestEkycLiveness struct {
	Image          string `json:"query_image_content"`
	MykadRequestID string `json:"mykad_request_id"`
}

type ResponseEkycMykad struct {
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

func (c Client) EkycMyKad(request RequestEkycMykad) (*ResponseEkycMykad, error) {
	if c.err != nil {
		return nil, c.err
	}

	response := new(ResponseEkycMykad)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, fmt.Sprintf("%s", requestURL), RequestService{
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

func (c Client) EkycFaceCompare(request RequestEkycFaceCompare) (*ResponseEkycFaceCompare, error) {
	if c.err != nil {
		return nil, c.err
	}

	response := new(ResponseEkycFaceCompare)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, fmt.Sprintf("%s", requestURL), RequestService{
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

func (c Client) EkycLiveness(request RequestEkycLiveness) (*ResponseEkycLiveness, error) {
	if c.err != nil {
		return nil, c.err
	}

	response := new(ResponseEkycLiveness)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, fmt.Sprintf("%s", requestURL), RequestService{
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
