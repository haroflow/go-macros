package color

import (
	gocolor "image/color"
	"math"

	"github.com/haroflow/go-macros/automation"
)

func Commands() []automation.Command {
	moduleName := "color"
	return []automation.Command{
		{
			ModuleName:  moduleName,
			MethodName:  "getBrightness",
			Parameters:  "color: color",
			Description: "Returns the brightness of a color, in the range 0 to 255.",
			Action:      GetBrightness,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "distance",
			Parameters:  "c1: color, c2: color",
			Description: "Returns the distance between the colors, in the range 0 (equal) to (255^2)*3 (difference between white and black)",
			Action:      Distance,
		},
	}
}

func GetBrightness(c gocolor.RGBA) uint8 {
	return uint8((uint(c.R) + uint(c.G) + uint(c.B)) / 3)
}

func Distance(a gocolor.RGBA, b gocolor.RGBA) float64 {
	return math.Pow(float64(a.R)-float64(b.R), 2.0) +
		math.Pow(float64(a.G)-float64(b.G), 2.0) +
		math.Pow(float64(a.B)-float64(b.B), 2.0)
}

// TODO: tohexstring
