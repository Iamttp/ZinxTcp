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
func (p *Person) setState(state State)             { p.state = state }
func (p *Person) getState() State                  { return p.state }
func (p *Person) setTypePerson(person typePerson)  { p.person = person }
func (p *Person) getTypePerson() typePerson        { return p.person }
func (p *Person) setOwner(owner bool)              { p.owner = owner }
func (p *Person) getOwner() bool                   { return p.owner }
func (p *Person) setSpeedVal(speedVal float32)     { p.speedVal = speedVal }
func (p *Person) getSpeedVal() float32             { return p.speedVal }
func (p *Person) setMoveVec(moveVec *util.Vector2) { p.moveVec = moveVec }
func (p *Person) getMoveVec() *util.Vector2        { return p.moveVec }
