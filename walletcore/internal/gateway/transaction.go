package gateway

import "github.com/sidroniolima/go-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(account *entity.Transaction) error
}
