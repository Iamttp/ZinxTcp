package Person

import "awesomeProject/myZinx/MMODemo/Server/util"

type Person struct {
	state    State
	person   typePerson
	owner    bool
	speedVal float32
	moveVec  *util.Vector2
}

func (p *Person) move()       {}
func (p *Person) attack()     {}
func (p *Person) hit(val int) {}
func (p *Person) dead()       {}

// getter setter
func (p *Person) SetState(state State)             { p.state = state }
func (p *Person) GetState() State                  { return p.state }
func (p *Person) SetTypePerson(person typePerson)  { p.person = person }
func (p *Person) GetTypePerson() typePerson        { return p.person }
func (p *Person) SetOwner(owner bool)              { p.owner = owner }
func (p *Person) GetOwner() bool                   { return p.owner }
func (p *Person) SetSpeedVal(speedVal float32)     { p.speedVal = speedVal }
func (p *Person) GetSpeedVal() float32             { return p.speedVal }
func (p *Person) SetMoveVec(moveVec *util.Vector2) { p.moveVec = moveVec }
func (p *Person) GetMoveVec() *util.Vector2        { return p.moveVec }
