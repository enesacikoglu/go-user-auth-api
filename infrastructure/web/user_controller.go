package web

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go-user-auth-api/application"
	"go-user-auth-api/domain"
	"go-user-auth-api/infrastructure/errors"
	"net/http"
	"strconv"
)

type UserController struct {
	userService application.UserService
}

func NewUserController(userService application.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// GetUserById godoc
// @tags user-controller
// @Summary Get user info by id
// @Accept  json
// @Produce  json
// @Param id path int true "User Id"
// @Success 200 {object} domain.UserDto
// @Router /users/{id} [get]
func (controller *UserController) GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.userService.GetUserById(id)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

// GetUserByEmail godoc
// @tags user-controller
// @Summary Get user info by email
// @Accept  json
// @Produce  json
// @Param email query string true "User Email"
// @Success 200 {object} domain.UserDto
// @Router /users/email [get]
func (controller *UserController) GetUserByEmail(c echo.Context) error {
	email := c.QueryParam("email")
	if !govalidator.IsEmail(email) {
		return errors.BadRequest("Email must be valid")
	}
	user, err := controller.userService.GetUserByEmail(email)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

// CreateUser godoc
// @tags user-controller
// @Summary Create User
// @Accept  json
// @Produce  json
// @Param request body domain.UserCreateRequest true "UserCreateRequest"
// @Success 201
// @Router /users [post]
func (controller *UserController) CreateUser(c echo.Context) error {
	var userRequest domain.UserCreateRequest
	err := c.Bind(&userRequest)
	if err != nil {
		return err
	}

	if err := userRequest.Validate(); err != nil {
		return err
	}

	err = controller.userService.CreateUser(userRequest)
	if err != nil {
		return err
	} else {
		return c.NoContent(201)
	}
}

func addUserIdToResponseHeader(c *echo.Response, userId int64) {
	c.Header().Set("userId", strconv.FormatInt(userId, 10))
}

// UpdateUserById godoc
// @tags user-controller
// @Summary Update User by id
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Param request body domain.UserUpdateRequest true "UserUpdateRequest"
// @Success 200
// @Router /users/{id} [put]
func (controller *UserController) UpdateUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var userRequest domain.UserUpdateRequest
	err := c.Bind(&userRequest)

	if err := userRequest.Validate(); err != nil {
		return err
	}

	if err != nil {
		return err
	}
	err = controller.userService.UpdateUserById(id, userRequest)
	if err != nil {
		return err
	} else {
		return c.NoContent(200)
	}
}

// DeleteUserById godoc
// @tags user-controller
// @Summary Delete User by id
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Success 204
// @Router /users/{id} [delete]
func (controller *UserController) DeleteUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.userService.DeleteUserById(id)
	if err != nil {
		return err
	} else {
		return c.NoContent(204)
	}
}

// GetUserRolesById godoc
// @tags user-controller
// @Summary Get User Roles by user id
// @Accept  json
// @Produce  json
// @Param id path int true "User Id"
// @Success 200 {array} domain.RoleDto
// @Router /users/{id}/roles [get]
func (controller *UserController) GetUserRolesById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	userRoles, err := controller.userService.GetUserRolesById(id)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, userRoles)
	}
}

// GetUserRolePermissionsByUserIdAndRoleId godoc
// @tags user-controller
// @Summary Get User Role Permission info
// @Accept  json
// @Produce  json
// @Param userId path int true "User Id"
// @Param roleId path int true "Role Id"
// @Success 200 {array} domain.PermissionDto
// @Router /users/{userId}/roles/{roleId}/permissions [get]
func (controller *UserController) GetUserRolePermissionsByUserIdAndRoleId(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	userRolePermissions, err := controller.userService.GetUserRolePermissionsByUserIdAndRoleId(userId, roleId)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, userRolePermissions)
	}
}

// GetAll godoc
// @tags user-controller
// @Summary Get All Users info
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.UserDto
// @Router /users [get]
func (controller *UserController) GetAll(c echo.Context) error {
	users, err := controller.userService.GetAll()
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, users)
	}
}

func (controller *UserController) Register(e *echo.Echo) {
	e.POST("/users", controller.CreateUser)
	e.GET("/users/email", controller.GetUserByEmail)
	e.GET("/users", controller.GetAll)

	e.GET("/users/:id", controller.GetUserById)
	e.PUT("/users/:id", controller.UpdateUserById)
	e.DELETE("/users/:id", controller.DeleteUserById)

	e.GET("/users/:id/roles", controller.GetUserRolesById)
	e.GET("/users/:userId/roles/:roleId/permissions", controller.GetUserRolePermissionsByUserIdAndRoleId)
}
