package database

import (
	"database/sql"
	"testing"

	"github.com/sidroniolima/go-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	db.Exec("CREATE TABLE accounts (id VARCHAR(255) NOT NULL, client_id VARCHAR(255) NOT NULL, balance FLOAT, created_at DATE)")
	db.Exec("CREATE TABLE clients (id VARCHAR(255) NOT NULL, name VARCHAR(255), email VARCHAR(255), created_at DATE)")

	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("John", "john@doe.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {

	account := entity.NewAccount(s.client)

	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {

	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt,
	)

	account := entity.NewAccount(s.client)

	err := s.accountDB.Save(account)
	s.Nil(err)

	accountDB, err := s.accountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(account.Balance, accountDB.Balance)
	s.NotNil(accountDB.CreatedAt)
}
