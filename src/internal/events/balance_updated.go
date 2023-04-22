package event

import "time"

type BalanceUpdated struct {
	Name    string
	Payload interface{}
}

func NewBalanceUpdated() *TransactionCreated {
	return &TransactionCreated{
		Name: "BalanceUpdated",
	}
}

func (b *BalanceUpdated) GetName() string {
	return b.Name
}

func (b *BalanceUpdated) GetPayLoad() interface{} {
	return b.Payload
}
func (b *BalanceUpdated) GetDateTime() time.Time {
	return time.Now()
}
func (b *BalanceUpdated) SetPayLoad(payload interface{}) {
	b.Payload = payload
}
