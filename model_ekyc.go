package sdk

// BoundingBox :
type BoundingBox struct {
	X int32 `bson:"x" json:"x"`
	Y int32 `bson:"y" json:"y"`
}

// BoundingBoxResult :
type BoundingBoxResult struct {
	Prediction string        `bson:"prediction" json:"prediction"`
	Confidence float64       `bson:"confidence" json:"confidence"`
	Boxes      []BoundingBox `bson:"boxes" json:"boxes"`
}

// FaceDetResult :
type FaceDetResult struct {
	Confidence float64       `bson:"confidence" json:"confidence"`
	Boxes      []BoundingBox `bson:"boxes" json:"boxes"`
}

// OCRRequestIDs :
type OCRRequestIDs struct {
	Region1 string `bson:"region1" json:"region1,omitempty"`
	Region2 string `bson:"region2" json:"region2,omitempty"`
	Region3 string `bson:"region3" json:"region3,omitempty"`
}

// MykadScore :
type MykadScore struct {
	Landmarks float64 `bson:"landmarks" json:"landmarks"`
	OCR       float64 `bson:"ocr" json:"ocr"`
	Searcher  float64 `bson:"searcher" json:"searcher"`
	Total     float64 `bson:"total" json:"total"`
}

// MykadAddress :
type MykadAddress struct {
	City         string `bson:"city" json:"city"`
	State        string `bson:"state" json:"state"`
	Postcode     string `bson:"postcode" json:"postcode"`
	AddressLine1 string `bson:"addressLine1" json:"addressLine1"`
	AddressLine2 string `bson:"addressLine2" json:"addressLine2"`
	AddressLine3 string `bson:"addressLine3" json:"addressLine3"`
	AddressLine4 string `bson:"addressLine4" json:"addressLine4,omitempty"`
	AddressLine5 string `bson:"addressLine5" json:"addressLine5,omitempty"`
}

// MykadTCRegionResult :
type MykadTCRegionResult struct {
	Success bool   `bson:"success" json:"success"`
	Message string `bson:"message" json:"message"`
	Result  struct {
		Prediction string  `bson:"prediction" json:"prediction"`
		Confidence float64 `bson:"confidence" json:"confidence"`
		Alternated bool    `bson:"alternated" json:"alternated"`
	} `bson:"result" json:"result"`
}

// MykadResult :
type MykadResult struct {
	Scores MykadScore      `bson:"scores" json:"scores"`
	Faces  []FaceDetResult `bson:"faces" json:"faces"`
	Data   struct {
		IC       string       `bson:"ic" json:"ic"`
		Name     string       `bson:"name" json:"name"`
		Gender   string       `bson:"gender" json:"gender"`
		IsMuslim bool         `bson:"isMuslim" json:"isMuslim"`
		Address  MykadAddress `bson:"address" json:"address"`
	} `bson:"data" json:"data"`
	Object struct {
		Prediction string  `bson:"prediction" json:"prediction"`
		Confidence float64 `bson:"confidence" json:"confidence"`
	} `bson:"object" json:"object"`
	Colour struct {
		Detected   string  `bson:"detected" json:"detected"`
		Percentage float64 `bson:"percentage" json:"percentage"`
	} `bson:"colour" json:"colour"`
	Glare struct {
		Present bool `bson:"present" json:"present"`
		Details []struct {
			Region     int32         `bson:"region" json:"region"`
			Confidence float64       `bson:"confidence" json:"confidence"`
			Boxes      []BoundingBox `bson:"boxes" json:"boxes"`
		} `bson:"details" json:"details"`
	} `bson:"glare" json:"glare"`
	TamperCheck struct {
		Prediction string `bson:"prediction" json:"prediction"`
		Physical   struct {
			Success    bool    `bson:"success" json:"success"`
			Prediction string  `bson:"prediction" json:"prediction"`
			Confidence float64 `bson:"confidence" json:"confidence"`
			FailReason string  `bson:"failReason" json:"failReason,omitempty"`
		} `bson:"physical" json:"physical"`
		FaceVerification struct {
			Success bool   `bson:"success" json:"success"`
			Message string `bson:"message" json:"message"`
			Result  struct {
				Similarity   float64 `bson:"similarity" json:"similarity"`
				IsSamePerson bool    `bson:"isSamePerson" json:"isSamePerson"`
			} `bson:"result" json:"result"`
		} `bson:"faceVerification" json:"faceVerification"`
		Region1     MykadTCRegionResult `bson:"region1" json:"region1"`
		Region2     MykadTCRegionResult `bson:"region2" json:"region2"`
		Region3     MykadTCRegionResult `bson:"region3" json:"region3"`
		FrontalFace MykadTCRegionResult `bson:"frontalFace" json:"frontalFace"`
	} `bson:"tamperCheck" json:"tamperCheck"`
	Duration      float64       `bson:"duration" json:"duration"`
	OCRRequestIDs OCRRequestIDs `bson:"ocrRequestIds" json:"ocrRequestIds,omitempty"`
	Meta          struct {
		ResizeScale float64 `bson:"resize_scale" json:"resizeScale"`
	} `bson:"meta" json:"meta"`
}

// FaceCompareResult :
type FaceCompareResult struct {
	ID           string `json:"id"`
	Action       string `json:"action"`
	Similarity   uint8  `json:"similarity"`
	IsSamePerson bool   `json:"isSamePerson"`
	Status       string `json:"status"`
	UpdatedAt    string `json:"updatedAt"`
}
