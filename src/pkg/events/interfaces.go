package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayLoad() interface{}
	SetPayLoad(interface{})
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface)
	Remove(eventName string, handler EventHandlerInterface)
	Has(eventName string, handler EventHandlerInterface) bool
	Clear()
}
