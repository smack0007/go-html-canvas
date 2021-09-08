package game

import (
	"types"
)

type CardSuit byte

const (
	Clubs CardSuit = iota
	Diamonds
	Hearts
	Spades
)

type CardValue byte

const (
	Ace CardValue = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit CardSuit
	Value CardValue
}

var CardImageRectangles = map[Card]types.Rect{
	{Suit: Clubs, Value: Ten}: { X: 560, Y: 760, Width: 140, Height: 190},
	{Suit: Clubs, Value: Two}: { X: 280, Y: 1140, Width: 140, Height: 190},
	{Suit: Clubs, Value: Three}: { X: 700, Y: 190, Width: 140, Height: 190},
	{Suit: Clubs, Value: Four}: { X: 700, Y: 0, Width: 140, Height: 190},
	{Suit: Clubs, Value: Five}: { X: 560, Y: 1710, Width: 140, Height: 190},
	{Suit: Clubs, Value: Six}: { X: 560, Y: 1520, Width: 140, Height: 190},
	{Suit: Clubs, Value: Seven}: { X: 560, Y: 1330, Width: 140, Height: 190},
	{Suit: Clubs, Value: Eight}: { X: 560, Y: 1140, Width: 140, Height: 190},
	{Suit: Clubs, Value: Nine}: { X: 560, Y: 950, Width: 140, Height: 190},
	{Suit: Clubs, Value: Ace}: { X: 560, Y: 570, Width: 140, Height: 190},
	{Suit: Clubs, Value: Jack}: { X: 560, Y: 380, Width: 140, Height: 190},
	{Suit: Clubs, Value: King}: { X: 560, Y: 190, Width: 140, Height: 190},
	{Suit: Clubs, Value: Queen}: { X: 560, Y: 0, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Ten}: { X: 420, Y: 190, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Two}: { X: 420, Y: 1710, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Three}: { X: 420, Y: 1520, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Four}: { X: 420, Y: 1330, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Five}: { X: 420, Y: 1140, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Six}: { X: 420, Y: 950, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Seven}: { X: 420, Y: 760, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Eight}: { X: 420, Y: 570, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Nine}: { X: 420, Y: 380, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Ace}: { X: 420, Y: 0, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Jack}: { X: 280, Y: 1710, Width: 140, Height: 190},
	{Suit: Diamonds, Value: King}: { X: 280, Y: 1520, Width: 140, Height: 190},
	{Suit: Diamonds, Value: Queen}: { X: 280, Y: 1330, Width: 140, Height: 190},
	{Suit: Hearts, Value: Ten}: { X: 140, Y: 1520, Width: 140, Height: 190},
	{Suit: Hearts, Value: Two}: { X: 700, Y: 380, Width: 140, Height: 190},
	{Suit: Hearts, Value: Three}: { X: 280, Y: 950, Width: 140, Height: 190},
	{Suit: Hearts, Value: Four}: { X: 280, Y: 760, Width: 140, Height: 190},
	{Suit: Hearts, Value: Five}: { X: 280, Y: 570, Width: 140, Height: 190},
	{Suit: Hearts, Value: Six}: { X: 280, Y: 380, Width: 140, Height: 190},
	{Suit: Hearts, Value: Seven}: { X: 280, Y: 190, Width: 140, Height: 190},
	{Suit: Hearts, Value: Eight}: { X: 280, Y: 0, Width: 140, Height: 190},
	{Suit: Hearts, Value: Nine}: { X: 140, Y: 1710, Width: 140, Height: 190},
	{Suit: Hearts, Value: Ace}: { X: 140, Y: 1330, Width: 140, Height: 190},
	{Suit: Hearts, Value: Jack}: { X: 140, Y: 1140, Width: 140, Height: 190},
	{Suit: Hearts, Value: King}: { X: 140, Y: 950, Width: 140, Height: 190},
	{Suit: Hearts, Value: Queen}: { X: 140, Y: 760, Width: 140, Height: 190},
	{Suit: Spades, Value: Ten}: { X: 0, Y: 60, Width: 140, Height: 190},
	{Suit: Spades, Value: Two}: { X: 140, Y: 380, Width: 140, Height: 190},
	{Suit: Spades, Value: Three}: { X: 140, Y: 190, Width: 140, Height: 190},
	{Suit: Spades, Value: Four}: { X: 140, Y: 0, Width: 140, Height: 190},
	{Suit: Spades, Value: Five}: { X: 0,Y: 710, Width: 140, Height: 190},
	{Suit: Spades, Value: Six}: { X: 0,Y: 520, Width: 140, Height: 190},
	{Suit: Spades, Value: Seven}: { X: 0,Y: 330, Width: 140, Height: 190},
	{Suit: Spades, Value: Eight}: { X: 0,Y: 140, Width: 140, Height: 190},
	{Suit: Spades, Value: Nine}: { X: 0,Y: 50, Width: 140, Height: 190},
	{Suit: Spades, Value: Ace}: { X: 0,Y: 70, Width: 140, Height: 190},
	{Suit: Spades, Value: Jack}: { X: 0,Y: 80, Width: 140, Height: 190},
	{Suit: Spades, Value: King}: { X: 0,Y: 90, Width: 140, Height: 190},
	{Suit: Spades, Value: Queen}: { X: 0, Y: 0, Width: 140, Height: 190},
}
