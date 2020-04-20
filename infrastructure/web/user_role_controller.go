package web

import (
	"github.com/labstack/echo/v4"
	"go-user-auth-api/application"
	"go-user-auth-api/domain"
	"net/http"
	"strconv"
)

type UserRoleController struct {
	service application.UserRoleService
}

func NewUserRoleController(service application.UserRoleService) *UserRoleController {
	return &UserRoleController{
		service: service,
	}
}

// CreateUserRole godoc
// @tags user-role-controller
// @Summary Create UserRole
// @Accept  json
// @Produce  json
// @Param request body domain.UserRoleCreateRequest true "UserRoleCreateRequest"
// @Success 201
// @Router /user-roles [post]
func (controller *UserRoleController) CreateUserRole(c echo.Context) error {
	var userRoleRequest domain.UserRoleCreateRequest
	err := c.Bind(&userRoleRequest)
	if err != nil {
		return err
	}

	if err := userRoleRequest.Validate(); err != nil {
		return err
	}

	err = controller.service.CreateUserRole(userRoleRequest)
	if err != nil {
		return err
	} else {
		return c.NoContent(201)
	}
}

// DeleteByUserIdAndRoleId godoc
// @tags user-role-controller
// @Summary Delete UserRole by UserId and RoleId
// @Accept  json
// @Produce  json
// @Param userId path int true "User Id"
// @Param roleId path int true "Role Id"
// @Success 204
// @Router /user-roles/{userId}/roles/{roleId} [delete]
func (controller *UserRoleController) DeleteByUserIdAndRoleId(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	err := controller.service.DeleteByUserIdAndRoleId(userId, roleId)
	if err != nil {
		return err
	} else {
		return c.NoContent(204)
	}
}

// GetAll godoc
// @tags user-role-controller
// @Summary Get All UserRoles info
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.UserRoleDto
// @Router /user-roles [get]
func (controller *UserRoleController) GetAll(c echo.Context) error {
	permissions, err := controller.service.GetAll()
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, permissions)
	}
}

// GetUserRoleWithPermissions godoc
// @tags user-role-controller
// @Summary Get User Role with Permissions
// @Accept  json
// @Produce  json
// @Param userId path int true "User Id"
// @Success 200 {object} domain.UserDto
// @Router /user-roles/{userId} [get]
func (controller *UserRoleController) GetUserRoleWithPermissions(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))
	userDto, err := controller.service.GetUserRoleWithPermissionsByUserId(userId)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, userDto)
	}
}

func (controller *UserRoleController) Register(e *echo.Echo) {
	e.POST("/user-roles", controller.CreateUserRole)
	e.GET("/user-roles", controller.GetAll)
	e.DELETE("/user-roles/:userId/roles/:roleId", controller.DeleteByUserIdAndRoleId)
	e.GET("/user-roles/:userId", controller.GetUserRoleWithPermissions)
}
