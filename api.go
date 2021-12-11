package sdk

// Error :
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`

	Debug       string      `json:"debug"`
	Description interface{} `json:"description"`
}
