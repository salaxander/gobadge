package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/tinygo-org/gobadge/fonts"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

const (
	WIDTH  = 160
	HEIGHT = 128
)

const (
	BLACK = iota
	WHITE
	RED
)

const (
	logoDisplayTime = 10 * time.Second
)

var colors = []color.RGBA{
	color.RGBA{0, 0, 0, 255},
	color.RGBA{255, 255, 255, 255},
	color.RGBA{255, 0, 0, 255},
}

var rainbow []color.RGBA
var pressed uint8
var quit bool

func Badge() {
	quit = false
	display.FillScreen(colors[BLACK])

	rainbow = make([]color.RGBA, 256)
	for i := 0; i < 256; i++ {
		rainbow[i] = getRainbowRGB(uint8(i))
	}

	for {
		logo()
		if quit {
			break
		}
		scroll("This badge", "runs", "TINYGO")
		if quit {
			break
		}
		myNameIsRainbow(YourName)
		if quit {
			break
		}
	}
}

func myNameIs(name string) {
	display.FillScreen(colors[WHITE])

	var r int16 = 8

	// black corners detail
	display.FillRectangle(0, 0, r, r, colors[BLACK])
	display.FillRectangle(0, HEIGHT-r, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, 0, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, HEIGHT-r, r, r, colors[BLACK])

	// round corners
	tinydraw.FilledCircle(&display, r, r, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, r, r, colors[RED])
	tinydraw.FilledCircle(&display, r, HEIGHT-r-1, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, HEIGHT-r-1, r, colors[RED])

	// top band
	display.FillRectangle(r, 0, WIDTH-2*r-1, r, colors[RED])
	display.FillRectangle(0, r, WIDTH, 26, colors[RED])

	// bottom band
	display.FillRectangle(r, HEIGHT-r-1, WIDTH-2*r-1, r+1, colors[RED])
	display.FillRectangle(0, HEIGHT-2*r-1, WIDTH, r, colors[RED])

	// top text : my NAME is
	w32, _ := tinyfont.LineWidth(&fonts.Regular12pt7b, "my NAME is")
	tinyfont.WriteLine(&display, &fonts.Regular12pt7b, (WIDTH-int16(w32))/2, 24, "my NAME is", colors[WHITE])

	// middle text
	w32, _ = tinyfont.LineWidth(&fonts.Bold9pt7b, name)
	tinyfont.WriteLine(&display, &fonts.Bold9pt7b, (WIDTH-int16(w32))/2, 72, name, colors[BLACK])

	// gophers
	tinyfont.WriteLineColors(&display, &fonts.Regular32pt, WIDTH-48, 110, "BE", []color.RGBA{getRainbowRGB(100), getRainbowRGB(200)})
}

func myNameIsRainbow(name string) {
	myNameIs(name)

	w32, _ := tinyfont.LineWidth(&fonts.Bold9pt7b, name)
	for i := 0; i < 230; i++ {
		tinyfont.WriteLineColors(&display, &fonts.Bold9pt7b, (WIDTH-int16(w32))/2, 72, name, rainbow[i:])
		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			quit = true
			break
		}
	}
}

func blinky(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, topline)
	w32bottom, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, bottomline)
	for i := int16(0); i < 10; i++ {
		// show black text
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32top))/2, 50, topline, colors[BLACK])
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 100, bottomline, colors[BLACK])

		// repeat the other way around
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32top))/2, 50, topline, colors[WHITE])
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 100, bottomline, colors[WHITE])

		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			quit = true
			break
		}
	}
}

func blinkyRainbow(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, topline)
	w32bottom, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, bottomline)
	for i := int16(0); i < 20; i++ {
		// show black text
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32top))/2, 50, topline, getRainbowRGB(uint8(i*12)))
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 100, bottomline, getRainbowRGB(uint8(i*12)))

		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			quit = true
			break
		}
	}
}

func scroll(topline, middleline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text, so we could center them later
	w32top, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, topline)
	w32middle, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, middleline)
	w32bottom, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, bottomline)
	tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32top))/2, 34, topline, getRainbowRGB(200))
	tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32middle))/2, 60, middleline, getRainbowRGB(80))
	tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 100, bottomline, getRainbowRGB(120))

	display.SetScrollArea(0, 0)
	for k := 0; k < 4; k++ {
		for i := int16(159); i >= 0; i-- {

			pressed, _ = buttons.Read8Input()
			if pressed&machine.BUTTON_SELECT_MASK > 0 {
				quit = true
				break
			}
			display.SetScroll(i)
			time.Sleep(10 * time.Millisecond)
		}
	}
	display.SetScroll(0)
	display.StopScroll()
}

func logo() {
	display.FillRectangleWithBuffer(0, 0, WIDTH, HEIGHT, logoRGBA)
	time.Sleep(logoDisplayTime)
}
