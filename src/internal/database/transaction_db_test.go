package database

import (
	"database/sql"
	"testing"

	"github.com/deividroger/ms-wallet/src/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db

	_, err = db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATETIME)")

	s.Nil(err)

	_, err = db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance FLOAT, created_at DATETIME)")

	s.Nil(err)
	_, err = db.Exec(`
		CREATE TABLE transactions (
							id VARCHAR(255),
							account_id_from VARCHAR(255),
							account_id_to VARCHAR(255),
							amount FLOAT,
							created_at DATETIME
						)`)
	s.Nil(err)

	client, err := entity.NewClient("John Doe", "client1@client1.com")

	s.Nil(err)

	s.client = client

	client2, err := entity.NewClient("John Doe 2", "client2@client2.com")
	s.Nil(err)

	s.client2 = client2

	accountFrom := entity.NewAccount(s.client)
	accountFrom.Credit(1000)

	s.accountFrom = accountFrom

	account2 := entity.NewAccount(s.client2)
	account2.Credit(1000)

	s.accountTo = account2

	s.transactionDB = NewTransactionDB(db)

}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuit(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestSave() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)

	err = s.transactionDB.Create(transaction)

	s.Nil(err)
}
