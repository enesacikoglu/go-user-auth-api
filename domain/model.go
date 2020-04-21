package domain

import (
	"github.com/asaskevich/govalidator"
	"go-user-auth-api/infrastructure/errors"
	"strconv"
	"time"
)

// Response Models
type UserDto struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	CreatedBy    string    `json:"createdBy"`
	ModifiedBy   string    `json:"modifiedBy"`
	CreatedDate  time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
	Roles        []RoleDto `json:"roles,omitempty"`
}

type UserRoleDto struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	RoleId      int       `json:"roleId"`
	CreatedBy   string    `json:"createdBy"`
	CreatedDate time.Time `json:"createdDate"`
}

type RoleDto struct {
	Id           int             `json:"id"`
	Name         string          `json:"name"`
	CreatedBy    string          `json:"createdBy"`
	ModifiedBy   string          `json:"modifiedBy"`
	CreatedDate  time.Time       `json:"createdDate"`
	ModifiedDate time.Time       `json:"modifiedDate"`
	Permissions  []PermissionDto `json:"permissions,omitempty"`
}

type PermissionDto struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	CreatedBy     string    `json:"createdBy"`
	ModifiedBy    string    `json:"modifiedBy"`
	CreatedDate   time.Time `json:"createdDate"`
	ModifiedDate  time.Time `json:"modifiedDate"`
	ApplicationId int       `json:"applicationId,omitempty"`
}

type ApplicationDto struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	CreatedBy    string    `json:"createdBy"`
	ModifiedBy   string    `json:"modifiedBy"`
	CreatedDate  time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
}

type RolePermissionDto struct {
	Id            int       `json:"id"`
	RoleId        int       `json:"roleId"`
	PermissionId  int       `json:"permissionId"`
	CreatedBy     string    `json:"createdBy"`
	CreatedDate   time.Time `json:"createdDate"`
	ApplicationId int       `json:"applicationId"`
}

// Requests Models
type UserCreateRequest struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	CreatedBy  string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
}

type UserUpdateRequest struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	ModifiedBy string `json:"modifiedBy"`
}

type UserRoleCreateRequest struct {
	UserId    int    `json:"userId"`
	RoleId    int    `json:"roleId"`
	CreatedBy string `json:"createdBy"`
}

type CreateRoleCommand struct {
	Name       string `json:"name"`
	CreatedBy  string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
}

type RoleCreatedEvent struct {
	Id    int `json:"id"`
}

type DeleteRoleCommand struct {
	Id    int `json:"id"`
}

type RoleDeletedEvent RoleCreatedEvent

type ApplicationCreateRequest struct {
	Name       string `json:"name"`
	CreatedBy  string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
}

type ApplicationUpdateRequest struct {
	Name       string `json:"name"`
	ModifiedBy string `json:"modifiedBy"`
}

type RolePermissionCreateRequest struct {
	RoleId        int    `json:"roleId,omitempty" valid:"required"`
	PermissionId  int    `json:"permissionId,omitempty" valid:"required"`
	CreatedBy     string `json:"createdBy"`
	ApplicationId int    `json:"applicationId"`
}

type RoleUpdateRequest CreateRoleCommand

type PermissionCreateRequest struct {
	Name       string `json:"name" valid:"required~Name can not be blank"`
	CreatedBy  string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
}

type PermissionUpdateRequest PermissionCreateRequest

func (u UserCreateRequest) Validate() error {
	if !govalidator.IsNotNull(u.Name) {
		return errors.BadRequest("Name cannot be empty")
	}

	if !govalidator.IsNotNull(u.Surname) {
		return errors.BadRequest("Surname cannot be empty")
	}

	if !govalidator.IsEmail(u.Email) {
		return errors.BadRequest("Email field is not valid")
	}

	if !govalidator.IsNotNull(u.CreatedBy) {
		return errors.BadRequest("CreatedBy cannot be empty")
	}
	if !govalidator.IsNotNull(u.ModifiedBy) {
		return errors.BadRequest("ModifiedBy cannot be empty")
	}
	return nil
}

func (u UserUpdateRequest) Validate() error {
	if !govalidator.IsNotNull(u.Name) {
		return errors.BadRequest("Name cannot be empty")
	}

	if !govalidator.IsNotNull(u.Surname) {
		return errors.BadRequest("Surname cannot be empty")
	}

	if !govalidator.IsEmail(u.Email) {
		return errors.BadRequest("Email field is not valid")
	}

	if !govalidator.IsNotNull(u.ModifiedBy) {
		return errors.BadRequest("ModifiedBy cannot be empty")
	}
	return nil
}

func (a ApplicationCreateRequest) Validate() error {
	if !govalidator.IsNotNull(a.Name) {
		return errors.BadRequest("Name cannot be empty")
	}

	if !govalidator.IsNotNull(a.CreatedBy) {
		return errors.BadRequest("CreatedBy cannot be empty")
	}
	if !govalidator.IsNotNull(a.ModifiedBy) {
		return errors.BadRequest("ModifiedBy cannot be empty")
	}
	return nil
}

func (a ApplicationUpdateRequest) Validate() error {
	if !govalidator.IsNotNull(a.Name) {
		return errors.BadRequest("Name cannot be empty")
	}

	if !govalidator.IsNotNull(a.ModifiedBy) {
		return errors.BadRequest("ModifiedBy cannot be empty")
	}
	return nil
}

func (r CreateRoleCommand) Validate() error {
	if !govalidator.IsNotNull(r.Name) {
		return errors.BadRequest("Name cannot be empty")
	}

	if !govalidator.IsNotNull(r.CreatedBy) {
		return errors.BadRequest("CreatedBy cannot be empty")
	}
	if !govalidator.IsNotNull(r.ModifiedBy) {
		return errors.BadRequest("ModifiedBy cannot be empty")
	}
	return nil
}

func (p PermissionCreateRequest) Validate() error {
	if !govalidator.IsNotNull(p.Name) {
		return errors.BadRequest("Name cannot be empty")
	}

	if !govalidator.IsNotNull(p.CreatedBy) {
		return errors.BadRequest("CreatedBy cannot be empty")
	}

	if !govalidator.IsNotNull(p.ModifiedBy) {
		return errors.BadRequest("ModifiedBy cannot be empty")
	}
	return nil
}

func (ur UserRoleCreateRequest) Validate() error {
	if ur.UserId == 0 {
		return errors.BadRequest("UserId cannot be empty")
	}

	if !govalidator.IsNumeric(strconv.Itoa(ur.UserId)) {
		return errors.BadRequest("UserId must be numeric")
	}

	if ur.RoleId == 0 {
		return errors.BadRequest("RoleId cannot be empty")
	}

	if !govalidator.IsNumeric(strconv.Itoa(ur.RoleId)) {
		return errors.BadRequest("RoleId must be numeric")
	}

	if !govalidator.IsNotNull(ur.CreatedBy) {
		return errors.BadRequest("CreatedBy cannot be empty")
	}
	return nil
}

func (rp RolePermissionCreateRequest) Validate() error {
	if rp.RoleId == 0 {
		return errors.BadRequest("RoleId cannot be empty")
	}

	if !govalidator.IsNumeric(strconv.Itoa(rp.RoleId)) {
		return errors.BadRequest("RoleId must be numeric")
	}

	if rp.PermissionId == 0 {
		return errors.BadRequest("PermissionId cannot be empty")
	}

	if !govalidator.IsNumeric(strconv.Itoa(rp.PermissionId)) {
		return errors.BadRequest("PermissionId must be numeric")
	}

	if !govalidator.IsNotNull(rp.CreatedBy) {
		return errors.BadRequest("CreatedBy cannot be empty")
	}

	if rp.ApplicationId == 0 {
		return errors.BadRequest("ApplicationId cannot be empty")
	}

	if !govalidator.IsNumeric(strconv.Itoa(rp.ApplicationId)) {
		return errors.BadRequest("ApplicationId must be numeric")
	}
	return nil
}
