package delivery

///handler = controller
import (
	"be13/project/features/comment"
	"be13/project/middlewares"
	"be13/project/utils/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type commentDelivery struct {
	commentService comment.ServiceInterface
}

func NewHome(Service comment.ServiceInterface, e *echo.Echo) {
	handler := &commentDelivery{
		commentService: Service,
	}

	e.POST("/comment", handler.Create, middlewares.JWTMiddleware())

}

func (delivery *commentDelivery) Create(c echo.Context) error {
	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "Host" {
	// 	return c.JSON(http.StatusUnauthorized, helper.FailedResponse("Please Complete your Validation"))
	// }

	userIdtoken := middlewares.ExtractTokenUserId(c)
	log.Println("user id", userIdtoken)

	commentReq := CommentRequest{}
	errbind := c.Bind(&commentReq)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errbind.Error()))
	}

	dataCore := UserRequestToUserCore(commentReq)

	dataCore.UserID = uint(userIdtoken)

	errResult := delivery.commentService.CreateComment(dataCore)
	if errResult != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errResult.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Added Comment"))

}
