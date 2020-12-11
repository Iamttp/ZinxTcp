package Person

import (
	"awesomeProject/myZinx/MMODemo/Server/util"
)

type State uint8

const (
	Idle   State = 0
	Run    State = 1
	Dead   State = 2
	Attack State = 3
	Hit    State = 4
)

type TypePerson uint8

const (
	Sword  TypePerson = 0
	Anchor TypePerson = 1
)

type IPerson interface {
	Move()
	Attack()
	Hit(val int)
	Dead()

	// getter setter
	SetState(state State)
	GetState() State

	SetTypePerson(person TypePerson)
	GetTypePerson() TypePerson

	SetOwner(owner bool)
	GetOwner() bool

	SetSpeedVal(speedVal float32)
	GetSpeedVal() float32

	SetMoveVec(moveVec *util.Vector2)
	GetMoveVec() *util.Vector2
}
