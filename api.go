package sdk

// Error :
type Error struct {
	Code        string      `json:"code"`
	Message     string      `json:"message"`
	Description interface{} `json:"description"`
}
