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
	setState(state State)
	getState() State

	setTypePerson(person typePerson)
	getTypePerson() typePerson

	setOwner(owner bool)
	getOwner() bool

	setSpeedVal(speedVal float32)
	getSpeedVal() float32

	setMoveVec(moveVec *util.Vector2)
	getMoveVec() *util.Vector2
}
