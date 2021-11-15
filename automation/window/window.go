package window

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
)

func Show(windowName string) {
	// ids, err := robotgo.FindIds(windowName)
	// if err != nil {
	// 	panic(err) // FIXME
	// }
	// fmt.Println(ids)

	// robotgo.CloseWindow(ids[0]) // ok
	// robotgo.MaxWindow(ids[0])// ok

	h := robotgo.FindWindow(windowName)
	fmt.Println(h)
	win.BringWindowToTop(h)
	win.SetForegroundWindow(h)
	win.SetFocus(h)
	win.EnableWindow(h, true)
}
