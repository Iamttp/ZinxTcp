package Person

import "awesomeProject/myZinx/MMODemo/Server/util"

type Person struct {
	state    State
	person   TypePerson
	owner    bool
	speedVal float32
	moveVec  *util.Vector2
}

func (p *Person) Move()       {}
func (p *Person) Attack()     {}
func (p *Person) Hit(val int) {}
func (p *Person) Dead()       {}

// getter setter
func (p *Person) SetState(state State)             { p.state = state }
func (p *Person) GetState() State                  { return p.state }
func (p *Person) SetTypePerson(person TypePerson)  { p.person = person }
func (p *Person) GetTypePerson() TypePerson        { return p.person }
func (p *Person) SetOwner(owner bool)              { p.owner = owner }
func (p *Person) GetOwner() bool                   { return p.owner }
func (p *Person) SetSpeedVal(speedVal float32)     { p.speedVal = speedVal }
func (p *Person) GetSpeedVal() float32             { return p.speedVal }
func (p *Person) SetMoveVec(moveVec *util.Vector2) { p.moveVec = moveVec }
func (p *Person) GetMoveVec() *util.Vector2        { return p.moveVec }

// 一定自己实现NewXXX 充当构造函数，防止空指针错误
func NewPerson() IPerson {
	return &Person{
		state:    0,
		person:   0,
		owner:    false,
		speedVal: 0,
		moveVec:  util.NewVector2Zero(),
	}
}
