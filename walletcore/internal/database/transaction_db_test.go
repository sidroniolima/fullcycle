package database

import (
	"database/sql"
	"testing"

	"github.com/sidroniolima/go-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	clientFrom    *entity.Client
	clientTo      *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	db.Exec("CREATE TABLE accounts (id VARCHAR(255) NOT NULL, client_id VARCHAR(255) NOT NULL, balance FLOAT, created_at DATE)")
	db.Exec("CREATE TABLE clients (id VARCHAR(255) NOT NULL, name VARCHAR(255), email VARCHAR(255), created_at DATE)")
	db.Exec("CREATE TABLE transactions (id VARCHAR(255) NOT NULL, account_id_from VARCHAR(255) NOT NULL, account_id_to VARCHAR(255) NOT NULL, amount float, created_at DATE)")

	s.transactionDB = NewTransactionDB(db)
	s.clientTo, _ = entity.NewClient("John", "john@doe.com")
	s.clientFrom, _ = entity.NewClient("John", "john@doe.com")

	s.accountTo = entity.NewAccount(s.clientTo)
	s.accountFrom = entity.NewAccount(s.clientFrom)
	s.accountFrom.Balance = 1000
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)

	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
