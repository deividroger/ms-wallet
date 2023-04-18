package createaccount

import (
	"testing"

	"github.com/deividroger/ms-wallet/src/internal/entity"
	"github.com/deividroger/ms-wallet/src/internal/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {

	client, _ := entity.NewClient("John Doe", "teste@teste.com.br")

	clientGatewayMock := &mocks.ClientGatewayMock{}
	clientGatewayMock.On("Get", client.ID).Return(client, nil)

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, clientGatewayMock)

	inputDto := CreateAccountInputDto{ClientID: client.ID}

	output, err := uc.Execute(inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)

	clientGatewayMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)

	clientGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)

}
