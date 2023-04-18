package createclient

import (
	"testing"

	"github.com/deividroger/ms-wallet/src/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(&m)

	output, err := uc.Execute(CreateClientInputDto{
		Name:  "John Doe",
		Email: "teste@teste.com.br",
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotNil(t, output.ID)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "teste@teste.com.br", output.Email)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)

}
