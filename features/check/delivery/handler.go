package delivery

import (
	"be13/project/features/check"
	"be13/project/utils/helper"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type checkDelivery struct {
	checkServices check.ServiceInterface
}

func NewCheck(service check.ServiceInterface, e *echo.Echo) {
	handler := &checkDelivery{
		checkServices: service,
	}

	e.GET("/checks", handler.Checkroom)

}

func (delivery *checkDelivery) Checkroom(c echo.Context) error {

	room := c.QueryParam("room_id")

	id, errConv := strconv.Atoi(room)
	if errConv != nil {
		return errors.New("id must integer")
	}
	CheckIn := c.QueryParam("check_in")
	CheckOut := c.QueryParam("check_out")

	data, err := delivery.checkServices.Checkroom(id, CheckIn, CheckOut)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed Get available house"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get room's status", map[string]string{
		"availability": data,
	}))

}
