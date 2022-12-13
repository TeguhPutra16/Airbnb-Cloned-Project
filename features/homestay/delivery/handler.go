package delivery

///handler = controller
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

	dataCore.UserID = uint(userIdtoken)
	dataCore.Status = "Available"

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
