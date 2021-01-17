package delivery

import (
	"net/http"

	"github.com/labstack/echo"
)

func ChannelInfo(c echo.Context) error {
	successMsg := map[string]string{"status": "success"}
	return c.JSON(http.StatusOK, successMsg)
}
