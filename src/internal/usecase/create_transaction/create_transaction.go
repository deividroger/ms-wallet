package createtransaction

import (
	"context"

	"github.com/deividroger/ms-wallet/src/internal/entity"
	"github.com/deividroger/ms-wallet/src/internal/gateway"
	"github.com/deividroger/ms-wallet/src/pkg/events"
	"github.com/deividroger/ms-wallet/src/pkg/uow"
)

type CreateTransactionInputDto struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDto struct {
	ID            string
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	Uow                uow.UowInterface
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	Uow uow.UowInterface,
	eventDispatcher events.EventDispatcherInterface,

	transactionCreated events.EventInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:                Uow,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (uc *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {

	output := &CreateTransactionOutputDto{}

	err := uc.Uow.Do(ctx, func(_ *uow.Uow) error {

		accountRepository := uc.getAccountRepository(ctx)
		transactionRepository := uc.getTransactionRepository(ctx)

		accountFrom, err := accountRepository.FindById(input.AccountIDFrom)

		if err != nil {
			return err
		}
		account2, err := accountRepository.FindById(input.AccountIDTo)

		if err != nil {
			return err
		}

		transaction, err := entity.NewTransaction(accountFrom, account2, input.Amount)

		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountFrom)

		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(account2)

		if err != nil {
			return err
		}

		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}
		output.ID = transaction.ID
		output.AccountIDFrom = input.AccountIDFrom
		output.AccountIDTo = input.AccountIDTo
		output.Amount = input.Amount
		return nil

	})

	if err != nil {
		return nil, err
	}

	uc.TransactionCreated.SetPayLoad(output)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

	return output, nil

}

func (uc *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")

	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := uc.Uow.GetRepository(ctx, "TransactionDB")

	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)

}
