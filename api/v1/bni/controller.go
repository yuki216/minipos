package bni

import (
	"github.com/labstack/echo/v4"
	"go-hexagonal-auth/api/common"
	"go-hexagonal-auth/api/v1/bni/request"
	"go-hexagonal-auth/api/v1/bni/response"
	"go-hexagonal-auth/business/bni"
	"go-hexagonal-auth/config"
)

//Controller Get item API controller
type Controller struct {
	service bni.Service
	cfg config.Config
}

//NewController Construct item API controller
func NewController(service bni.Service, cfg config.Config) *Controller {
	return &Controller{
		service,
		cfg,
	}
}


func (controller *Controller) CreateBilling(c echo.Context) error {

	billingRequest := new(request.CreateBilling)

	if err := c.Bind(billingRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	content,err:= controller.service.CreateBilling(billingRequest)
	if err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	res := response.CreateResponse{
		VirtualAccount: content.VirtualAccount,
		TrxID:          content.TrxID,
	}

	return c.JSON(common.NewSuccessResponse(res))

}

func (controller *Controller) UpdateBilling(c echo.Context) error {

	billingRequest := new(request.CreateBilling)

	if err := c.Bind(billingRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	content,err:= controller.service.CreateBilling(billingRequest)
	if err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	res := response.CreateResponse{
		VirtualAccount: content.VirtualAccount,
		TrxID:          content.TrxID,
	}

	return c.JSON(common.NewSuccessResponse(res))


}

func (controller *Controller) InquiryBilling(c echo.Context) error {

	billingRequest := new(request.CreateBilling)

	if err := c.Bind(billingRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	content,err:= controller.service.CreateBilling(billingRequest)
	if err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	res := response.CreateResponse{
		VirtualAccount: content.VirtualAccount,
		TrxID:          content.TrxID,
	}

	return c.JSON(common.NewSuccessResponse(res))


}