package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClient(t *testing.T) {
	client, err := NewClient("John", "john@doe.com")

	if err != nil {
		t.Errorf("Error creating client: %v", err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John", client.Name)
	assert.Equal(t, "john@doe.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.cmo")
	err := client.Update("John", "john@hello.com")

	assert.Nil(t, err)
	assert.Equal(t, "John", client.Name)
	assert.Equal(t, "john@hello.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.cmo")
	err := client.Update("", "john@hello.com")

	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.cmo")
	account := NewAccount(client)

	err := client.AddAccount(account)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))

}
