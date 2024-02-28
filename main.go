package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/hotkey"
	"os/exec"
	"runtime"
)

var hk *hotkey.Hotkey
var selectedKey = "F9"

func main() {
	a := app.New()
	w := a.NewWindow("Turn Off Monitor")
	w.Resize(fyne.NewSize(300, 260))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	label := widget.NewLabel("To turn off your monitor press the button!")

	button := widget.NewButton("Turn Off Monitor", func() {
		err := turnOffMonitor()
		if err != nil {
			fyneErrorHandler(err, a, w)
		}
	})
	exitButton := widget.NewButton("Quit", func() {
		a.Quit()
	})
	settingsButton := widget.NewButton("Change Key", func() {
		appSettings(a, w)
	})

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	w.SetContent(container.NewVBox(label, button, settingsButton, exitButton))
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	go listenHotkey(hotkey.KeyF9)

	w.ShowAndRun()
}
func turnOffMonitor() error {
	switch runtime.GOOS {
	case "windows":
		// Windows: PowerShell command to turn off the monitor
		cmd := exec.Command("powershell", `(Add-Type "[DllImport(""user32.dll"")] public static extern int PostMessage(int hWnd, int hMsg, int wParam, int lParam);" -Name "Win32PostMessage" -Namespace Win32Functions -PassThru)::PostMessage(0xffff, 0x0112, 0xF170, 2)`)
		return cmd.Run()
	case "linux":
		// Linux: Using xset to turn off the monitor
		cmd := exec.Command("xset", "dpms", "force", "off")
		return cmd.Run()
	case "darwin":
		// macOS: Using pmset to put the display to sleep
		cmd := exec.Command("pmset", "displaysleepnow")
		return cmd.Run()
	default:
		return errors.New("OS:" + runtime.GOOS + " is not yet supported")
	}
}

func appSettings(a fyne.App, w fyne.Window) {
	w = a.NewWindow("Settings")
	w.CenterOnScreen()
	w.SetFixedSize(true)
	labelSettings := widget.NewLabel("Press the key you want to use to turn off the monitor:")
	currentKey := widget.NewLabel("Current Key:" + selectedKey)
	w.SetContent(container.NewVBox(labelSettings, container.NewCenter(currentKey)))
	w.Show()

	// add a listener to the key press
	w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		var key hotkey.Key

		switch e.Name {
		case "F1":
			key = hotkey.KeyF1
		case "F2":
			key = hotkey.KeyF2
		case "F3":
			key = hotkey.KeyF3
		case "F4":
			key = hotkey.KeyF4
		case "F5":
			key = hotkey.KeyF5
		case "F6":
			key = hotkey.KeyF6
		case "F7":
			key = hotkey.KeyF7
		case "F8":
			key = hotkey.KeyF8
		case "F9":
			key = hotkey.KeyF9
		case "F10":
			key = hotkey.KeyF10
		case "F11":
			key = hotkey.KeyF11
		case "0":
			key = hotkey.Key0
		case "1":
			key = hotkey.Key1
		case "2":
			key = hotkey.Key2
		case "3":
			key = hotkey.Key3
		case "4":
			key = hotkey.Key4
		case "5":
			key = hotkey.Key5
		case "6":
			key = hotkey.Key6
		case "7":
			key = hotkey.Key7
		case "8":
			key = hotkey.Key8
		case "9":
			key = hotkey.Key9
		case "A":
			key = hotkey.KeyA
		case "B":
			key = hotkey.KeyB
		case "C":
			key = hotkey.KeyC
		case "D":
			key = hotkey.KeyD
		case "E":
			key = hotkey.KeyE
		case "F":
			key = hotkey.KeyF
		case "G":
			key = hotkey.KeyG
		case "H":
			key = hotkey.KeyH
		case "I":
			key = hotkey.KeyI
		case "J":
			key = hotkey.KeyJ
		case "K":
			key = hotkey.KeyK
		case "L":
			key = hotkey.KeyL
		case "M":
			key = hotkey.KeyM
		case "N":
			key = hotkey.KeyN
		case "O":
			key = hotkey.KeyO
		case "P":
			key = hotkey.KeyP
		case "Q":
			key = hotkey.KeyQ
		case "R":
			key = hotkey.KeyR
		case "S":
			key = hotkey.KeyS
		case "T":
			key = hotkey.KeyT
		case "U":
			key = hotkey.KeyU
		case "V":
			key = hotkey.KeyV
		case "W":
			key = hotkey.KeyW
		case "X":
			key = hotkey.KeyX
		case "Y":
			key = hotkey.KeyY
		case "Z":
			key = hotkey.KeyZ
		default:
			return
		}
		selectedKey = string(e.Name)
		go func() {
			err := listenHotkey(key)
			if err != nil {
				fyneErrorHandler(err, a, w)
			}
		}()
		w.Close()
	})
}
func listenHotkey(key hotkey.Key) (err error) {
	if hk != nil {
		hk.Unregister() // Unregister the previous hotkey
	}
	hk = hotkey.New(nil, key)
	if err = hk.Register(); err != nil {
		return
	}
	// Start listen hotkey event
	for range hk.Keydown() {
		err = turnOffMonitor()
		if err != nil {
			return
		}
	}
	return
}

func fyneErrorHandler(err error, a fyne.App, w fyne.Window) {
	w = a.NewWindow("error")
	w.SetFixedSize(true)
	button := widget.NewButton("exit", func() {
		a.Quit()
	})
	w.CenterOnScreen()
	w.SetContent(container.NewVBox(widget.NewLabel(err.Error()), button))
	w.Show()
}
