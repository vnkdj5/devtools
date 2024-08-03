package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vnkdj5/devtools/pkg/constants/errorcodes"
	"github.com/vnkdj5/devtools/pkg/types"
	"github.com/vnkdj5/devtools/pkg/utils/errorutils"
	"github.com/vnkdj5/devtools/pkg/utils/hmac"
	"github.com/vnkdj5/devtools/pkg/utils/httputils"
	"github.com/vnkdj5/devtools/pkg/utils/stringutils"
	"go.uber.org/zap"
)

type CryptoHandler struct {
	logger *zap.Logger
}

func NewCryptoHandler(logger *zap.Logger) *CryptoHandler {
	return &CryptoHandler{logger: logger}
}

func (h *CryptoHandler) GenerateHMAC(c echo.Context) error {
	request := new(types.InputHMACRequest)
	response := new(types.Response)
	if err := httputils.ParseAndValidateRequest(c, request, h.logger); err != nil {
		response.Error = &types.Error{
			Code:        errorutils.GetErrorCode(errorcodes.ValidationErrorCode, errorcodes.RequestParsingErrorCode),
			Description: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	// Convert data to string for HMAC generation
	dataString, err := stringutils.ConvertDataToString(request.Data)
	if err != nil {
		response.Error = &types.Error{
			Code:        errorutils.GetErrorCode(errorcodes.ValidationErrorCode, errorcodes.RequestParsingErrorCode),
			Description: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	hmac, err := hmac.GenerateHMAC(dataString, request.Algorithm, request.SecretKey)
	if err != nil {
		response.Error = &types.Error{
			Code:        errorutils.GetErrorCode(errorcodes.HMACGenerationError, errorcodes.HMACGenerationError),
			Description: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Data = hmac

	return c.JSON(http.StatusOK, response)

}
