package delivery

///handler = controller
import (
	"be13/project/features/comment"
	"be13/project/middlewares"
	"be13/project/utils/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type commentDelivery struct {
	commentService comment.ServiceInterface
}

func NewComment(Service comment.ServiceInterface, e *echo.Echo) {
	handler := &commentDelivery{
		commentService: Service,
	}

	e.POST("/comments", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/comments/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/comments/:id", handler.DeleteById, middlewares.JWTMiddleware())

}

func (delivery *commentDelivery) Create(c echo.Context) error {
	userIdtoken := middlewares.ExtractTokenUserId(c)
	log.Println("user id", userIdtoken)

	commentReq := CommentRequest{}

	errbind := c.Bind(&commentReq)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errbind.Error()))
	}
	if commentReq.Ratings < 0 || commentReq.Ratings > 5 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("rates in range 0 to 5"))
	}

	dataCore := RequestToCore(commentReq)

	dataCore.UserID = uint(userIdtoken)

	errResult := delivery.commentService.CreateComment(dataCore)
	if errResult != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errResult.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Added Comment"))

}

func (delivery *commentDelivery) DeleteById(c echo.Context) error {
	/////////////hanya bisa hapus komen sendiri////////////////////////////////////////////
	userIdtoken := middlewares.ExtractTokenUserId(c)
	log.Println("user_id_token", userIdtoken)
	id, _ := strconv.Atoi(c.Param("id"))
	del, err := delivery.commentService.DeleteById(id, userIdtoken) //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr Hapus data"))
	}
	result := CoreToRespon(del)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil menghapus comment", result))
}

func (delivery *commentDelivery) Update(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	komenInput := CommentRequest{}
	errBind := c.Bind(&komenInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := RequestToCore(komenInput)
	userIdtoken := middlewares.ExtractTokenUserId(c)

	err := delivery.commentService.UpdateComment(id, userIdtoken, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update komentar"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success Update Comment"))
}
