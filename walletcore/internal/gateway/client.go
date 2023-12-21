package gateway

import "github.com/sidroniolima/go-ms-wallet/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(clent *entity.Client) error
}
