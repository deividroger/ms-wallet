package gateway

import "github.com/deividroger/ms-wallet/src/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
