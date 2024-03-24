package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vnkdj5/devtools/pkg/constants/errorcodes"
	"github.com/vnkdj5/devtools/pkg/types"
	"github.com/vnkdj5/devtools/pkg/utils/base64utils"
	"github.com/vnkdj5/devtools/pkg/utils/errorutils"
	"github.com/vnkdj5/devtools/pkg/utils/httputils"
	"go.uber.org/zap"
)

type Base64Handler struct {
	logger *zap.Logger
}

func NewBase64Handler(logger *zap.Logger) *Base64Handler {
	return &Base64Handler{logger: logger}
}

func (h *Base64Handler) Encode(c echo.Context) error {
	request := new(types.InputDataRequest)
	response := new(types.Response)
	if err := httputils.ParseAndValidateRequest(c, request, h.logger); err != nil {
		response.Error = &types.Error{
			Code:        errorutils.GetErrorCode(errorcodes.ValidationErrorCode, errorcodes.RequestParsingErrorCode),
			Description: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	encoded := base64utils.Encode(request.Data)
	response.Data = encoded

	return c.JSON(http.StatusOK, response)

}

func (h *Base64Handler) Decode(c echo.Context) error {
	request := new(types.InputDataRequest)
	response := new(types.Response)
	if err := httputils.ParseAndValidateRequest(c, request, h.logger); err != nil {
		response.Error = &types.Error{
			Code:        errorutils.GetErrorCode(errorcodes.ValidationErrorCode, errorcodes.RequestParsingErrorCode),
			Description: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	encoded, err := base64utils.Decode(request.Data)
	if err != nil {
		response.Error = &types.Error{
			Code:           errorutils.GetErrorCode(errorcodes.ValidationErrorCode, errorcodes.RequestParsingErrorCode),
			Description:    err.Error(),
			AdditionalInfo: nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response.Data = encoded

	return c.JSON(http.StatusOK, response)
}
