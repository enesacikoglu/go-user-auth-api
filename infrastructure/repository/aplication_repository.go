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

type ApplicationRepositoryImp struct {
	mssqlConnection MSSqlConnection
}

func NewApplicationRepository(mssqlConnection MSSqlConnection) *ApplicationRepositoryImp {
	return &ApplicationRepositoryImp{
		mssqlConnection: mssqlConnection,
	}
}

func (repository *ApplicationRepositoryImp) CreateApplication(application domain.Application) error {
	insertApplicationsQuery := `
		INSERT INTO Applications(Id,Name,CreatedBy, ModifiedBy, CreatedDate, ModifiedDate) 
        VALUES( NEXT VALUE FOR DBO.SEQ_APPLICATIONS,@name,@createdBy,@modifiedBy,GETUTCDATE(),GETUTCDATE())
	`
	ctx := context.Background()

	prepareContext, err := repository.mssqlConnection.db.PrepareContext(ctx, insertApplicationsQuery)

	if err != nil {
		return err
	}

	defer prepareContext.Close()

	_, err = prepareContext.Exec(sql.Named("name", application.Name),
		sql.Named("createdBy", application.CreatedBy), sql.Named("modifiedBy", application.ModifiedBy))

	if err != nil {
		sqlError := err.(mssql.Error)
		log.Error(sqlError.Error())
		return errors.InternalServerError(sqlError.Message)
	}

	return nil
}

func (repository *ApplicationRepositoryImp) GetApplicationById(id int) (*domain.Application, error) {
	findRoleByIdQuery := `
		SELECT * FROM Applications 
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

	var application domain.Application
	err = prepareContext.QueryRowContext(connCtx, sql.Named("id", id)).
		Scan(&application.Id, &application.Name, &application.CreatedBy,
			&application.ModifiedBy, &application.CreatedDate, &application.ModifiedDate)
	if err != nil {
		return nil, err
	}
	return &application, nil
}

func (repository *ApplicationRepositoryImp) FindAll() ([]domain.Application, error) {
	findAll := `
     SELECT * FROM Applications 
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

	applications := make([]domain.Application, 0)
	for rows.Next() {
		var application domain.Application
		if err := rows.Scan(&application.Id, &application.Name, &application.CreatedBy,
			&application.ModifiedBy, &application.CreatedDate, &application.ModifiedDate); err != nil {
			return nil, fmt.Errorf("%w couldn't scan applications", err)
		}

		applications = append(applications, application)
	}

	log.Debugf("fetched total %d applications", len(applications))

	return applications, nil
}
