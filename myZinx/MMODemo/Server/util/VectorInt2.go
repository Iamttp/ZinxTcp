package util

type Vector2Int struct {
	X int32
	Y int32
}

func (v *Vector2Int) Add(b *Vector2Int) *Vector2Int {
	v.X += b.X
	v.Y += b.Y
	return v
}

func (v *Vector2Int) Sub(b *Vector2Int) *Vector2Int {
	v.X -= b.X
	v.Y -= b.Y
	return v
}

func (v *Vector2Int) Mul(b *Vector2Int) *Vector2Int {
	v.X *= b.X
	v.Y *= b.Y
	return v
}

func (v *Vector2Int) Div(b *Vector2Int) *Vector2Int {
	v.X /= b.X
	v.Y /= b.Y
	return v
}

func (v *Vector2Int) SquareLen() int32 {
	return v.X*v.X + v.Y*v.Y
}

func NewVector2IntZero() *Vector2Int {
	return &Vector2Int{
		X: 0,
		Y: 0,
	}
}
