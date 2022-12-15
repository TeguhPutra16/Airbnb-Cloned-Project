package delivery

import (
	"be13/project/features/check"
	_delivery "be13/project/features/homestay/delivery"
	"be13/project/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type checkDelivery struct {
	checkServices check.ServiceInterface
}

func NewAuth(service check.ServiceInterface, e *echo.Echo) {
	handler := &checkDelivery{
		checkServices: service,
	}

	e.GET("/checks", handler.Checkroom)

}

func (delivery *checkDelivery) Checkroom(c echo.Context) error {
	checkInput := CheckRequest{}
	errBind := c.Bind(&checkInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataUser, err := delivery.checkServices.GetAllhomestay(checkInput.CheckIn, checkInput.CheckOut)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed login"))
	}

	data := _delivery.ListCoreToRespon(dataUser)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success login", data))

}
