package types

// Response type export
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error *Error      `json:"error,omitempty"`
}

// Error type export
type Error struct {
	Code           string      `json:"code,omitempty"`
	Description    string      `json:"description,omitempty"`
	AdditionalInfo interface{} `json:"additionalInfo,omitempty"`
}

// Request type export
type Request struct {
	URI         string      `json:"uri"`
	Method      string      `json:"method"`
	QueryString string      `json:"queryString"`
	Body        interface{} `json:"body,omitempty"`
}
