package web

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go-user-auth-api/application"
	"go-user-auth-api/infrastructure/errors"
	"net/http"
	"strconv"
)

type UserPermissionController struct {
	service application.UserPermissionService
}

func NewUserPermissionController(service application.UserPermissionService) *UserPermissionController {
	return &UserPermissionController{
		service: service,
	}
}

// GetUserPermissionsByEmailAndAppId godoc
// @tags user-permission-controller
// @Summary Get UserPermission info by email and applicationId
// @Accept  json
// @Produce  json
// @Param email query string true "User email"
// @Param appId query int true "Application Id"
// @Success 200 {array} domain.PermissionDto
// @Router /user-permissions [get]
func (controller *UserPermissionController) GetUserPermissionsByEmailAndAppId(c echo.Context) error {
	email := c.QueryParam("email")
	if !govalidator.IsEmail(email) {
		return errors.BadRequest("Email must be valid")
	}
	appId ,err:= strconv.Atoi(c.QueryParam("appId"))
	if err!=nil {
		return errors.BadRequest("Invalid ApplicationId")
	}

	user, err := controller.service.GetUserPermissionsByEmailAndAppId(email, appId)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

func (controller *UserPermissionController) Register(e *echo.Echo) {
	e.GET("/user-permissions", controller.GetUserPermissionsByEmailAndAppId)
}
