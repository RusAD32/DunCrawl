package Traps

import . "../Interfaces"

type NoTrap struct{}

func (*NoTrap) Trigger(p *Player) {}
