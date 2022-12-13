package delivery

import (
	"be13/project/features/homestay"
	"be13/project/middlewares"
	"be13/project/utils/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type homeStayDelivery struct {
	homeStayService homestay.ServiceEntities
}

func NewHome(Service homestay.ServiceEntities, e *echo.Echo) {
	handler := &homeStayDelivery{
		homeStayService: Service,
	}

	e.POST("/homestay", handler.Create, middlewares.JWTMiddleware())
	e.GET("/homestays", handler.GetAllhomestay, middlewares.JWTMiddleware())
	e.PUT("/homestay/:id", handler.Update, middlewares.JWTMiddleware())
	e.GET("/homestay/:id", handler.GetById, middlewares.JWTMiddleware())
	e.DELETE("/homestay/:id", handler.DeleteById, middlewares.JWTMiddleware())
	e.GET("/homestay", handler.GetBytime, middlewares.JWTMiddleware())

}

func (delivery *homeStayDelivery) Create(c echo.Context) error {
	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "Host" {
	// 	return c.JSON(http.StatusUnauthorized, helper.FailedResponse("Please Complete your Validation"))
	// }

	userIdtoken := middlewares.ExtractTokenUserId(c)
	log.Println("user id", userIdtoken)

	homestayReq := HomestayRequest{}
	errbind := c.Bind(&homestayReq)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errbind.Error()))
	}

	dataCore := UserRequestToUserCore(homestayReq)

	dataCore.Status = "Available"
	dataCore.UserID = uint(userIdtoken)

	errResult := delivery.homeStayService.Create(dataCore)
	if errResult != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errResult.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Added Home Stay"))

}

func (delivery *homeStayDelivery) GetBytime(c echo.Context) error {
	start := c.QueryParam("start")

	end := c.QueryParam("end")

	data, err := delivery.homeStayService.GetBytime(start, end)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = ListCoreToRespon(data)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca  user", ResponData))
}

func (delivery *homeStayDelivery) DeleteById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	del, err := delivery.homeStayService.DeleteById(id) //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr Hapus data"))
	}
	result := CoreToRespon(del)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil menghapus user", result))
}

func (delivery *homeStayDelivery) Update(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	homeInput := HomestayRequest{}
	errBind := c.Bind(&homeInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := UserRequestToUserCore(homeInput)
	err := delivery.homeStayService.Update(id, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success Update data"))
}

func (delivery *homeStayDelivery) GetAllhomestay(c echo.Context) error {

	result, err := delivery.homeStayService.GetAllhomestay() //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = ListCoreToRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca homestay", ResponData))

}

func (delivery *homeStayDelivery) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := delivery.homeStayService.GetById(id) //memanggil fungsi service yang ada di folder service//jika return nya 2 maka variable harus ada 2

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca ruangan dan commentnya", result))
}
