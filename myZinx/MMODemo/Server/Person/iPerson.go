package Person

import (
	"awesomeProject/myZinx/MMODemo/Server/util"
)

type State uint8

const (
	idle   State = 0
	run    State = 1
	dead   State = 2
	attack State = 3
	hit    State = 4
)

type typePerson uint8

const (
	Sword  typePerson = 0
	Anchor typePerson = 1
)

type IPerson interface {
	move()
	attack()
	hit(val int)
	dead()

	// getter setter
	SetState(state State)
	GetState() State

	SetTypePerson(person typePerson)
	GetTypePerson() typePerson

	SetOwner(owner bool)
	GetOwner() bool

	SetSpeedVal(speedVal float32)
	GetSpeedVal() float32

	SetMoveVec(moveVec *util.Vector2)
	GetMoveVec() *util.Vector2
}
