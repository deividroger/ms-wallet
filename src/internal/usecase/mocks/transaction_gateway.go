package mocks

import (
	"github.com/deividroger/ms-wallet/src/internal/entity"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	ret := m.Called(transaction)
	return ret.Error(0)
}
