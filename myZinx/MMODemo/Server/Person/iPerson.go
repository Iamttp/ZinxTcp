package Person

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
