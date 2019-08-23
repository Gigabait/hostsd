package httpserver

import (
	// other
	"github.com/labstack/echo"
)

func requestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			log.Info().
				Str("From", ec.RealIP()).
				Str("To", ec.Request().Host).
				Str("Method", ec.Request().Method).
				Str("Path", ec.Request().URL.Path).
				Int64("Length", ec.Request().ContentLength).
				Str("UA", ec.Request().UserAgent()).
				Msg("HTTP request")

			_ = next(ec)
			return nil
		}
	}
}
