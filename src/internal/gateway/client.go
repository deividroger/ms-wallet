package gateway

import "github.com/deividroger/ms-wallet/src/internal/entity"

type ClientGateway interface {
	Get(int string) (*entity.Client, error)
	Save(*entity.Client) error
}
