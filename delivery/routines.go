package delivery

import (
	"net/http"
	"strconv"

	"messy-stuff/usecase"

	"github.com/labstack/echo"
)

func Routines(c echo.Context) error {
	workers, _ := strconv.Atoi(c.QueryParam("workers"))
	totalTime := usecase.ProcessRoutines(workers)

	successMsg := map[string]int64{"totalTime": totalTime}
	return c.JSON(http.StatusOK, successMsg)
}
