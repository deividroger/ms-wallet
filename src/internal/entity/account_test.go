package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "email@email.com")
	account := NewAccount(client)

	assert.NotNil(t, account)
	assert.Equal(t, account.Client.ID, client.ID)

}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)

	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "email@email.com")
	account := NewAccount(client)

	account.Credit(10)

	assert.Equal(t, float64(10), account.Balance)

}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "email@email.com")
	account := NewAccount(client)

	account.Credit(100)

	account.Debit(50)

	assert.Equal(t, float64(50), account.Balance)
}
