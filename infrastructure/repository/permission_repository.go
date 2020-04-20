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

type PermissionRepositoryImp struct {
	mssqlConnection MSSqlConnection
}

func NewPermissionRepository(mssqlConnection MSSqlConnection) *PermissionRepositoryImp {
	return &PermissionRepositoryImp{
		mssqlConnection: mssqlConnection,
	}
}

func (repository *PermissionRepositoryImp) CreatePermission(permission domain.Permission) error {
	insertPermissionQuery := `
		INSERT INTO Permissions(Id,Name,CreatedBy, ModifiedBy, CreatedDate, ModifiedDate) 
        VALUES( NEXT VALUE FOR DBO.SEQ_PERMISSIONS,@name,@createdBy,@modifiedBy,GETUTCDATE(),GETUTCDATE())
	`

	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, insertPermissionQuery)

	if err != nil {
		return err
	}

	defer prepareContext.Close()

	_, err = prepareContext.Exec(sql.Named("name", permission.Name), sql.Named("createdBy", permission.CreatedBy), sql.Named("modifiedBy", permission.ModifiedBy))

	if err != nil {
		sqlError := err.(mssql.Error)
		log.Error(sqlError.Error())
		return errors.InternalServerError(sqlError.Message)
	}

	return nil
}

func (repository *PermissionRepositoryImp) GetPermissionById(id int) (*domain.Permission, error) {
	findPermissionByIdQuery := `
		SELECT * FROM Permissions 
        WHERE Id = @id
	`

	connCtx, cancel := context.WithTimeout(context.Background(), defaultQueryTimeout)
	defer cancel()

	conn, err := repository.mssqlConnection.db.Conn(connCtx)
	if err != nil {
		return nil, fmt.Errorf("%w couldn't open connection from pool", err)
	}

	defer conn.Close()

	prepareContext, err := conn.PrepareContext(connCtx, findPermissionByIdQuery)
	if err != nil {
		return nil, err
	}

	defer prepareContext.Close()

	var permission domain.Permission
	err = prepareContext.QueryRowContext(connCtx, sql.Named("id", id)).
		Scan(&permission.Id, &permission.Name, &permission.CreatedBy,
			&permission.ModifiedBy, &permission.CreatedDate, &permission.ModifiedDate)
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (repository *PermissionRepositoryImp) DeletePermissionById(id int) error {

	deletePermissionQuery := `
		DELETE FROM Permissions 
        where Id=@id and Id not in (select PermissionId from RolePermissions where PermissionId=@id)
	`

	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, deletePermissionQuery)

	if err != nil {
		return err
	}

	defer prepareContext.Close()

	exec, err := prepareContext.Exec(sql.Named("id", id))

	if err != nil {
		return err
	}

	affected, _ := exec.RowsAffected()

	if affected != 1 {
		return errors.BadRequest(fmt.Sprintf("Permission could not deleted with given id %d it is associated with any role or not exist", id))
	}

	return nil
}

func (repository *PermissionRepositoryImp) FindAll() ([]domain.Permission, error) {
	findAll := `
     SELECT * FROM Permissions 
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

	permissions := make([]domain.Permission, 0)
	for rows.Next() {
		var permission domain.Permission
		if err := rows.Scan(&permission.Id, &permission.Name, &permission.CreatedBy,
			&permission.ModifiedBy, &permission.CreatedDate, &permission.ModifiedDate); err != nil {
			return nil, fmt.Errorf("%w couldn't scan permissions", err)
		}

		permissions = append(permissions, permission)
	}

	log.Debugf("fetched total %d permissions", len(permissions))

	return permissions, nil
}
