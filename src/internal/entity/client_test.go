package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "email@email.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, client.Name, "John Doe")
	assert.Equal(t, client.Email, "email@email.com")

}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "email@email.com")

	err := client.Update("Nome 2", "gmail@gmail.com")

	assert.Nil(t, err)
	assert.Equal(t, client.Name, "Nome 2")
	assert.Equal(t, client.Email, "gmail@gmail.com")

}

func TestUpdateClientWhenArgsAreInvalid(t *testing.T) {

	client, _ := NewClient("John Doe", "email@email.com")

	err := client.Update("", "")

	assert.NotNil(t, err)
	assert.Error(t, err, "name is invalid")

}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "email@email.com")

	account := NewAccount(client)

	err := client.AddAccount(account)

	assert.Nil(t, err)

	assert.Equal(t, len(client.Accounts), 1)

}
