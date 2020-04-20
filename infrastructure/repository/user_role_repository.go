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

type UserRoleRepositoryImp struct {
	mssqlConnection MSSqlConnection
}

func NewUserRoleRepository(mssqlConnection MSSqlConnection) *UserRoleRepositoryImp {
	return &UserRoleRepositoryImp{
		mssqlConnection: mssqlConnection,
	}
}

func (repository *UserRoleRepositoryImp) CreateUserRole(userRole domain.UserRole) error {
	insertPermissionQuery := `
		INSERT INTO UserRoles(Id,UserId,RoleId,CreatedBy,CreatedDate) 
        VALUES( NEXT VALUE FOR DBO.SEQ_USER_ROLES,@userId,@roleId,@createdBy,GETUTCDATE())
	`

	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, insertPermissionQuery)

	if err != nil {
		return err
	}

	defer prepareContext.Close()

	exec, err := prepareContext.Exec(sql.Named("userId", userRole.UserId),
		sql.Named("roleId", userRole.RoleId), sql.Named("createdBy", userRole.CreatedBy))

	if err != nil {
		sqlError := err.(mssql.Error)
		log.Error(sqlError.Error())
		return errors.InternalServerError(sqlError.Message)
	}

	affected, _ := exec.RowsAffected()

	if affected != 1 {
		return errors.BadRequest(fmt.Sprintf("User role could not created with given user id %d and roleId %d", userRole.UserId, userRole.RoleId))
	}

	return nil
}

func (repository *UserRoleRepositoryImp) FindAll() ([]domain.UserRole, error) {
	findAll := `
     SELECT * FROM UserRoles 
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

	userRoles := make([]domain.UserRole, 0)
	for rows.Next() {
		var userRole domain.UserRole
		if err := rows.Scan(&userRole.Id, &userRole.UserId,
			&userRole.RoleId, &userRole.CreatedBy, &userRole.CreatedDate); err != nil {
			return nil, fmt.Errorf("%w couldn't scan userRoles", err)
		}

		userRoles = append(userRoles, userRole)
	}

	log.Debugf("fetched total %d userRoles", len(userRoles))

	return userRoles, nil
}

func (repository *UserRoleRepositoryImp) DeleteByUserIdAndRoleId(userId int, roleId int) error {
	deleteUserRoleQuery := `
		DELETE FROM UserRoles 
        where UserId=@userId and RoleId=@roleId
	`

	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, deleteUserRoleQuery)

	if err != nil {
		return err
	}

	defer prepareContext.Close()

	_, err = prepareContext.Exec(sql.Named("userId", userId), sql.Named("roleId", roleId))

	if err != nil {
		sqlError := err.(mssql.Error)
		log.Error(sqlError.Error())
		return errors.InternalServerError(sqlError.Message)
	}

	return nil
}
