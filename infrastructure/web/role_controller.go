package web

import (
	"github.com/labstack/echo/v4"
	"go-user-auth-api/application/handler"
	"go-user-auth-api/domain"
	commandQueryHandler "go-user-auth-api/infrastructure/commandquerybus"
	"net/http"
	"strconv"
)

type RoleController struct {
	commandQueryHandler commandQueryHandler.Handler
}

func NewRoleController(commandQueryHandler commandQueryHandler.Handler) *RoleController {
	return &RoleController{
		commandQueryHandler: commandQueryHandler,
	}
}

// CreateRole godoc
// @tags role-controller
// @Summary Create Role
// @Accept  json
// @Produce  json
// @Param request body domain.RoleCreateRequest true "RoleCreateRequest"
// @Success 201
// @Router /roles [post]
func (controller *RoleController) CreateRole(c echo.Context) error {
	var command domain.CreateRoleCommand
	err := c.Bind(&command)
	if err != nil {
		return err
	}

	if err := command.Validate(); err != nil {
		return err
	}

	resp, errors := controller.commandQueryHandler.Handle(command)
	if errors != nil {
		return errors[0]
	} else {
		event := resp.(domain.RoleCreatedEvent)
		c.Response().Header().Set("id", strconv.Itoa(event.Id))
		return c.NoContent(201)
	}
}

// GetRoleById godoc
// @tags role-controller
// @Summary Get Role Info by id
// @Accept  json
// @Produce  json
// @Param id path int true "Role Id"
// @Success 200 {object} domain.RoleDto
// @Router /roles/{id} [get]
func (controller *RoleController) GetRoleById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	role, err := controller.commandQueryHandler.Handle(handler.GetRoleByIdQuery{Id: id})
	//	role, err := controller.service.GetRoleById(id)
	if err != nil {
		return err[0]
	} else {
		return c.JSON(http.StatusOK, role)
	}
}

// DeleteRoleById godoc
// @tags role-controller
// @Summary Delete Role by id
// @Accept  json
// @Produce  json
// @Param id path int true "Role Id"
// @Success 204
// @Router /roles/{id} [delete]
func (controller *RoleController) DeleteRoleById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, errors := controller.commandQueryHandler.Handle(domain.DeleteRoleCommand{Id: id})
	if errors != nil {
		return errors[0]
	} else {
		event := resp.(domain.RoleDeletedEvent)
		c.Response().Header().Set("id", strconv.Itoa(event.Id))
		return c.NoContent(204)
	}
}

// GetAll godoc
// @tags role-controller
// @Summary Get All Roles info
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.RoleDto
// @Router /roles [get]
func (controller *RoleController) GetAll(c echo.Context) error {
	roles, err := controller.commandQueryHandler.Handle(handler.GetAllRolesQuery{})
	if err != nil {
		return err[0]
	} else {
		return c.JSON(http.StatusOK, roles)
	}
}

func (controller *RoleController) Register(e *echo.Echo) {
	e.POST("/roles", controller.CreateRole)
	e.GET("/roles/:id", controller.GetRoleById)
	e.DELETE("/roles/:id", controller.DeleteRoleById)
	e.GET("/roles", controller.GetAll)
}
