package shape

type Rectangle struct{
	Widht float32
	Lenght float32
}

func (rectangle *Rectangle) Area() float32 {
	return rectangle.Widht * rectangle.Lenght
}

func (rectangle *Rectangle) Perimeter() float32 {
	return 2 * (rectangle.Widht + rectangle.Lenght)
}