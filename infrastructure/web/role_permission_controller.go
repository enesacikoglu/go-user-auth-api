package web

import (
	"github.com/labstack/echo/v4"
	"go-user-auth-api/application"
	"go-user-auth-api/domain"
	"net/http"
	"strconv"
)

type RolePermissionController struct {
	service application.RolePermissionService
}

func NewRolePermissionController(service application.RolePermissionService) *RolePermissionController {
	return &RolePermissionController{
		service: service,
	}
}

// CreateRolePermission godoc
// @tags role-permission-controller
// @Summary Create Role Permission
// @Accept  json
// @Produce  json
// @Param request body domain.RolePermissionCreateRequest true "RolePermissionCreateRequest"
// @Success 201
// @Router /role-permissions [post]
func (controller *RolePermissionController) CreateRolePermission(c echo.Context) error {
	var rolePermissionRequest domain.RolePermissionCreateRequest
	err := c.Bind(&rolePermissionRequest)
	if err != nil {
		return err
	}

	if err := rolePermissionRequest.Validate(); err != nil {
		return err
	}

	err = controller.service.CreateRolePermission(rolePermissionRequest)
	if err != nil {
		return err
	} else {
		return c.NoContent(201)
	}
}

// DeleteByRoleIdAndPermissionIdAndApplicationId godoc
// @tags role-permission-controller
// @Summary Delete RolePermission by RoleId,PermissionId,ApplicationId
// @Accept  json
// @Produce  json
// @Param roleId path int true "Role Id"
// @Param permissionId path int true "Application Id"
// @Param applicationId path int true "Application Id"
// @Success 204
// @Router /role-permissions/{roleId}/permissions/{permissionId}/applications/{applicationId} [delete]
func (controller *RolePermissionController) DeleteByRoleIdAndPermissionIdAndApplicationId(c echo.Context) error {
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	permissionId, _ := strconv.Atoi(c.Param("permissionId"))
	applicationId, _ := strconv.Atoi(c.Param("applicationId"))
	err := controller.service.DeleteByRoleIdAndPermissionIdAndApplicationId(roleId, permissionId, applicationId)
	if err != nil {
		return err
	} else {
		return c.NoContent(204)
	}
}

// GetAll godoc
// @tags role-permission-controller
// @Summary Get All RolePermissions info
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.RolePermissionDto
// @Router /role-permissions [get]
func (controller *RolePermissionController) GetAll(c echo.Context) error {
	permissions, err := controller.service.GetAll()
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, permissions)
	}
}

func (controller *RolePermissionController) Register(e *echo.Echo) {
	e.POST("/role-permissions", controller.CreateRolePermission)
	e.GET("/role-permissions", controller.GetAll)
	e.DELETE("/role-permissions/:roleId/permissions/:permissionId/applications/:applicationId", controller.DeleteByRoleIdAndPermissionIdAndApplicationId)
}
