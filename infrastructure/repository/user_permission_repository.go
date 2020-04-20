package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	"go-user-auth-api/domain"
)

type UserPermissionRepositoryImp struct {
	mssqlConnection MSSqlConnection
}

func NewUserPermissionRepository(mssqlConnection MSSqlConnection) *UserPermissionRepositoryImp {
	return &UserPermissionRepositoryImp{
		mssqlConnection: mssqlConnection,
	}
}

func (repository *UserPermissionRepositoryImp) GetUserPermissionByEmailAndApplicationId(email string, applicationId int) ([]domain.Permission, error) {

	findUserPermissionsByEmailAndApplicationIdQuery := `
    select 
       pr.Id             as PermissionId,
       pr.Name           as PermissionName,
       pr.CreatedBy      as PermissionCreatedBy,
       pr.ModifiedBy     as PermissionModifiedBy,
       pr.CreatedDate    as PermissionCreatedDate,
       pr.ModifiedDate   as PermissionModifiedDate,
       rp.ApplicationId  as ApplicationId
from Users us
       LEFT JOIN UserRoles ur on us.Id = ur.UserId
       JOIN Roles rl ON ur.RoleId = rl.Id
       JOIN RolePermissions rp on ur.RoleId = rp.RoleId
       JOIN Permissions pr ON rp.PermissionId = pr.Id
WHERE us.Email = @email and rp.ApplicationId = @applicationId
	`


	connCtx, cancel := context.WithTimeout(context.Background(), defaultQueryTimeout)
	defer cancel()

	conn, err := repository.mssqlConnection.db.Conn(connCtx)
	if err != nil {
		return nil, fmt.Errorf("%w couldn't open connection from pool", err)
	}

	defer conn.Close()

	queryCtx, cancel := context.WithTimeout(context.Background(), defaultQueryTimeout)
	defer cancel()

	rows, err := conn.QueryContext(queryCtx, findUserPermissionsByEmailAndApplicationIdQuery,
		sql.Named("email", email),
		sql.Named("applicationId", applicationId))
	if err != nil {
		return nil, fmt.Errorf("%w  connection executure query", err)
	}

	permissions := make([]domain.Permission, 0)
	for rows.Next() {
		var permission domain.Permission
		if err := rows.Scan(&permission.Id,
			&permission.Name,&permission.CreatedBy,
			&permission.ModifiedBy, &permission.CreatedDate, &permission.ModifiedDate, &permission.ApplicationId); err != nil {
			return nil, fmt.Errorf("%w couldn't scan userPermissions", err)
		}
		permissions = append(permissions, permission)
	}

	log.Debugf("fetched total %d userPermissions", len(permissions))

	return permissions, nil

}
