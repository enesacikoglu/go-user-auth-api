package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/labstack/gommon/log"
	"go-user-auth-api/domain"
	"go-user-auth-api/infrastructure/errors"
)

type RolePermissionRepositoryImp struct {
	mssqlConnection MSSqlConnection
}

func NewRolePermissionRepository(mssqlConnection MSSqlConnection) *RolePermissionRepositoryImp {
	return &RolePermissionRepositoryImp{
		mssqlConnection: mssqlConnection,
	}
}

func (repository *RolePermissionRepositoryImp) CreateRolePermission(rolePermission domain.RolePermission) error {
	insertPermissionQuery := `
		INSERT INTO RolePermissions(Id,RoleId,PermissionId,CreatedBy,CreatedDate,ApplicationId) 
        VALUES( NEXT VALUE FOR DBO.SEQ_ROLE_PERMISSIONS,@roleId,@permissionId,@createdBy,GETUTCDATE(),@applicationId)
	`

	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, insertPermissionQuery)

	if err != nil {
		return err
	}

	defer prepareContext.Close()

	_, err = prepareContext.Exec(sql.Named("roleId", rolePermission.RoleId),
		sql.Named("permissionId", rolePermission.PermissionId),
		sql.Named("createdBy", rolePermission.CreatedBy), sql.Named("applicationId", rolePermission.ApplicationId))

	if err != nil {
		sqlError := err.(mssql.Error)
		log.Error(sqlError.Error())
		return errors.InternalServerError(sqlError.Message)
	}

	return nil
}

func (repository *RolePermissionRepositoryImp) FindAll() ([]domain.RolePermission, error) {
	findAll := `
     SELECT * FROM RolePermissions 
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

	rows, err := conn.QueryContext(queryCtx, findAll)
	if err != nil {
		return nil, fmt.Errorf("%w  connection executure query", err)
	}

	rolePermissions := make([]domain.RolePermission, 0)
	for rows.Next() {
		var rolePermission domain.RolePermission
		if err := rows.Scan(&rolePermission.Id, &rolePermission.RoleId,
			&rolePermission.PermissionId, &rolePermission.CreatedBy,
			&rolePermission.CreatedDate, &rolePermission.ApplicationId); err != nil {
			return nil, fmt.Errorf("%w couldn't scan permissions", err)
		}

		rolePermissions = append(rolePermissions, rolePermission)
	}

	log.Debugf("fetched total %d rolePermissions", len(rolePermissions))

	return rolePermissions, nil
}
func (repository *RolePermissionRepositoryImp) DeleteByRoleIdAndPermissionIdAndApplicationId(roleId int, permissionId int, applicationId int) error {
	deleteRolePermissionQuery := `
		DELETE FROM RolePermissions 
        where RoleId=@roleId and PermissionId=@permissionId and ApplicationId=@applicationId
	`

	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, deleteRolePermissionQuery)

	if err != nil {
		return err
	}

	defer prepareContext.Close()

	_, err = prepareContext.Exec(sql.Named("roleId", roleId), sql.Named("permissionId", permissionId), sql.Named("applicationId", applicationId))

	if err != nil {
		sqlError := err.(mssql.Error)
		log.Error(sqlError.Error())
		return errors.InternalServerError(sqlError.Message)
	}

	return nil
}
