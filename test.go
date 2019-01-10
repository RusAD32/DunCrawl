package main

import (
	. "DunCrawl/Generator"
	. "DunCrawl/Interfaces"
	"DunCrawl/UI"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math"
)

func labGenTest() {
	l := GenerateLabyrinth(9, 15)
	PrintLabyrinth(l)
}

var l *Labyrinth
var g UI.UIGame

func update(screen *ebiten.Image) error {
	//UI.MoveThroughLabyrinth(l)
	_ = screen.Fill(color.White)
	g.Update()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	g.Draw(screen)
	//ebitenutil.DebugPrintAt(screen, UI.PrintMemUsage(), 0, 300)
	//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%f", scale), 300, 200)
	//w, h := screen.Size()
	//UI.DrawLabyrinth(screen, &l,5, 5, w/5, h/5, color.Black)
	return nil
}

func ebitenTest() {
	l = GenerateLabyrinth(10, 10)
	g.Init(l, 600, 480)
	PrintLabyrinth(l)
	//go UI.EnterLabyrinth(&l)
	x, y := ebiten.ScreenSizeInFullscreen()
	//scale := ebiten.DeviceScaleFactor()
	scale := math.Max(ebiten.DeviceScaleFactor(), 1200/float64(x))
	scale = math.Max(scale, 960/float64(y))
	if err := ebiten.Run(update, 600, 480, 2.0/scale, "DunCrawl"); err != nil {
		panic(err.Error())
	}
}

func main() {
	ebitenTest()
}
