package UI

import "github.com/hajimehoshi/ebiten"

type SkQueue struct {
	skills      []*SkillIcon
	x, y, xOffs int
}

func (q *SkQueue) Update(skills []*SkillIcon) {
	q.skills = skills
}

func (q *SkQueue) Draw(screen *ebiten.Image) {
	for i := len(q.skills) - 1; i >= 0; i-- {
		curSk := q.skills[i]
		opts := *curSk.opts
		curSk.x = q.x + q.xOffs*i
		curSk.y = q.y
		curSk.opts.GeoM.Translate(float64(q.x+q.xOffs*i), float64(q.y))
		curSk.Draw(screen)
		*curSk.opts = opts
	}
}
