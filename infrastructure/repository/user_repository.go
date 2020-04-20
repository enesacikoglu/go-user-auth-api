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
	"time"
)

const (
	defaultQueryTimeout = 5 * time.Second
)

type UserRepositoryImp struct {
	mssqlConnection MSSqlConnection
}

func NewUserRepository(mssqlConnection MSSqlConnection) *UserRepositoryImp {
	return &UserRepositoryImp{
		mssqlConnection: mssqlConnection,
	}
}

func (repository *UserRepositoryImp) CreateUser(user domain.User) error {
	insertUserQuery := `
		INSERT INTO Users(Id,Email, Name, Surname,CreatedBy, ModifiedBy, CreatedDate, ModifiedDate) 
        VALUES(NEXT VALUE FOR DBO.SEQ_USERS, @email,@name,@surname,@createdBy,@modifiedBy,GETUTCDATE(),GETUTCDATE())
	`
	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, insertUserQuery)

	if err != nil {
		return err
	}

	defer prepareContext.Close()

	_, err = prepareContext.Exec(sql.Named("email", user.Email), sql.Named("name", user.Name),
		sql.Named("surname", user.Surname), sql.Named("createdBy", user.CreatedBy), sql.Named("modifiedBy", user.ModifiedBy))

	if err != nil {
		sqlError := err.(mssql.Error)
		log.Error(sqlError.Error())
		return errors.InternalServerError(sqlError.Message)
	}

	return nil
}

func (repository *UserRepositoryImp) GetUser(id int) (domain.UserRolePermissions, error) {
	findUserRolePermissionByIdQuery := `
       select 
       us.Id             as UserId,
       us.Email          as Email,
       us.Name           as UserName,
       us.Surname        as UserSurname,
       us.CreatedBy      as UserCreatedBy,
       us.ModifiedBy     as UserModifiedBy,
       us.CreatedDate    as UserCreatedDate,
       us.ModifiedDate   as UserModifiedDate,
       rl.Id             as RoleId,
       rl.Name           as RoleName,
       rl.CreatedBy      as RoleCreatedBy,
       rl.ModifiedBy     as RoleModifiedBy,
       rl.CreatedDate    as RoleCreatedDate,
       rl.ModifiedDate   as RoleModifiedDate,
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
    WHERE us.Id = @id
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

	rows, err := conn.QueryContext(queryCtx, findUserRolePermissionByIdQuery, sql.Named("id", id))
	if err != nil {
		return nil, fmt.Errorf("%w  connection executure query", err)
	}

	var userRolePermissions domain.UserRolePermissions
	for rows.Next() {
		var userRolePermission domain.UserRolePermission
		if err := rows.Scan(&userRolePermission.UserId, &userRolePermission.Email, &userRolePermission.UserName,
			&userRolePermission.Surname, &userRolePermission.UserCreatedBy, &userRolePermission.UserModifiedBy,
			&userRolePermission.UserCreatedDate, &userRolePermission.UserModifiedDate, &userRolePermission.RoleId,
			&userRolePermission.RoleName, &userRolePermission.RoleCreatedBy, &userRolePermission.RoleModifiedBy,
			&userRolePermission.RoleCreatedDate, &userRolePermission.RoleModifiedDate, &userRolePermission.PermissionId,
			&userRolePermission.PermissionName, &userRolePermission.PermissionCreatedBy, &userRolePermission.PermissionModifiedBy, &userRolePermission.PermissionCreatedDate, &userRolePermission.PermissionModifiedDate,
			&userRolePermission.ApplicationId); err != nil {
			return nil, fmt.Errorf("%w couldn't scan userRolePermission", err)
		}

		userRolePermissions = append(userRolePermissions, userRolePermission)
	}

	log.Debugf("fetched total %d userRolePermissions", len(userRolePermissions))

	return userRolePermissions, nil
}

func (repository *UserRepositoryImp) GetUserById(id int) (*domain.User, error) {
	findUserByIdQuery := `
		SELECT * FROM Users 
        WHERE Id = @id
	`
	connCtx, cancel := context.WithTimeout(context.Background(), defaultQueryTimeout)
	defer cancel()

	conn, err := repository.mssqlConnection.db.Conn(connCtx)
	if err != nil {
		return nil, fmt.Errorf("%w couldn't open connection from pool", err)
	}

	defer conn.Close()

	prepareContext, err := conn.PrepareContext(connCtx, findUserByIdQuery)
	if err != nil {
		return nil, err
	}

	defer prepareContext.Close()

	var user domain.User
	err = prepareContext.QueryRowContext(connCtx, sql.Named("id", id)).
		Scan(&user.Id, &user.Email, &user.Name, &user.Surname, &user.CreatedBy, &user.ModifiedBy, &user.CreatedDate, &user.ModifiedDate)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *UserRepositoryImp) GetUserByEmail(email string) (*domain.User, error) {
	findUserByEmailQuery := `
		SELECT * FROM Users 
        WHERE Email = @email
	`
	connCtx, cancel := context.WithTimeout(context.Background(), defaultQueryTimeout)
	defer cancel()

	conn, err := repository.mssqlConnection.db.Conn(connCtx)
	if err != nil {
		return nil, fmt.Errorf("%w couldn't open connection from pool", err)
	}

	defer conn.Close()

	prepareContext, err := conn.PrepareContext(connCtx, findUserByEmailQuery)
	if err != nil {
		return nil, err
	}

	defer prepareContext.Close()

	var user domain.User
	err = prepareContext.QueryRowContext(connCtx, sql.Named("email", email)).
		Scan(&user.Id, &user.Email, &user.Name, &user.Surname, &user.CreatedBy, &user.ModifiedBy, &user.CreatedDate, &user.ModifiedDate)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepositoryImp) UpdateUserById(id int, user domain.User) error {
	updateUserQuery := `
		UPDATE Users SET Email=@email,Name=@name,Surname=@surname,CreatedBy=@createdBy, 
        ModifiedBy=@modifiedBy,CreatedDate=@createdDate,ModifiedDate=GETUTCDATE()
        where Id=@id
	`


	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, updateUserQuery)

	if err != nil {
		return err
	}

	defer prepareContext.Close()

	_, err = prepareContext.Exec(sql.Named("id", id), sql.Named("name", user.Name), sql.Named("email", user.Email),
		sql.Named("surname", user.Surname), sql.Named("createdBy", user.CreatedBy), sql.Named("modifiedBy", user.ModifiedBy),
		sql.Named("createdDate", user.CreatedDate))

	if err != nil {
		sqlError := err.(mssql.Error)
		log.Error(sqlError.Error())
		return errors.InternalServerError(sqlError.Message)
	}

	return nil
}

func (repository *UserRepositoryImp) DeleteUserById(id int) error {

	updateUserQuery := `
		DELETE FROM USERS 
        where Id=@id and Id not in (select UserId from UserRoles where UserId=@id)
	`


	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, updateUserQuery)

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
		return errors.BadRequest(fmt.Sprintf("User could not deleted with given id %d it is associated with any role or not exist", id))
	}

	return nil
}

func (repository *UserRepositoryImp) GetUserRolesById(id int) ([]domain.Role, error) {
	findUserRoleByIdQuery := `
       select 
       rl.Id             as RoleId,
       rl.Name           as RoleName,
       rl.CreatedBy      as RoleCreatedBy,
       rl.ModifiedBy     as RoleModifiedBy,
       rl.CreatedDate    as RoleCreatedDate,
       rl.ModifiedDate   as RoleModifiedDate
	from Users us
       LEFT JOIN UserRoles ur on us.Id = ur.UserId
            JOIN Roles rl ON ur.RoleId = rl.Id
    WHERE us.Id = @id
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

	rows, err := conn.QueryContext(queryCtx, findUserRoleByIdQuery, sql.Named("id", id))
	if err != nil {
		return nil, fmt.Errorf("%w  connection executure query", err)
	}

	roles := make([]domain.Role, 0)
	for rows.Next() {
		var role domain.Role
		if err := rows.Scan(&role.Id,
			&role.Name, &role.CreatedBy, &role.ModifiedBy, &role.CreatedDate, &role.ModifiedDate); err != nil {
			return nil, fmt.Errorf("%w couldn't scan userRole", err)
		}
		roles = append(roles, role)
	}

	log.Debugf("fetched total %d userRole", len(roles))

	return roles, nil
}

func (repository *UserRepositoryImp) GetUserRolePermissionsByUserIdAndRoleId(id int, roleId int) ([]domain.Permission, error) {
	findUserRolePermissionByIdAndRoleIdQuery := `
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
    WHERE us.Id = @id and rl.Id=@roleId
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

	rows, err := conn.QueryContext(queryCtx, findUserRolePermissionByIdAndRoleIdQuery,
		sql.Named("id", id),
		sql.Named("roleId", roleId))

	if err != nil {
		sqlError := err.(mssql.Error)
		log.Error(sqlError.Error())
		return nil, errors.InternalServerError(sqlError.Message)
	}

	permissions := make([]domain.Permission, 0)
	for rows.Next() {
		var permission domain.Permission
		if err := rows.Scan(&permission.Id, &permission.Name, &permission.CreatedBy,
			&permission.ModifiedBy, &permission.CreatedDate, &permission.ModifiedDate, &permission.ApplicationId); err != nil {
			return nil, fmt.Errorf("%w couldn't scan userRolePermissions", err)
		}
		permissions = append(permissions, permission)
	}

	log.Debugf("fetched total %d userRolePermissions", len(permissions))

	return permissions, nil
}

func (repository *UserRepositoryImp) FindAll() ([]domain.User, error) {
	findAll := `
     SELECT * FROM Users 
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

	users := make([]domain.User, 0)
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.Surname,
			&user.CreatedBy, &user.ModifiedBy, &user.CreatedDate, &user.ModifiedDate); err != nil {
			return nil, fmt.Errorf("%w couldn't scan roles", err)
		}
		users = append(users, user)
	}

	log.Debugf("fetched total %d users", len(users))

	return users, nil
}
