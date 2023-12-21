package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	clientFrom, _ := NewClient("John Doe", "john@doe.com")
	accountFrom := NewAccount(clientFrom)

	clientTo, _ := NewClient("Silvia Saint", "silvia@saint.com")
	accountTo := NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	transation, err := NewTransaction(accountFrom, accountTo, 100)

	assert.Nil(t, err)
	assert.NotNil(t, transation)
	assert.Equal(t, float64(900), accountFrom.Balance)
	assert.Equal(t, float64(1100), accountTo.Balance)
}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	clientFrom, _ := NewClient("John Doe", "john@doe.com")
	accountFrom := NewAccount(clientFrom)

	clientTo, _ := NewClient("Silvia Saint", "silvia@saint.com")
	accountTo := NewAccount(clientTo)

	accountFrom.Credit(100)
	accountTo.Credit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, 200)

	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, float64(100), accountFrom.Balance)
	assert.Equal(t, float64(1000), accountTo.Balance)
}
