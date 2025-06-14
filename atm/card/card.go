package card

import (
	"errors"
	"fmt"
)

type Card struct {
	NameOnCard string
	CardNumber string
	pin        string
}

func (c *Card) Authenticate(pin string) error {
	if c.pin == pin {
		fmt.Println("Successfully Authenticated")
		return nil
	}

	return errors.New("incorrect Pin")
}

func NewCard(cardNumber string, nameOnCard string, pin string) *Card {
	return &Card{
		CardNumber: cardNumber,
		NameOnCard: nameOnCard,
		pin:        pin,
	}
}

func (c *Card) SetPin(pin string) {
	c.pin = pin
}
