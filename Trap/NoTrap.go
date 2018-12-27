package Traps

import . "DunCrawl/Interfaces"

type NoTrap struct{}

func (*NoTrap) Trigger(p *Player) {}
