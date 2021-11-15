package image

import (
	"github.com/go-vgo/robotgo"
	"github.com/haroflow/go-macros/automation"
)

func Commands() []automation.Command {
	moduleName := "image"
	return []automation.Command{
		{
			ModuleName:  moduleName,
			MethodName:  "saveJpeg",
			Parameters:  "img: image, path: string",
			Description: "Saves the image to disk on the specified path as a JPEG.",
			Action:      SaveJpeg,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "savePng",
			Parameters:  "img: image, path: string",
			Description: "Saves the image to disk on the specified path as a PNG.",
			Action:      SavePng,
		},
	}
}

func SaveJpeg(img robotgo.CBitmap, path string) error {
	return robotgo.SaveJpeg(robotgo.ToImage(robotgo.ToMMBitmapRef(img)), path)
}

func SavePng(img robotgo.CBitmap, path string) error {
	return robotgo.SavePng(robotgo.ToImage(robotgo.ToMMBitmapRef(img)), path)
}

// TODO: Adicionar funções de ajuste de HSL, contraste, grayscale, sharpen, etc.
