package util

type Vector2 struct {
	X float32
	Y float32
}

func (v *Vector2) Add(b *Vector2) *Vector2 {
	v.X += b.X
	v.Y += b.Y
	return v
}

func (v *Vector2) Sub(b *Vector2) *Vector2 {
	v.X -= b.X
	v.Y -= b.Y
	return v
}

func (v *Vector2) Mul(b *Vector2) *Vector2 {
	v.X *= b.X
	v.Y *= b.Y
	return v
}

func (v *Vector2) Div(b *Vector2) *Vector2 {
	v.X /= b.X
	v.Y /= b.Y
	return v
}

func (v *Vector2) SquareLen() float32 {
	return v.X*v.X + v.Y*v.Y
}

func NewVector2Zero() *Vector2 {
	return &Vector2{
		X: 0,
		Y: 0,
	}
}
