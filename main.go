package main

import (
	"github.com/nsf/termbox-go"
	"math"
	"time"
)

var keyboardEventsChan = make(chan keyboardEventType)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	dotColor     = termbox.ColorYellow

	Pi2 = math.Pi * 2.0
)

var (
	phi      = 0.0
	deltaPhi = 0.01
	draw     = '*'
	a        = 1.0
	b        = 1.0
)

func move() {
	phi += deltaPhi
	if phi >= Pi2 {
		phi -= Pi2
	}
}

func render() error {
	if err := termbox.Clear(defaultColor, defaultColor); err != nil {
		panic("rrr")
	}

	w, h := termbox.Size()

	h2 := float64(h) / 2.0
	w2 := float64(w) / 2.0

	for y := 0; y < h; y++ {
		termbox.SetCell(int(w2), y, '|', termbox.ColorGreen, bgColor)
	}

	for x := 0; x < w; x++ {
		termbox.SetCell(x, int(h2), '-', termbox.ColorGreen, bgColor)
	}

	termbox.SetCell(int(w2), int(h2), '+', termbox.ColorGreen, bgColor)

	for omega := 0.0; omega < Pi2; omega += 0.01 {
		y := h2 + math.Sin(a*omega+phi)*h2
		x := w2 + math.Sin(b*omega)*w2
		termbox.SetCell(int(x), int(y), draw, dotColor, bgColor)
	}

	return termbox.Flush()
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKeyboard(keyboardEventsChan)

	if err := render(); err != nil {
		panic(err)
	}

mainloop:
	for {
		select {
		case e := <-keyboardEventsChan:
			switch e {
			case UseStars:
				draw = '*'
			case UseDots:
				draw = '.'
			case IncA:
				a += 1.0
			case DecA:
				a = math.Max(1.0, a-1.0)
			case IncB:
				b += 1.0
			case DecB:
				b = math.Max(1.0, b-1.0)
			case Exit:
				break mainloop
			}
		default:
			move()
			if err := render(); err != nil {
				panic(err)
			}

			time.Sleep(20 * time.Millisecond)
		}
	}
}
