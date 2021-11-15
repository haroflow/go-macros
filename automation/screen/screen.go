package screen

import (
	"image/color"

	"github.com/go-vgo/robotgo"
	"github.com/haroflow/go-macros/automation"
)

func Commands() []automation.Command {
	moduleName := "screen"

	return []automation.Command{
		{
			ModuleName:  moduleName,
			MethodName:  "capture",
			Parameters:  "",
			Description: "Captures and returns a screenshot of the entire screen.",
			Action:      Capture,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "captureRect",
			Parameters:  "x: int, y: int, width: int, height: int",
			Description: "Captures and returns a screenshot of a rectangular area on the screen.",
			Action:      CaptureRect,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "getPixelColor",
			Parameters:  "x: int, y: int",
			Description: "Returns the color from a point on screen.",
			Action:      GetPixelColor,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "getWidth",
			Parameters:  "",
			Description: "Returns the screen width.",
			Action:      GetWidth,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "getHeight",
			Parameters:  "",
			Description: "Returns the screen height.",
			Action:      GetHeight,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "search",
			Parameters:  "img: image, tolerance: float from 0.0 to 1.0",
			Description: "Search for an image on the screen. The tolerance parameter should be in the range 0.0 to 1.0, denoting how closely the colors in the image have to match, with 0.0 being exact and 1.0 being any. Returns a point.",
			Action:      Search,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "searchAll",
			Parameters:  "",
			Description: "Search for all instances of an image on the screen. The tolerance parameter should be in the range 0.0 to 1.0, denoting how closely the colors in the image have to match, with 0.0 being exact and 1.0 being any. Returns a list of points.",
			Action:      SearchAll,
		},
	}
}

func Capture() robotgo.CBitmap {
	// Unfortunately, robotgo.FindBitmap is not working
	// when converting an image.Image to a MMBitmapRef. I don't know why...
	// So for now we'll use robotgo.CBitmap and robotgo.CaptureScreen.
	return robotgo.CBitmap(robotgo.CaptureScreen())
}

func CaptureRect(x, y, w, h int) robotgo.CBitmap {
	// Unfortunately, robotgo.FindBitmap is not working
	// when converting an image.Image to a MMBitmapRef. I don't know why...
	// So for now we'll use robotgo.CBitmap and robotgo.CaptureScreen.
	if w == 0 || h == 0 {
		return robotgo.CBitmap(robotgo.CaptureScreen())
	}
	return robotgo.CBitmap(robotgo.CaptureScreen(x, y, w, h))
}

func GetPixelColor(x, y int) color.Color {
	c := robotgo.GetPxColor(x, y)

	return color.RGBA{
		B: uint8(uint(c)),
		G: uint8(uint(c) >> 8),
		R: uint8(uint(c) >> 16),
		A: 0,
	}
}

func GetWidth() int {
	x, _ := robotgo.GetScreenSize()
	return x
}

func GetHeight() int {
	_, y := robotgo.GetScreenSize()
	return y
}

func Search(img robotgo.CBitmap, tolerance float64) Point {
	// Unfortunately, robotgo.FindBitmap is not working
	// when converting an image.Image to a MMBitmapRef. I don't know why...
	// So for now we'll use robotgo.CBitmap and robotgo.CaptureScreen.
	c := robotgo.ToMMBitmapRef(img)

	screen := robotgo.CaptureScreen()
	x, y := robotgo.FindBitmap(c, screen, tolerance)
	return Point{X: x, Y: y}
}

func SearchAll(img robotgo.CBitmap, tolerance float64) []Point {
	// Unfortunately, robotgo.FindBitmap is not working
	// when converting an image.Image to a MMBitmapRef. I don't know why...
	// So for now we'll use robotgo.CBitmap and robotgo.CaptureScreen.
	c := robotgo.ToMMBitmapRef(img)

	screen := robotgo.CaptureScreen()
	points := robotgo.FindEveryBitmap(c, screen, tolerance)

	ret := make([]Point, len(points))
	for i, p := range points {
		ret[i] = Point{X: p.X, Y: p.Y}
	}

	return ret
}
