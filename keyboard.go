package main

import "github.com/nsf/termbox-go"

type keyboardEventType int

const (
	// Exit stops the program
	Exit keyboardEventType = 1 + iota
	// UseDots switches the "graphics" to use dots
	UseDots
	// UseStars switches the "graphics" to use stars (asterisks)
	UseStars
	// IncA increases the frequency on the x axis
	IncA
	// DecA decreases the frequency on the x axis
	DecA
	// IncB increases the frequency on the y axis
	IncB
	// DecB decreases the frequency on the y axis
	DecB
)

func listenToKeyboard(eventChan chan keyboardEventType) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				eventChan <- Exit
			default:
				switch ev.Ch {
				case '*':
					eventChan <- UseStars
				case '.':
					eventChan <- UseDots
				case 'q', 'Q':
					eventChan <- Exit
				case 'x':
					eventChan <- DecA
				case 'X':
					eventChan <- IncA
				case 'y':
					eventChan <- DecB
				case 'Y':
					eventChan <- IncB
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
