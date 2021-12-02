package sdk

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

type RequestGetEkycResult struct {
	ID       string   `json:"id"`
	Includes []string `json:"includes"`
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

type ResponseGetEkycResult struct {
	Code string        `json:"code"`
	Item GetEkycResult `json:"item"`
	Err  *Error        `json:"error"`
}

// EkycMyKad :
func (c Client) EkycMyKad(request RequestEkycMykad) (*ResponseEkycMykad, *Error) {
	if c.err != nil {
		return nil, &Error{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: c.err.Error(),
		}
	}

	response := new(ResponseEkycMykad)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, requestURL, RequestService{
		Service:  "ekyc",
		Version:  "v1",
		Function: "mykad",
		Request:  request,
	}, response); err != nil {
		return nil, &Error{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
		}
	}

	if response.Err != nil {
		return response, response.Err
	}

	return response, nil
}

// EkycFaceCompare :
func (c Client) EkycFaceCompare(request RequestEkycFaceCompare) (*ResponseEkycFaceCompare, *Error) {
	if c.err != nil {
		return nil, &Error{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: c.err.Error(),
		}
	}

	response := new(ResponseEkycFaceCompare)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, requestURL, RequestService{
		Service:  "ekyc",
		Version:  "v1",
		Function: "face-compare",
		Request:  request,
	}, response); err != nil {
		return nil, &Error{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
		}
	}

	if response.Err != nil {
		return response, response.Err
	}
	return response, nil
}

// EkycLiveness :
func (c Client) EkycLiveness(request RequestEkycLiveness) (*ResponseEkycLiveness, *Error) {
	if c.err != nil {
		return nil, &Error{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: c.err.Error(),
		}
	}

	response := new(ResponseEkycLiveness)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, requestURL, RequestService{
		Service:  "ekyc",
		Version:  "v1",
		Function: "face-compare-with-mykad",
		Request:  request,
	}, response); err != nil {
		return nil, &Error{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
		}
	}

	if response.Err != nil {
		return response, response.Err
	}
	return response, nil
}

// GetEkycResult :
func (c Client) GetEkycResult(request RequestGetEkycResult) (*ResponseGetEkycResult, *Error) {
	if c.err != nil {
		return nil, &Error{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: c.err.Error(),
		}
	}

	response := new(ResponseGetEkycResult)
	requestURL := c.prepareAPIURL(pathAPIService)
	if err := c.httpAPI(methodPOST, requestURL, RequestService{
		Service:  "ekyc",
		Version:  "v1",
		Function: "get-ekyc-result",
		Request:  request,
	}, response); err != nil {
		return nil, &Error{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
		}
	}

	if response.Err != nil {
		return response, response.Err
	}
	return response, nil
}
