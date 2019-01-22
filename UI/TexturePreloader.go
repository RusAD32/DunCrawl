package UI

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type TexPreloader map[string]*ebiten.Image

func NewTexPreloader() *TexPreloader {
	m := TexPreloader(make(map[string]*ebiten.Image))
	return &m
}

func (p *TexPreloader) GetImgByPath(path string) *ebiten.Image {
	img, ok := (*p)[path]
	if !ok {
		pic, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterLinear)
		(*p)[path] = pic
		if err != nil {
			panic(err)
		}
		return pic
	}
	return img
}

func (p *TexPreloader) DeleteAllButThese(paths []string) {
	for k := range *p {
		if !contains(paths, k) {
			delete(*p, k)
		}
	}
}

func (p *TexPreloader) EnsureThese(paths []string) {
	p.DeleteAllButThese(paths)
	for _, path := range paths {
		p.GetImgByPath(path)
	}
}

func contains(arr []string, el string) bool {
	for _, v := range arr {
		if v == el {
			return true
		}
	}
	return false
}
