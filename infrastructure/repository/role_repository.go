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

type RoleRepositoryImp struct {
	mssqlConnection MSSqlConnection
}

func NewRoleRepository(mssqlConnection MSSqlConnection) *RoleRepositoryImp {
	return &RoleRepositoryImp{
		mssqlConnection: mssqlConnection,
	}
}

func (repository *RoleRepositoryImp) CreateRole(role domain.Role) error {
	insertPermissionQuery := `
		INSERT INTO ROLES(Id,Name,CreatedBy, ModifiedBy, CreatedDate, ModifiedDate) 
        VALUES( NEXT VALUE FOR DBO.SEQ_ROLES,@name,@createdBy,@modifiedBy,GETUTCDATE(),GETUTCDATE())
	`
	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, insertPermissionQuery)

	if err != nil {
		return err
	}

	defer prepareContext.Close()

	_, err = prepareContext.Exec(sql.Named("name", role.Name),
		sql.Named("createdBy", role.CreatedBy), sql.Named("modifiedBy", role.ModifiedBy))

	if err != nil {
		sqlError := err.(mssql.Error)
		log.Error(sqlError.Error())
		return errors.InternalServerError(sqlError.Message)
	}

	return nil
}

func (repository *RoleRepositoryImp) GetRoleById(id int) (*domain.Role, error) {
	findRoleByIdQuery := `
		SELECT * FROM Roles 
        WHERE Id = @id
	`

	connCtx, cancel := context.WithTimeout(context.Background(), defaultQueryTimeout)
	defer cancel()

	conn, err := repository.mssqlConnection.db.Conn(connCtx)
	if err != nil {
		return nil, fmt.Errorf("%w couldn't open connection from pool", err)
	}

	defer conn.Close()

	prepareContext, err := conn.PrepareContext(connCtx, findRoleByIdQuery)
	if err != nil {
		return nil, err
	}

	defer prepareContext.Close()

	var role domain.Role
	err = prepareContext.QueryRowContext(connCtx, sql.Named("id", id)).
		Scan(&role.Id, &role.Name, &role.CreatedBy,
			&role.ModifiedBy, &role.CreatedDate, &role.ModifiedDate)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (repository *RoleRepositoryImp) DeleteRoleById(id int) error {

	deleteRoleQuery := `
		DELETE FROM Roles 
        where Id=@id and Id not in (select RoleId from RolePermissions where RoleId=@id)
	`

	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, deleteRoleQuery)

	if err != nil {
		return err
	}

	defer prepareContext.Close()

	exec, err := prepareContext.Exec(sql.Named("id", id))

	if err != nil {
		sqlError := err.(mssql.Error)
		log.Error(sqlError.Error())
		return errors.InternalServerError(sqlError.Message)
	}

	affected, _ := exec.RowsAffected()

	if affected != 1 {
		return errors.BadRequest(fmt.Sprintf("Role could not deleted with given id %d it is associated with any permission or not exist", id))
	}

	return nil
}

func (repository *RoleRepositoryImp) FindAll() ([]domain.Role, error) {
	findAll := `
     SELECT * FROM Roles 
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

	roles := make([]domain.Role, 0)
	for rows.Next() {
		var role domain.Role
		if err := rows.Scan(&role.Id, &role.Name, &role.CreatedBy,
			&role.ModifiedBy, &role.CreatedDate, &role.ModifiedDate); err != nil {
			return nil, fmt.Errorf("%w couldn't scan roles", err)
		}

		roles = append(roles, role)
	}

	log.Debugf("fetched total %d roles", len(roles))

	return roles, nil
}
