package UI

import (
	. "DunCrawl/Interfaces"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

type consts struct {
	// constants regarding position of different elements of UI
	labXPos, labYPos, labW, labH                             int
	doorX, doorXOff, doorY, doorH, doorW                     int
	backdoorX, backdoorY, backdoorH, backdoorW               int
	enemyX, enemyXOff, enemyY, enemyH, enemyW                int
	hpX, hpY, hpW, hpH, infoX, infoY, statusX, statusY       int
	selfSkButX, skButXOff, skButY, skButW, skButH, dmgSkButX int
	// end of constant declaration
}

func getConstants(w, h int) consts {
	c := consts{}
	c.labXPos = 5
	c.labYPos = 5
	c.labW = w / 5
	c.labH = h / 5

	c.doorX = w * 15 / 100
	c.doorXOff = w / 4
	c.doorY = h / 5
	c.doorW = w * 2 / 10
	c.doorH = h * 5 / 8

	c.enemyX = 0
	c.enemyXOff = w / 4
	c.enemyY = h / 4
	c.enemyW = w / 4
	c.enemyH = w / 4

	c.hpX = w / 10
	c.hpY = h * 8 / 10
	c.hpW = w * 8 / 10
	c.hpH = h / 16
	c.infoX = w / 3
	c.infoY = h * 9 / 10
	c.statusX = w * 8 / 10
	c.statusY = h * 9 / 10

	c.backdoorX = w * 2 / 10
	c.backdoorY = h * 9 / 10
	c.backdoorW = w * 6 / 10
	c.backdoorH = h / 10

	c.selfSkButX = w / 16
	c.dmgSkButX = w/2 + w/16
	c.skButY = h * 66 / 100
	c.skButXOff = w / 8
	c.skButW = w / 10
	c.skButH = h / 10
	return c
	// end of constant declaration
}

func LoadResource(name string) ([]byte, error) {
	reader, err := ebitenutil.OpenFile(fmt.Sprintf("resources/%s", name))
	result := make([]byte, 200000)
	n, err := reader.Read(result)
	if err != nil {
		return nil, err
	}
	return result[:n], nil
}

func getRoomCoords(i, j, startX, startY, roomW, roomH int) (float64, float64) {
	return float64(roomW*j + startX), float64(roomH*i + startY)
}

func drawTriangle(screen *ebiten.Image, ax, ay, bx, by, cx, cy float64, col color.Color) {
	ebitenutil.DrawLine(screen, ax, ay, bx, by, col)
	ebitenutil.DrawLine(screen, bx, by, cx, cy, col)
	ebitenutil.DrawLine(screen, cx, cy, ax, ay, col)
}

func DrawPlayer(screen *ebiten.Image, x, y float64, roomW, roomH, dir int, col color.Color) {
	playerW := float64(roomW) * 0.8
	playerWoffs := float64(roomW) * 0.1
	playerH := float64(roomH) * 0.8
	playerHoffs := float64(roomH) * 0.1
	switch dir {
	case 0:
		{
			ax := x + playerWoffs
			ay := y + playerHoffs
			bx := ax + playerW
			by := ay + playerH/2.0
			cx := ax
			cy := ay + playerH
			drawTriangle(screen, ax, ay, bx, by, cx, cy, col)
		}
	case 3:
		{
			ax := x + playerWoffs
			ay := y + playerHoffs + playerH
			bx := ax + playerW/2
			by := ay - playerH
			cx := ax + playerW
			cy := ay
			drawTriangle(screen, ax, ay, bx, by, cx, cy, col)
		}
	case 2:
		{
			ax := x + playerWoffs + playerW
			ay := y + playerHoffs
			bx := ax - playerW
			by := ay + playerH/2.0
			cx := ax
			cy := ay + playerH
			drawTriangle(screen, ax, ay, bx, by, cx, cy, col)
		}
	case 1:
		{
			ax := x + playerWoffs
			ay := y + playerHoffs
			bx := ax + playerW/2.0
			by := ay + playerH
			cx := ax + playerW
			cy := ay
			drawTriangle(screen, ax, ay, bx, by, cx, cy, col)
		}
	}
}

func DrawLabyrinth(screen *ebiten.Image, l *Labyrinth, startX, startY, w, h int, col color.Color) {
	roomW := (w - startX*2) / l.GetWidth()
	roomH := (h - startY*2) / l.GetLength()
	rooms := l.GetRooms()
	for i := 0; i < l.GetWidth(); i++ {
		for j := 0; j < l.GetLength(); j++ {
			room := rooms[i*l.GetLength()+j]
			walls := room.GetNeighbours()
			roomX, roomY := getRoomCoords(i, j, startX, startY, roomW, roomH)
			nextX := roomX + float64(roomW)
			nextY := roomY + float64(roomH)
			if !walls[int(Forward)].CanGoThrough() {
				ebitenutil.DrawLine(screen, roomX, roomY, roomX+float64(roomW), roomY, col)
			}
			if !walls[int(Left)].CanGoThrough() {
				ebitenutil.DrawLine(screen, roomX, roomY, roomX, nextY, col)
			}
			if !walls[int(Right)].CanGoThrough() {
				ebitenutil.DrawLine(screen, nextX, roomY, nextX, nextY, col)
			}
			if !walls[int(Back)].CanGoThrough() {
				ebitenutil.DrawLine(screen, roomX, nextY, nextX, nextY, col)
			}
			if rooms[i*l.GetLength()+j] == l.GetCurrentRoom() {
				DrawPlayer(screen, roomX, roomY, roomW, roomH, l.GetPrevious(), col)
			}
		}
	}
}
