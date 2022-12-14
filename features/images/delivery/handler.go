package delivery

///handler = controller
import (
	"be13/project/features/images"
	"be13/project/utils/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UploadDelivery struct {
	uploadService images.ServiceEntities
}

func NewImage(Service images.ServiceEntities, e *echo.Echo) {
	handler := &UploadDelivery{
		uploadService: Service,
	}

	e.POST("/upload", handler.Create())
	// e.GET("/user", handler.GetAll, middlewares.JWTMiddleware())
	// e.PUT("/user/:id", handler.Update, middlewares.JWTMiddleware())
	// e.DELETE("/user/:id", handler.DeleteById, middlewares.JWTMiddleware())
	// e.GET("/user/:id", handler.GetById, middlewares.JWTMiddleware())

}

func (delivery *UploadDelivery) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UploadRequest
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Something Error In Server"))
		}

		file, fileheader, err := c.Request().FormFile("images")
		if err == nil {
			res, err := helper.UploadStatusImages(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			log.Print(res)
			input.Images = res
		}
		cnv := RequestToCore(input)

		res := delivery.uploadService.UplodaImg(cnv)
		if res != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusCreated, helper.SuccessResponse("success add status"))
	}
}
