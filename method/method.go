package method

import "fmt"

type Point struct {
	X, Y int
}

type Center struct {
	X, Y int
}
type Cicle struct {
	*Center
	Point
}

func (p *Point) Distance(q Point) int {
	return q.X - p.X
}

func (p *Point) Distance2(q Point) int {
	return q.X - (*p).X
}

func (c *Cicle) Length(p Point) int {
	return p.X - c.Center.X
}

func (c *Cicle) ReturnX() int {
	fmt.Println(c.Center.X)
	return c.Center.X
}

func (c *Cicle) ReturnX2() int {
	fmt.Println((*c.Center).X)
	return (*c.Center).X
}
