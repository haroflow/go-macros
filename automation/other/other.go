package other

import (
	"time"

	"github.com/haroflow/go-macros/automation"
)

func Commands() []automation.Command {
	return []automation.Command{
		{
			ModuleName:  "",
			MethodName:  "sleep",
			Parameters:  "milliseconds: int",
			Description: "Waits for the specified milliseconds",
			Action:      Sleep,
		},
	}
}

func Sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
