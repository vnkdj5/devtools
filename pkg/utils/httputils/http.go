package httputils

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func ParseAndValidateRequest(c echo.Context, request interface{}, logger *zap.Logger) error {
	if err := c.Bind(request); err != nil {
		logger.Debug("Error occurred while parsing request ", zap.Error(err), zap.Any("requestURI", c.Request().RequestURI))
		return fmt.Errorf("Error occurred while parsing request : %w", err)
	}

	binder := &echo.DefaultBinder{}
	if headerBindErr := binder.BindHeaders(c, request); headerBindErr != nil {
		logger.Debug("Error occurred while parsing headers ", zap.Error(headerBindErr), zap.Any("requestURI", c.Request().RequestURI))
		return fmt.Errorf("Error occurred while parsing headers : %w", headerBindErr)
	}

	if validateErr := c.Validate(request); validateErr != nil {
		logger.Debug("Validation failed", zap.Error(validateErr), zap.Any("requestURI", c.Request().RequestURI))
		return fmt.Errorf("Validation failed. Errors: %w", validateErr)
	}
	return nil
}
