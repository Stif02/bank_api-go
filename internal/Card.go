package internal

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

type Card struct {
	ID     uuid.UUID
	Active bool
	Block  bool
	CVV    int
}

func (c *Card) ViewCardStatus() string {
	return fmt.Sprintf("Status: Activate: %v, Blocked: %v", c.Active, c.Block)
}

func (c *Card) ActivateCard() {
	c.Active = true
}

func (c *Card) BlockedCard() {
	c.Block = true
	c.Active = false
}

func CreateCard() *Card {
	minOfCVV := 100
	maxOfCVV := 999

	return &Card{
		ID:  uuid.Must(uuid.NewRandom()),
		CVV: rand.Intn(maxOfCVV-minOfCVV) + minOfCVV,
	}
}
