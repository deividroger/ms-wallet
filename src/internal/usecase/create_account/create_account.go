package createaccount

import (
	"fmt"

	"github.com/deividroger/ms-wallet/src/internal/entity"
	"github.com/deividroger/ms-wallet/src/internal/gateway"
)

type CreateAccountInputDto struct {
	ClientID string `json:"client_id"`
}

type CreateAccountOutputDto struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: a,
		ClientGateway:  c,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDto) (*CreateAccountOutputDto, error) {
	client, err := uc.ClientGateway.Get(input.ClientID)
	fmt.Println(client)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)

	err = uc.AccountGateway.Save(account)

	if err != nil {

		return nil, err
	}

	return &CreateAccountOutputDto{ID: account.ID}, nil

}
