package delivery

///handler = controller
import (
	"be13/project/features/user"
	"be13/project/middlewares"
	"be13/project/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserDeliv struct {
	UserService user.ServiceEntities
}

func New(Service user.ServiceEntities, e *echo.Echo) {
	handler := &UserDeliv{
		UserService: Service,
	}

	e.POST("/user", handler.Create)
	e.GET("/user", handler.GetAll, middlewares.JWTMiddleware())
	e.PUT("/user/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/user/:id", handler.DeleteById, middlewares.JWTMiddleware())
	e.GET("/user/:id", handler.GetById, middlewares.JWTMiddleware())

}
func (delivery *UserDeliv) Create(c echo.Context) error {

	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "Admin" {
	// 	return c.JSON(http.StatusUnauthorized, helper.PesanGagalHelper("tidak bisa diakses khusus admin!!!"))
	// }

	Inputuser := UserRequest{} //penangkapan data user reques dari entities user
	errbind := c.Bind(&Inputuser)

	generatePass := user.Bcript(Inputuser.Password)
	Inputuser.Password = generatePass

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errbind.Error()))
	}
	dataCore := UserRequestToUserCore(Inputuser) //data mapping yang diminta create
	errResultCore := delivery.UserService.Create(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("berhasil create user"))
}
func (delivery *UserDeliv) GetAll(c echo.Context) error {
	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "admin" {
	// 	return c.JSON(http.StatusUnauthorized, helper.PesanGagalHelper("tidak bisa diakses khusus admin!!!"))
	// }
	result, err := delivery.UserService.GetAll() //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = ListUserCoreToUserRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca  user", ResponData))

}
func (delivery *UserDeliv) Update(c echo.Context) error {

	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "Admin" {
	// 	return c.JSON(http.StatusUnauthorized, helper.PesanGagalHelper("tidak bisa diakses khusus admin!!!"))
	// }
	id, _ := strconv.Atoi(c.Param("id"))

	userInput := UserRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := UserRequestToUserCore(userInput)
	err := delivery.UserService.Update(id, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success Update data"))
}
func (delivery *UserDeliv) DeleteById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	del, err := delivery.UserService.DeleteById(id) //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr Hapus data"))
	}
	result := UserCoreToUserRespon(del)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil menghapus user", result))
}

func (delivery *UserDeliv) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := delivery.UserService.GetById(id) //memanggil fungsi service yang ada di folder service//jika return nya 2 maka variable harus ada 2

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = UserCoreToUserRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca  user", ResponData))
}
