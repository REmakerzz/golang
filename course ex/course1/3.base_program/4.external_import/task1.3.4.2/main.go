package main

import (
	"fmt"
	"github.com/ptflp/gocolors"
)

func main() {
	fmt.Println(ColorizeRed("Hello"))
	fmt.Println(ColorizeGreen("ALLO"))
	fmt.Println(ColorizeBlue("World"))
	fmt.Println(ColorizeYellow("foo"))
	fmt.Println(ColorizeMagenta("boo"))
	fmt.Println(ColorizeCyan("bar"))
	fmt.Println(ColorizeWhite("Kata"))
	fmt.Println(ColorizeCustom("Russian Federation", 100, 200, 50))
}

// ColorizeRed that takes a string and returns it in red.
func ColorizeRed(a string) string {
	return gocolors.Colorize(gocolors.Red, a)
}

// ColorizeGreen that takes a string and returns it in green.
func ColorizeGreen(a string) string {
	return gocolors.Colorize(gocolors.Green, a)
}

// ColorizeBlue that takes a string and returns it in Blue.
func ColorizeBlue(a string) string {
	return gocolors.Colorize(gocolors.Blue, a)
}

// ColorizeYellow that takes a string and returns it in yellow.
func ColorizeYellow(a string) string {
	return gocolors.Colorize(gocolors.Yellow, a)
}

// ColorizeMagenta that takes a string and returns it in magenta.
func ColorizeMagenta(a string) string {
	return gocolors.Colorize(gocolors.Magenta, a)
}

// ColorizeCyan that takes a string and returns it in cyan.
func ColorizeCyan(a string) string {
	return gocolors.Colorize(gocolors.Cyan, a)
}

// ColorizeWhite that takes a string and returns it in white.
func ColorizeWhite(a string) string {
	return gocolors.Colorize(gocolors.White, a)
}

// ColorizeCustom that takes a string and RGB color values and returns it in a custom color.
func ColorizeCustom(a string, r, g, b uint8) string {
	return gocolors.Colorize(gocolors.RGB(int(r), int(g), int(b)), a)
}
