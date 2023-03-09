package database

import (
	"database/sql"
	"testing"

	"github.com/deividroger/ms-wallet/src/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	AccountDB *AccountDB
	client    *entity.Client
}

func (suite *AccountDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")

	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATETIME)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance FLOAT, created_at DATETIME)")

	suite.AccountDB = NewAccountDb(db)

	suite.client, _ = entity.NewClient("John Doe", "email@email.com")
}

func (suite *AccountDBTestSuite) TearDownTest() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
	suite.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.AccountDB.Save(account)

	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindById() {
	_, err := s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)

	s.Nil(err)

	account := entity.NewAccount(s.client)
	err = s.AccountDB.Save(account)
	s.Nil(err)

	accountDB, err := s.AccountDB.FindById(account.ID)

	s.Nil(err)

	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.Client.ID, accountDB.Client.ID)

}
