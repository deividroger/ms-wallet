package event

import "time"

type TransactionCreated struct {
	Name    string
	Payload interface{}
}

func NewTransactionCreated() *TransactionCreated {
	return &TransactionCreated{
		Name: "TransactionCreated",
	}
}

func (te *TransactionCreated) GetName() string {
	return te.Name
}

func (te *TransactionCreated) GetPayLoad() interface{} {
	return te.Payload
}
func (te *TransactionCreated) GetDateTime() time.Time {
	return time.Now()
}
func (te *TransactionCreated) SetPayLoad(payload interface{}) {
	te.Payload = payload
}
