package createtransaction

import (
	"context"
	"testing"

	"github.com/deividroger/ms-wallet/src/internal/entity"
	event "github.com/deividroger/ms-wallet/src/internal/events"
	"github.com/deividroger/ms-wallet/src/internal/usecase/mocks"
	"github.com/deividroger/ms-wallet/src/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {

	client1, _ := entity.NewClient("John Doe", "fake@fake.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("John Doe 2", "fake2@fake.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	input := CreateTransactionInputDto{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	transactionCreated := event.NewTransactionCreated()

	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, transactionCreated)

	output, err := uc.Execute(ctx, input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)

}
