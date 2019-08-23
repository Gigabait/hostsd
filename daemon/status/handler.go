package status

import (
	// stdlib
	"net/http"

	// other
	"github.com/labstack/echo"
)

func handleStatusGET(ec echo.Context) error {
	// Stub for now.
	return ec.JSON(http.StatusOK, map[string]string{"status": "working"})
}
