package types

type InputDataRequest struct {
	Data string `json:"data" validate:"required"`
}

type BatchInputDataRequest struct {
	Data []string `json:"data" validate:"required"`
}

type InputOperationRequest struct {
	InputDataRequest
	Operation string `json:"op" validate:"required"`
}

type InputHMACRequest struct {
	Data      any    `json:"data" validate:"required"`
	Algorithm string `json:"algorithm" validate:"required"`
	SecretKey string `json:"secretKey" validate:"required"`
}
