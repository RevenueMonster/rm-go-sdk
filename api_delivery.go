package sdk

import "fmt"

// Delivery :
type Delivery struct {
	DeliveryID       string   `json:"deliveryId"`
	DeliveryVendorID string   `json:"deliveryVendorId"`
	CurrencyType     string   `json:"currencyType"`
	Amount           uint64   `json:"amount"`
	Status           string   `json:"status"`
	Courier          *Courier `json:"courier"`
}

// Courier :
type Courier struct {
	ID        string  `json:"id"`
	Surname   string  `json:"surname"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	PhotoURL  string  `json:"photoUrl"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// DeliveryType :
type DeliveryType string

// DeliveryType -> Type
const (
	DeliveryTypeDocument DeliveryType = "DOCUMENT"
	DeliveryTypeFood     DeliveryType = "FOOD"
)

// VehicleType :
type VehicleType string

// Vehicle -> Type
const (
	VehicleTypeMotobike VehicleType = "MOTOBIKE"
	VehicleTypeCar      VehicleType = "CAR"
)

// DeliveryVendor :
type DeliveryVendor struct {
	Vendor string `json:"vendor"`
}

// DeliveryPoint :
type DeliveryPoint struct {
	Address        string               `json:"address"`
	EntranceNumber string               `json:"entranceNumber"`
	FloorNumber    string               `json:"floorNumber"`
	BuildingNumber string               `json:"buildingNumber"`
	Remark         string               `json:"remark"`
	Contact        DeliveryPointContact `json:"contact"`
}

// DeliveryPointContact :
type DeliveryPointContact struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

// RequestCreateDelivery :
type RequestCreateDelivery struct {
	DeliveryVendor DeliveryVendor  `json:"deliveryVendor"`
	VehicleType    VehicleType     `json:"vehicleType"`
	Type           DeliveryType    `json:"type"`
	Points         []DeliveryPoint `json:"points"`
}

// ResponseCreateDelivery :
type ResponseCreateDelivery struct {
	Code string   `json:"code"`
	Item Delivery `json:"item"`
	Err  *Error   `json:"error"`
}

// CreateDelivery :
func (c Client) CreateDelivery(request RequestCreateDelivery) (*ResponseCreateDelivery, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathCreateDelivery.method
	requestURL := c.prepareAPIURL(pathCreateDelivery)

	response := new(ResponseCreateDelivery)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}

// ResponseGetDeliveryByID :
type ResponseGetDeliveryByID struct {
	Code string   `json:"code"`
	Item Delivery `json:"item"`
	Err  *Error   `json:"error"`
}

// GetDeliveryByID :
func (c Client) GetDeliveryByID(id string) (*ResponseGetDeliveryByID, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := parhGetDeliveryByID.method
	requestURL := c.prepareAPIURL(parhGetDeliveryByID)
	response := new(ResponseGetDeliveryByID)
	if err := c.httpAPI(method, fmt.Sprintf("%s/%s", requestURL, id), nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// RequestCalculateDeliveryFee :
type RequestCalculateDeliveryFee struct {
	DeliveryVendor DeliveryVendor  `json:"deliveryVendor"`
	VehicleType    VehicleType     `json:"vehicleType"`
	Type           DeliveryType    `json:"type"`
	Points         []DeliveryPoint `json:"points"`
}

// ResponseCalculateDeliveryFee :
type ResponseCalculateDeliveryFee struct {
	Code string `json:"code"`
	Err  *Error `json:"error"`
	Item struct {
		CurrencyType string `json:"currencyType"`
		Amount       uint64 `json:"amount"`
	} `json:"item"`
}

// CalculateDeliveryFee :
func (c Client) CalculateDeliveryFee(request RequestCalculateDeliveryFee) (*ResponseCalculateDeliveryFee, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathCreateDelivery.method
	requestURL := c.prepareAPIURL(pathCreateDelivery)

	response := new(ResponseCalculateDeliveryFee)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}

// RequestConfirmDelivery :
type RequestConfirmDelivery struct {
	DeliveryID string `json:"deliveryId"`
	CheckoutID string `json:"checkoutId"`
}

// ResponseConfirmDelivery :
type ResponseConfirmDelivery struct {
	Code string   `json:"code"`
	Item Delivery `json:"item"`
	Err  *Error   `json:"error"`
}

// ConfirmDelivery :
func (c Client) ConfirmDelivery(request RequestConfirmDelivery) (*ResponseConfirmDelivery, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathConfirmDelivery.method
	requestURL := c.prepareAPIURL(pathConfirmDelivery)

	response := new(ResponseConfirmDelivery)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
