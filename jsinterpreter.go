package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"sync/atomic"

	"github.com/dop251/goja"
	"github.com/go-vgo/robotgo"
	"github.com/haroflow/go-macros/automation"
	"github.com/haroflow/go-macros/automation/color"
	"github.com/haroflow/go-macros/automation/image"
	"github.com/haroflow/go-macros/automation/keyboard"
	"github.com/haroflow/go-macros/automation/mouse"
	"github.com/haroflow/go-macros/automation/other"
	"github.com/haroflow/go-macros/automation/screen"
)

var jsVM *goja.Runtime
var isRunningMacro atomic.Value
var registeredCommands []automation.Command

func initJavascriptVM() {
	jsVM = goja.New()
	isRunningMacro.Store(false)

	registerModuleCommands("mouse", mouse.Commands())
	registerModuleCommands("keyboard", keyboard.Commands())
	registerModuleCommands("screen", screen.Commands())
	registerModuleCommands("image", image.Commands())
	registerModuleCommands("color", color.Commands())
	registerGlobalCommands(other.Commands())

	logText := func(msg string) {
		cmd := fmt.Sprint("log('", msg, "');")
		ui.Eval(cmd)
	}

	logImage := func(img robotgo.CBitmap) {
		if img == nil {
			logText("logImage: image is null")
			return
		}
		buf := bytes.Buffer{}
		img2 := robotgo.ToImage(robotgo.ToMMBitmapRef(img))
		err := jpeg.Encode(&buf, img2, nil)
		if err != nil {
			logText(fmt.Sprint("cannot encode image to jpeg:", err))
			return
		}
		b64 := base64.StdEncoding.EncodeToString(buf.Bytes())

		cmd := fmt.Sprint("logImage('", b64, "');")
		ui.Eval(cmd)
	}

	registerModuleCommands("log", []automation.Command{
		{
			ModuleName:  "log",
			MethodName:  "text",
			Parameters:  "msg: string",
			Description: "Prints the message in the log window",
			Action:      logText,
		},
		{
			ModuleName:  "log",
			MethodName:  "image",
			Parameters:  "img: image",
			Description: "Prints the image in the log window",
			Action:      logImage,
		},
	})
}

func registerModuleCommands(moduleName string, commands []automation.Command) {
	m := map[string]interface{}{}

	for _, cmd := range commands {
		m[cmd.MethodName] = cmd.Action
		registeredCommands = append(registeredCommands, cmd)
	}

	jsVM.Set(moduleName, m)
}

func registerGlobalCommands(commands []automation.Command) {
	for _, cmd := range commands {
		registeredCommands = append(registeredCommands, cmd)
		jsVM.Set(cmd.MethodName, cmd.Action)
	}
}

func executeCode(code string) error {
	fmt.Println("running code...")

	isRunningMacro.Store(true)
	v, err := jsVM.RunString(code)
	isRunningMacro.Store(false)
	if err != nil {
		if ex, ok := err.(*goja.Exception); ok {
			fmt.Println("error:", ex.String())
			return fmt.Errorf(ex.String())
		}
		fmt.Println("error:", err)
		return err
	}
	fmt.Println("success:", v)

	return nil
}

func stopMacros() {
	fmt.Println("stopping")
	isRunningMacro.Store(false)
	jsVM.Interrupt("halt")
	fmt.Println("stopped")
}
