package createtransaction

import (
	"github.com/deividroger/ms-wallet/src/internal/entity"
	"github.com/deividroger/ms-wallet/src/internal/gateway"
	"github.com/deividroger/ms-wallet/src/pkg/events"
)

type CreateTransactionInputDto struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDto struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
	eventDispatcher events.EventDispatcherInterface,

	transactionCreated events.EventInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {

	accountFrom, err := uc.AccountGateway.FindById(input.AccountIDFrom)

	if err != nil {
		return nil, err
	}
	account2, err := uc.AccountGateway.FindById(input.AccountIDTo)

	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, account2, input.Amount)

	if err != nil {
		return nil, err
	}

	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	uc.TransactionCreated.SetPayLoad(transaction)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

	return &CreateTransactionOutputDto{
		ID: transaction.ID,
	}, nil

}
