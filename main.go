package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Initialize a window (Raylib translates this to the Android fullscreen surface automatically)
	rl.InitWindow(800, 450, "Single File App")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Clear screen to an absolute background color
		rl.ClearBackground(rl.RayWhite)

		// Draw your UI elements / text directly
		rl.DrawText("Congrats! Built from a single root file.", 100, 200, 20, rl.LightGray)
		rl.DrawRectangle(100, 250, 200, 50, rl.Maroon)
		rl.DrawText("Native APK Button", 120, 265, 16, rl.White)

		rl.EndDrawing()
	}
}
