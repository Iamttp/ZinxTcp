package util

import (
	"log"
	"testing"
)

func TestVector2(t *testing.T) {
	a := Vector2{
		X: 1.3,
		Y: 1.6,
	}
	b := Vector2{
		X: 1.9,
		Y: 3.5,
	}
	log.Println(a.Add(&b))
	log.Println(b.SquareLen())
}
