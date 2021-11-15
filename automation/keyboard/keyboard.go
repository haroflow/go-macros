package keyboard

import (
	"github.com/go-vgo/robotgo"
	"github.com/haroflow/go-macros/automation"
)

func Commands() []automation.Command {
	moduleName := "keyboard"
	return []automation.Command{
		{
			ModuleName:  moduleName,
			MethodName:  "type",
			Parameters:  "message: string",
			Description: "Types a message on the keyboard.",
			Action:      Type,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "press",
			Parameters:  "key: string, ...",
			Description: "Presses one or more keys simultaneously. For a complete list of available keys, see: https://github.com/go-vgo/robotgo/blob/master/docs/keys.md",
			Action:      Press,
		},
	}
}

func Type(msg string) {
	robotgo.TypeStr(msg)
}

func Press(key string, other ...string) {
	//https://github.com/go-vgo/robotgo/blob/master/docs/keys.md
	if len(other) > 0 {
		robotgo.KeyTap(key, other)
	} else {
		robotgo.KeyTap(key)
	}
}
