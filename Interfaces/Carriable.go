package Interfaces

type Carriable interface {
	Use(player *Player, values ...interface{})
	GetName() string
}
