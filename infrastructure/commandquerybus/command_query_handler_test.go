package commandquerybus_test

import (
	"go-user-auth-api/infrastructure/commandquerybus"
	"log"
	"os"
	"testing"
)

//testing state
var accounts map[int]account
var accountCache map[int]account
var accountCreatedCount int
var handler *commandquerybus.Handler
var accountCreatedResult *accountCreated
var getAccountResult *account
var errors []error
var logger *log.Logger

//models
type account struct {
	ID      int
	Balance float64
	UserID  int
}

//commands
type createAccount struct {
	UserID int
}

//events
type accountCreated struct {
	AccountID int
	UserID    int
	Balance   float64
}

//queries
type getAccount struct {
	AccountID int
	UserID    int
}

//handlers
func createAccountHandler(command interface{}) (interface{}, error) {
	cA, ok := command.(createAccount)
	if ok == false {
		panic("do something better")
	}
	id := len(accounts) + 1
	account := account{ID: id, Balance: 0, UserID: cA.UserID}
	accounts[id] = account
	return accountCreated{AccountID: account.ID, UserID: account.UserID, Balance: 0}, nil
}

func getAccountHandler(query interface{}) (interface{}, error) {
	gA := query.(getAccount)
	return accountCache[gA.AccountID], nil
}

func incrementAccountCreatedCount(event interface{}) error {
	accountCreatedCount++
	return nil
}

func addAccountToCache(event interface{}) error {
	accountCreated := event.(accountCreated)
	accountCache[accountCreated.AccountID] = account{UserID: accountCreated.UserID, ID: accountCreated.AccountID, Balance: 0}
	return nil
}

//steps
func givenStateSetup() {
	accountCreatedCount = 0
	accounts = map[int]account{}
	accountCache = map[int]account{}
	logger = log.New(os.Stdout, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)
	handler = commandquerybus.New()
}

func givenEventHandlersRegistered() {
	handler.RegisterEventHandler(accountCreated{}, incrementAccountCreatedCount)
	handler.RegisterEventHandler(accountCreated{}, addAccountToCache)
}

func givenCommandHandlersRegistered() {
	handler.RegisterCommandHandler(createAccount{}, createAccountHandler)
}

func givenQueryHandlersRegistered() {
	handler.RegisterQueryHandler(getAccount{}, getAccountHandler)
}

func whenIHandleCommand() {
	result, err := handler.Handle(createAccount{UserID: 1})
	if len(err) == 0 {
		acr := result.(accountCreated)
		accountCreatedResult = &acr
	}
	errors = err
}

func whenIHandleQuery() {
	gA := getAccount{UserID: accountCreatedResult.UserID, AccountID: accountCreatedResult.AccountID}
	result, _ := handler.Handle(gA)
	gar := result.(account)
	getAccountResult = &gar
}

func thenTheCommandAndEventHandlersAreCalled(t *testing.T) {
	if accountCreatedResult.AccountID != 1 || accountCreatedCount != 1 || len(accountCache) != 1 {
		t.Fatalf("ThenTheCommandAndEventHandlersAreCalled failed")
	}
}

func thenTheCommandHandlersAreCalled(t *testing.T) {
	if accountCreatedResult.AccountID != 1 || accountCreatedCount != 0 || len(accountCache) != 0 {
		t.Fatalf("ThenTheCommandHandlersAreCalled failed")
	}
}

func thenTheQueryHandlersAreCalled(t *testing.T) {

	if getAccountResult.ID != accountCreatedResult.AccountID || getAccountResult.UserID != accountCreatedResult.UserID || getAccountResult.Balance != accountCreatedResult.Balance {
		t.Fatalf("ThenTheQueryHandlersAreCalled falied")
	}
}

func thenAnErrorIsReturned(t *testing.T) {
	if len(errors) == 0 {
		t.Fatalf("ThenAnErrorIsReturned falied")
	}
}

func TestReturnsErrorWhenNoHandlers(t *testing.T) {
	givenStateSetup()
	whenIHandleCommand()
	thenAnErrorIsReturned(t)
}

func TestReturnsWhenNoEventHandlers(t *testing.T) {
	givenStateSetup()
	givenCommandHandlersRegistered()
	whenIHandleCommand()
	thenTheCommandHandlersAreCalled(t)
}

func TestEndToEnd(t *testing.T) {
	givenStateSetup()
	givenCommandHandlersRegistered()
	givenEventHandlersRegistered()
	givenQueryHandlersRegistered()
	whenIHandleCommand()
	whenIHandleQuery()
	thenTheCommandAndEventHandlersAreCalled(t)
	thenTheQueryHandlersAreCalled(t)
}