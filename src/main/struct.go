package main

import "fmt"

type Point struct {
	X, Y int
}

type Node struct {
	point *Point
	next  *Node
}

func main() {
	NoStructDemo()
	fmt.Println("-----------")
	StructDemo()
	fmt.Println("-----------")
	StructDemo2()
	fmt.Println("-----------")
	StructDemo3()
	fmt.Println("-----------")
	StructDemo4()
	fmt.Println("-----------")
	StructDemo5()
	fmt.Println("-----------")
	StructDemo6()
	fmt.Println("-----------")
	StructDemo7()
	fmt.Println("-----------")
	StructDemo8()
	fmt.Println("-----------")
	StructDemo9()
}

func NoStructDemo() {
	x := 10
	y := 20
	fmt.Printf("{%d %d}\n", x, y)

	x = 20
	y = 30
	fmt.Printf("{%d %d}\n", x, y)
}

func StructDemo() {
	point := struct {
		x, y int
	}{10, 20}
	fmt.Printf("{%d %d}\n", point.x, point.y)

	point.x = 20
	point.y = 30

	fmt.Printf("{%d %d}\n", point.x, point.y)

	fmt.Println(point)
}

func StructDemo2() {
	point1 := Point{10, 20}
	fmt.Println(point1)

	point2 := Point{Y: 20, X: 30}
	fmt.Println(point2)
}

func StructDemo3() {
	var point Point
	fmt.Println(point)
}

func StructDemo4() {
	point1 := Point{10, 20}
	point2 := point1

	point1.X = 20

	fmt.Println(point1)
	fmt.Println(point2)
}

func StructDemo5() {
	point := Point{10, 20}

	changeX(point)
	fmt.Println(point)
}

func changeX(point Point) {
	point.X = 20
	fmt.Println(point)
}

func StructDemo6() {
	point1 := Point{10, 20}
	point2 := &point1

	point1.X = 20

	fmt.Println(point1)
	fmt.Println(point2)
}

func StructDemo7() {
	point := Point{10, 20}
	changeX2(&point)

	fmt.Println(point)
}

func changeX2(point *Point) {
	point.X = 20
	fmt.Printf("&{%d %d}\n", point.X, point.Y)
}

func StructDemo8() {
	point := default_point()
	fmt.Println(point)
}

func default_point() *Point {
	point := new(Point)
	point.X = 10
	point.Y = 10
	return point
}

func StructDemo9() {
	node := new(Node)

	node.point = &Point{10, 20}
	node.next = new(Node)

	node.next.point = &Point{10, 30}

	fmt.Println(node.point)
	fmt.Println(node.next.point)
}
