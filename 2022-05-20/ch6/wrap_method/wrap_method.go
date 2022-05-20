package wrapmethod

import (
	"time"
)

// ----- legacy code -----
type Employee struct {
	amountDue     []int
	payDispatcher dispatcher
}

func (e Employee) Pay() {
	for _, amount := range e.amountDue {
		e.payDispatcher.pay(amount, time.Now())
	}
}

type dispatcher struct{}

func (c dispatcher) pay(amount int, date time.Time) {}

// ----- legacy code end-----

