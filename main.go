package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type App struct{}

// Update handles your application logic (runs 60 times a second)
func (a *App) Update() error {
	return nil
}

// Draw handles rendering your graphics to the screen
func (a *App) Draw(screen *ebiten.Image) {
	// Clear the screen with a background color
	screen.Fill(color.RGBA{0x22, 0x22, 0x22, 0xff})

	// Print text directly onto the interface
	ebitenutil.DebugPrint(screen, "Congrats! Single-file Pure Go APK is running!")
}

// Layout handles scaling across different mobile screens
func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Single File Pure Go App")

	if err := ebiten.RunGame(&App{}); err != nil {
		log.Fatal(err)
	}
}
