package artihmetic

// x, y 2개의 Integer 구조체
type Numbers struct {
	X int
	Y int
}

// Plus x, y 합을 계산해서 반환
func (o *Numbers) Plus() int {
	return o.X + o.Y
}

// x, y 차를 계산해서 반환
func (o *Numbers) Minus() int {
	return o.X - o.Y
}

// x, y 곱을 계산해서 반환
func (o *Numbers) Multiply() int {
	return o.X * o.Y
}

// x, y 나누기를 계산해서 반환
func (o *Numbers) Divide() int {
	return o.X / o.Y
}
