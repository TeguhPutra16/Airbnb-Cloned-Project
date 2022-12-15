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

	e.POST("/homestays", handler.Create, middlewares.JWTMiddleware())
	e.GET("/homestays", handler.GetAllhomestay, middlewares.JWTMiddleware())
	e.PUT("/homestays/:id", handler.Update, middlewares.JWTMiddleware())
	e.GET("/homestays/:id", handler.GetById, middlewares.JWTMiddleware())
	e.GET("/homestays/user", handler.GetUserHomestay, middlewares.JWTMiddleware())
	e.DELETE("/homestays/:id", handler.DeleteById, middlewares.JWTMiddleware())

}

func (delivery *homeStayDelivery) Create(c echo.Context) error {

	userIdtoken := middlewares.ExtractTokenUserId(c)
	log.Println("user id", userIdtoken)

	homestayReq := HomestayRequest{}
	errbind := c.Bind(&homestayReq)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errbind.Error()))
	}
	file, fileheader, err := c.Request().FormFile("images")

	if err == nil {
		res, err := helper.Uploader.UploadFile(file, fileheader.Filename)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Name file already use"))
		}
		// log.Print(res)
		homestayReq.Images = res
	}
	dataCore := UserRequestToUserCore(homestayReq)

	dataCore.UserID = uint(userIdtoken)

	errResult := delivery.homeStayService.Create(dataCore)
	if errResult != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errResult.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Added Home Stay"))

}

func (delivery *homeStayDelivery) DeleteById(c echo.Context) error {

	userIdtoken := middlewares.ExtractTokenUserId(c)
	log.Println("user_id_token", userIdtoken)
	id, _ := strconv.Atoi(c.Param("id"))
	del, err := delivery.homeStayService.DeleteById(id) //memanggil fungsi service yang ada di folder service
	log.Println("user_id_comment", del.UserID)
	if del.UserID != uint(userIdtoken) {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("tidak bisa hapus selain komen sendri"))

	}
	//memanggil fungsi service yang ada di folder service

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
	userIdtoken := middlewares.ExtractTokenUserId(c)
	log.Println("user_id_token", userIdtoken)

	if dataCore.UserID != uint(userIdtoken) {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("hanya bisa update data sendiri "))

	}
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

func (delivery *homeStayDelivery) GetUserHomestay(c echo.Context) error {
	userIdtoken := middlewares.ExtractTokenUserId(c)
	log.Println("user_id_token", userIdtoken)

	result, err := delivery.homeStayService.GethHomestaybyidUser(userIdtoken) //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = ListCoreToRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca homestay dari user", ResponData))

}
