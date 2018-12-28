package UI

import (
	. "DunCrawl/Interfaces"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
	"io/ioutil"
	"net/http"
	"runtime"
)

var keyDirMap = map[ebiten.Key]Direction{
	ebiten.KeyLeft:  Left,
	ebiten.KeyUp:    Forward,
	ebiten.KeyRight: Right,
	ebiten.KeyDown:  Back,
}

func LoadResource(name string) ([]byte, error) {
	if runtime.GOOS == "js" || runtime.GOARCH == "js" {
		resp, err := http.Get(fmt.Sprintf("resources/%s", name))
		if err != nil {
			return nil, err
		}
		result := make([]byte, 200000)
		n, err := resp.Body.Read(result)
		if err != nil {
			return nil, err
		}
		return result[:n], nil
	} else {
		return ioutil.ReadFile(fmt.Sprintf("resources/%s", name))
	}
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

func checkKey(key ebiten.Key, l *Labyrinth) {
	if inpututil.IsKeyJustPressed(key) {
		dir, ok := keyDirMap[key]
		if !ok {
			panic("no such button should be checked!")
		}
		neighbour := l.GetSliceNeighbours()[int(dir)]
		if neighbour {
			go func() { l.GoToRoom(dir) }()
			<-l.GetEventsChan()
		}
	}
}

func MoveThroughLabyrinth(l *Labyrinth) {
	checkKey(ebiten.KeyDown, l)
	checkKey(ebiten.KeyUp, l)
	checkKey(ebiten.KeyLeft, l)
	checkKey(ebiten.KeyRight, l)
}
