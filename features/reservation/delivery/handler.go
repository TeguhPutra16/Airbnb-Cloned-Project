package delivery

///handler = controller
import (
	"be13/project/features/reservation"
	"be13/project/middlewares"
	"be13/project/utils/helper"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type reservasiDeliv struct {
	reservasiService reservation.ServiceEntities
}

func NewRes(Service reservation.ServiceEntities, e *echo.Echo) {
	handler := &reservasiDeliv{
		reservasiService: Service,
	}

	e.POST("/reserves", handler.Create, middlewares.JWTMiddleware())
	e.GET("/reserves/user", handler.GetUserReservation, middlewares.JWTMiddleware())
	// e.GET("/users", handler.GetAll, middlewares.JWTMiddleware())
	// e.PUT("/users/:id", handler.Update, middlewares.JWTMiddleware())
	// e.DELETE("/users/:id", handler.DeleteById, middlewares.JWTMiddleware())
	// e.GET("/users/:id", handler.GetById, middlewares.JWTMiddleware())

}

func (delivery *reservasiDeliv) Create(c echo.Context) error {

	userIdtoken := middlewares.ExtractTokenUserId(c)

	Input := ReservasiRequest{} //penangkapan data user reques dari entities user
	errbind := c.Bind(&Input)

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data0"+errbind.Error()))
	}
	log.Println(Input.CheckIn)
	res, errConvtime1 := time.Parse("2006-01-02", Input.CheckIn)
	if errConvtime1 != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data1"))
	}
	res1, errConvtime2 := time.Parse("2006-01-02", Input.CheckOut)
	if errConvtime2 != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data2"))
	}

	log.Println(Input.UserID)

	dataCore := RequestToCore(Input, res, res1) //data mapping yang diminta create\
	dataCore.UserID = uint(userIdtoken)
	errResultCore := delivery.reservasiService.Create(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data3"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success Create Reservation "))
}

func (delivery *reservasiDeliv) GetUserReservation(c echo.Context) error {
	userIdtoken := middlewares.ExtractTokenUserId(c)
	log.Println("user_id_token", userIdtoken)

	result, err := delivery.reservasiService.GetAllreservasi(userIdtoken) //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = ListCoreToRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca homestay dari user", ResponData))

}
