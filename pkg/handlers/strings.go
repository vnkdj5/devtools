package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/vnkdj5/devtools/pkg/constants/errorcodes"
	"github.com/vnkdj5/devtools/pkg/types"
	"github.com/vnkdj5/devtools/pkg/utils/errorutils"
	"github.com/vnkdj5/devtools/pkg/utils/httputils"
	"github.com/vnkdj5/devtools/pkg/utils/stringutils"
	"go.uber.org/zap"
)

type StringsHandler struct {
	logger *zap.Logger
}

func NewStringsHandler(logger *zap.Logger) *StringsHandler {
	return &StringsHandler{logger: logger}
}

func (h *StringsHandler) Process(c echo.Context) error {
	request := new(types.InputOperationRequest)
	response := new(types.Response)
	if err := httputils.ParseAndValidateRequest(c, request, h.logger); err != nil {
		response.Error = &types.Error{
			Code:        errorutils.GetErrorCode(errorcodes.ValidationErrorCode, errorcodes.RequestParsingErrorCode),
			Description: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	var convertedText string
	switch strings.ToLower(request.Operation) {
	case "uppercase":
		{
			convertedText = strings.ToUpper(request.Data)
			break
		}
	case "lowercase":
		{
			convertedText = strings.ToLower(request.Data)
			break
		}
	case "titlecase":
		{
			convertedText = strings.ToTitle(request.Data)
			break
		}
	case "length":
		fallthrough
	case "size":
		{
			convertedText = fmt.Sprintf("%d", len(request.Data))
			break
		}
	case "reverse":
		convertedText = stringutils.Reverse(request.Data)
	default:
		{
			response.Error = &types.Error{
				Code:        errorutils.GetErrorCode(errorcodes.ValidationErrorCode, errorcodes.RequestParsingErrorCode),
				Description: fmt.Sprintf("Operation '%s'  is not supported", request.Operation),
			}
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	response.Data = convertedText

	return c.JSON(http.StatusOK, response)

}
