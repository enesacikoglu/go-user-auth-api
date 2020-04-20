package domain

import (
	"time"
)

type User struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	CreatedBy    string    `json:"createdBy"`
	ModifiedBy   string    `json:"modifiedBy"`
	CreatedDate  time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
}

type Role struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	CreatedBy    string    `json:"createdBy"`
	ModifiedBy   string    `json:"modifiedBy"`
	CreatedDate  time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
}

type UserRole struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	RoleId      int       `json:"roleId"`
	CreatedBy   string    `json:"createdBy"`
	CreatedDate time.Time `json:"createdDate"`
}

type Permission struct {
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	CreatedBy      string         `json:"createdBy"`
	ModifiedBy     string         `json:"modifiedBy"`
	CreatedDate    time.Time      `json:"createdDate"`
	ModifiedDate   time.Time      `json:"modifiedDate"`
	ApplicationId  int            `json:"applicationId,omitempty"`
}

type RolePermission struct {
	Id            int       `json:"id"`
	RoleId        int       `json:"roleId"`
	PermissionId  int       `json:"permissionId"`
	CreatedBy     string    `json:"createdBy"`
	CreatedDate   time.Time `json:"createdDate"`
	ApplicationId int       `json:"applicationId"`
}

type Application struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	CreatedBy    string    `json:"createdBy"`
	ModifiedBy   string    `json:"modifiedBy"`
	CreatedDate  time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
}

type UserRolePermission struct {
	UserId                 int            `json:"userId"`
	Email                  string         `json:"email"`
	UserName               string         `json:"userName"`
	Surname                string         `json:"surname"`
	UserCreatedBy          string         `json:"userCreatedBy"`
	UserModifiedBy         string         `json:"userModifiedBy"`
	UserCreatedDate        time.Time      `json:"userCreatedDate"`
	UserModifiedDate       time.Time      `json:"userModifiedDate"`
	RoleId                 int            `json:"roleId"`
	RoleName               string         `json:"roleName"`
	RoleCreatedBy          string         `json:"roleCreatedBy"`
	RoleModifiedBy         string         `json:"roleModifiedBy "`
	RoleCreatedDate        time.Time      `json:"roleCreatedDate"`
	RoleModifiedDate       time.Time      `json:"roleModifiedDate"`
	PermissionId           int            `json:"permissionId"`
	PermissionName         string         `json:"permissionName"`
	PermissionCreatedBy    string         `json:"permissionCreatedBy"`
	PermissionModifiedBy   string         `json:"permissionModifiedBy"`
	PermissionCreatedDate  time.Time      `json:"permissionCreatedDate"`
	PermissionModifiedDate time.Time      `json:"permissionModifiedDate"`
	ApplicationId          int            `json:"applicationId"`
}

type UserRolePermissions []UserRolePermission
