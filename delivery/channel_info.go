package delivery

import (
	"messy-stuff/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ChannelInfo(c echo.Context) error {
	workers, _ := strconv.Atoi(c.QueryParam("workers"))
	data := usecase.ProcessChannel(workers)
	return c.JSON(http.StatusOK, data)
}
