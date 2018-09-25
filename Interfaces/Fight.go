package Interfaces

import (
	"container/heap"
	"fmt"
	"strconv"
)

type Fight struct {
	p       *Player
	enemies []*Enemy
	pq      PriorityQueue
}

func (f *Fight) Turn(p *Player, enemies []*Enemy) {
	prompt := "Choose a skill to use on yourself\n"
	for i, v := range p.SelfSkills {
		prompt += fmt.Sprintf("%d. %s\n", i, v.GetName())
	}
	res := Prompt(prompt, MakeStrRange(1, len(p.SelfSkills)-1))
	var chosenSelfSkill PlayerSelfSkill
	if res != "" {
		skillNum, err := strconv.Atoi(res)
		if err != nil {
			Inform("Prompt returned bad value: " + res)
		}
		chosenSelfSkill = p.SelfSkills[skillNum]
	} else {
		Inform("Prompt returned empty string")
		return
	}
	f.pq.Push(chosenSelfSkill)
	for i, v := range enemies {
		//TODO я устал, потом накодю
	}
}

func (f *Fight) StartFight(p *Player, enemies []*Enemy) {
	heap.Init(&f.pq)
	for len(enemies) > 0 && p.CurMentHP > 0 && p.CurPhysHP > 0 {

	}
}
