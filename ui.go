package main

import (
	"fmt"
	"net"
	"net/http"

	_ "embed"

	"github.com/go-vgo/robotgo"
	"github.com/haroflow/go-macros/automation"
	"github.com/haroflow/go-macros/automation/mouse"
	hook "github.com/robotn/gohook"
	"github.com/zserge/lorca"
)

var ui lorca.UI

// startUI shows www/index.html and blocks until the user closes the window.
func startUI() error {
	var err error

	ui, err = lorca.New("", "", 500, 600)
	if err != nil {
		return err
	}
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	defer ln.Close()

	http.HandleFunc("/help", helpHandler)

	http.Handle("/", http.FileServer(http.FS(fs)))

	go http.Serve(ln, nil)
	go listenHotkeys()

	initJavascriptVM()
	registerJavascriptFunctions()
	ui.Load(fmt.Sprintf("http://%s/www", ln.Addr()))

	<-ui.Done()

	return nil
}

func helpHandler(rw http.ResponseWriter, r *http.Request) {
	type TmplData struct {
		Commands []automation.Command
	}

	helpTemplate.Execute(rw, TmplData{
		Commands: registeredCommands,
	})
}

// listenHotkeys waits for key presses to determine what macro to run.
func listenHotkeys() {
	evChan := robotgo.EventStart()
	defer robotgo.EventEnd()

	for e := range evChan {
		// TODO customize key binding
		if e.Kind == hook.KeyUp && e.Keycode == hook.Keycode["f2"] {
			fmt.Println("f2 pressed")
			if isRunningMacro.Load().(bool) {
				fmt.Println("macro running... stopping macro")
				stopMacros()
			} else {
				fmt.Println("nothing running... starting macro")
				code := ui.Eval("getCode()").String()
				go executeCode(code)
			}
		}
	}
}

// registerJavascriptFunctions enables JS functions to call Go functions.
func registerJavascriptFunctions() {
	ui.Bind("executeCode", executeCode)
	ui.Bind("stopMacros", stopMacros)
	ui.Bind("getMousePosition", getMousePosition)
}

type MousePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func getMousePosition() MousePosition {
	return MousePosition{
		mouse.GetX(),
		mouse.GetY(),
	}
}
