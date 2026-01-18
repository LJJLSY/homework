package main

import (
	"fmt"
	"math"
)

// 定义Shape接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 实现Shape接口
type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return (r.length + r.width) * 2
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	R := Rectangle{5, 3}
	C := Circle{5}
	fmt.Printf("矩形的长：%f，宽：%f，面积是：%f，周长是：%f\n", R.length, R.width, R.Area(), R.Perimeter())
	fmt.Printf("圆的半径：%f，面积是：%f，周长是：%f\n", C.radius, C.Area(), C.Perimeter())
}
