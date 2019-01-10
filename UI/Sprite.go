package UI

import (
	"github.com/hajimehoshi/ebiten"
)

type Sprite struct {
	frames      []*ebiten.Image
	curFrame    int
	totalFrames int
	frameCycle  int
	curCycle    int
	noLoop      bool
}

func NewSprite(imgs ...*ebiten.Image) *Sprite {
	return &Sprite{
		frames:      imgs,
		curFrame:    0,
		frameCycle:  4,
		totalFrames: len(imgs),
	}
}

func (s *Sprite) GetCurrentFrame() *ebiten.Image {
	if s.noLoop && s.curFrame == s.totalFrames-1 {
		return s.frames[s.curFrame]
	}
	s.curCycle++
	s.curFrame += s.curCycle / s.frameCycle
	s.curCycle %= s.frameCycle
	s.curFrame %= s.totalFrames
	return s.frames[s.curFrame]
}

func (s *Sprite) GetAnimationLength() int {
	return s.totalFrames * s.frameCycle
}
