package createtransaction

import (
	"testing"

	"github.com/sidroniolima/go-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	clientFrom, _ := entity.NewClient("John", "john@gmail.com")
	mockAccountFrom := entity.NewAccount(clientFrom)
	mockAccountFrom.Credit(5000)

	clientTo, _ := entity.NewClient("Joseph", "jose@gmail.com")
	mockAccountTo := entity.NewAccount(clientTo)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Save", mock.Anything).Return(nil)
	accountGatewayMock.On("FindByID", mockAccountFrom.ID).Return(mockAccountFrom, nil)
	accountGatewayMock.On("FindByID", mockAccountTo.ID).Return(mockAccountTo, nil)

	transactionGatewayMock := &TransactionGatewayMock{}
	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	useCaseCreateTransaction := NewCreateTransactionUseCase(accountGatewayMock, transactionGatewayMock)

	inputDTO := CreateTransactionInputDTO{AccountFromID: mockAccountFrom.ID, AccountToID: mockAccountTo.ID, Amount: 100}

	output, err := useCaseCreateTransaction.Execute(inputDTO)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotNil(t, output.ID)
	accountGatewayMock.AssertExpectations(t)
	transactionGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 2)
	transactionGatewayMock.AssertNumberOfCalls(t, "Create", 1)
}
