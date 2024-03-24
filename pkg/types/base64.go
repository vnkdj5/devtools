package types

type InputDataRequest struct {
	Data string `json:"data" validate:"required"`
}

type BatchInputDataRequest struct {
	Data []string `json:"data" validate:"required"`
}
