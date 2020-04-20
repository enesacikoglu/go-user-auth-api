package web

import (
	"github.com/labstack/echo/v4"
	"go-user-auth-api/application"
	"go-user-auth-api/domain"
	"net/http"
	"strconv"
)

type ApplicationController struct {
	service application.ApplicationsService
}

func NewApplicationController(service application.ApplicationsService) *ApplicationController {
	return &ApplicationController{
		service: service,
	}
}

// CreateApplication godoc
// @tags application-controller
// @Summary Create Application
// @Accept  json
// @Produce  json
// @Param request body domain.ApplicationCreateRequest true "ApplicationCreateRequest"
// @Success 201
// @Router /applications [post]
func (controller *ApplicationController) CreateApplication(c echo.Context) error {
	var applicationRequest domain.ApplicationCreateRequest
	err := c.Bind(&applicationRequest)
	if err != nil {
		return err
	}

	if err := applicationRequest.Validate(); err != nil {
		return err
	}

	err = controller.service.CreateApplication(applicationRequest)
	if err != nil {
		return err
	} else {
		return c.NoContent(201)
	}
}

// GetApplicationById godoc
// @tags application-controller
// @Summary Get Application Info by id
// @Accept  json
// @Produce  json
// @Param id path int true "Application Id"
// @Success 200 {object} domain.ApplicationDto
// @Router /applications/{id} [get]
func (controller *ApplicationController) GetApplicationById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	application, err := controller.service.GetApplicationById(id)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, application)
	}
}


// GetAll godoc
// @tags application-controller
// @Summary Get All Applications info
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.ApplicationDto
// @Router /applications [get]
func (controller *ApplicationController) GetAll(c echo.Context) error {
	applications, err := controller.service.GetAll()
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, applications)
	}
}

func (controller *ApplicationController) Register(e *echo.Echo) {
	e.POST("/applications", controller.CreateApplication)
	e.GET("/applications/:id", controller.GetApplicationById)
	e.GET("/applications", controller.GetAll)
}
