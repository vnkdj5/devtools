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
