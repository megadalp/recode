package main

import (
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func main() {
	go func() {
		//
		// keyboard.SimulateKeyPress("Hello")             // Simulate key press for every letter in string
		// keyboard.SimulateKeyPress(keys.Enter)          // Simulate key press for Enter
		// keyboard.SimulateKeyPress(keys.CtrlShiftRight) // Simulate key press for Ctrl+Shift+Right
		// keyboard.SimulateKeyPress('x')                 // Simulate key press for a single rune
		// keyboard.SimulateKeyPress('x', keys.Down, 'a') // Simulate key presses for multiple inputs
		//
		keyboard.SimulateKeyPress(keys.CtrlC)
		keyboard.SimulateKeyPress(keys.CtrlV)
		//
		// keyboard.SimulateKeyPress(keys.Escape) // Simulate key press for Escape, which quits the program
	}()

	// keyboard.Listen(func(key keys.Key) (stop bool, err error) {
	//     // if key.Code == keys.Escape || key.Code == keys.CtrlC {
	//     if key.Code == keys.Escape {
	//         os.Exit(0) // Exit program on Escape
	//     }
	//
	//     fmt.Println("\r" + key.String()) // Print every key press
	//     return false, nil                // Return false to continue listening
	// })
}
