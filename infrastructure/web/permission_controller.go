package web

import (
	"github.com/labstack/echo/v4"
	"go-user-auth-api/application"
	"go-user-auth-api/domain"
	"net/http"
	"strconv"
)

type PermissionController struct {
	service application.PermissionService
}

func NewPermissionController(service application.PermissionService) *PermissionController {
	return &PermissionController{
		service: service,
	}
}

// CreatePermission godoc
// @tags permission-controller
// @Summary Create Permission
// @Accept  json
// @Produce  json
// @Param request body domain.PermissionCreateRequest true "PermissionCreateRequest"
// @Success 200
// @Router /permissions [post]
func (controller *PermissionController) CreatePermission(c echo.Context) error {
	var permissionRequest domain.PermissionCreateRequest
	err := c.Bind(&permissionRequest)
	if err != nil {
		return err
	}

	if err := permissionRequest.Validate(); err != nil {
		return err
	}

	err = controller.service.CreatePermission(permissionRequest)
	if err != nil {
		return err
	} else {
		return c.NoContent(201)
	}
}

// GetPermissionById godoc
// @tags permission-controller
// @Summary Get Permission Info by id
// @Accept  json
// @Produce  json
// @Param id path int true "Permission Id"
// @Success 200 {object} domain.PermissionDto
// @Router /permissions/{id} [get]
func (controller *PermissionController) GetPermissionById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	permission, err := controller.service.GetPermissionById(id)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, permission)
	}
}

// DeletePermissionById godoc
// @tags permission-controller
// @Summary Delete Permission by id
// @Accept  json
// @Produce  json
// @Param id path int true "Permission Id"
// @Success 204
// @Router /permissions/{id} [delete]
func (controller *PermissionController) DeletePermissionById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.service.DeletePermissionById(id)
	if err != nil {
		return err
	} else {
		return c.NoContent(204)
	}
}

// GetAll godoc
// @tags permission-controller
// @Summary Get All Permissions info
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.PermissionDto
// @Router /permissions [get]
func (controller *PermissionController) GetAll(c echo.Context) error {
	permissions, err := controller.service.GetAll()
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, permissions)
	}
}
func (controller *PermissionController) Register(e *echo.Echo) {
	e.POST("/permissions", controller.CreatePermission)
	e.GET("/permissions/:id", controller.GetPermissionById)
	e.DELETE("/permissions/:id", controller.DeletePermissionById)
	e.GET("/permissions", controller.GetAll)
}
