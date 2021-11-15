package screen

import "fmt"

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p Point) String() string {
	return fmt.Sprint("x: ", p.X, " y: ", p.Y)
}
