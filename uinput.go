package main

import "github.com/bendahl/uinput"

// alternatively (to use specific version), use this:
// import "gopkg.in/bendahl/uinput.v1"

func main() {
	// initialize keyboard and check for possible errors
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("testkeyboard"))
	if err != nil {
		return
	}
	// always do this after the initialization in order to guarantee that the device will be properly closed
	defer keyboard.Close()

	// prints "a"
	keyboard.KeyPress(uinput.KeyA)

	// CtrlC
	// keyboard.KeyDown(uinput.KeyLeftctrl)
	// keyboard.KeyPress(uinput.KeyC)
	// keyboard.KeyUp(uinput.KeyLeftctrl)

	keyboard.KeyPress(uinput.KeyRight)
	keyboard.KeyPress(uinput.KeySpace)
	keyboard.KeyPress(uinput.KeyMinus)

	// keyboard.KeyDown(uinput.KeyLeftctrl)
	// keyboard.KeyPress(uinput.KeyV)
	// keyboard.KeyUp(uinput.KeyLeftctrl)

	// prints "A"
	// Note that you could use caps lock instead of using shift with KeyDown and KeyUp
	keyboard.KeyDown(uinput.KeyLeftshift)
	keyboard.KeyPress(uinput.KeyA)
	keyboard.KeyUp(uinput.KeyLeftshift)

	// prints "00000"
	for i := 0; i < 5; i++ {
		keyboard.KeyPress(uinput.Key0)
	}
}
