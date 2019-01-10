package Effects

import . "DunCrawl/Interfaces"

type basicEffect struct {
	id     EffectID
	cd     int
	info   string
	amount int
}

func newBasicEffect(id EffectID, cd int, info string, amount int) *basicEffect {
	return &basicEffect{
		id:     id,
		cd:     cd,
		info:   info,
		amount: amount,
	}
}

func (be *basicEffect) GetID() EffectID {
	return be.id
}

func (be *basicEffect) GetAmount() int {
	return be.amount
}

func (be *basicEffect) GetInfo() string {
	return be.info
}

func (be *basicEffect) DecreaseCD() {
	be.cd--
}

func (be *basicEffect) GetCD() int {
	return be.cd
}
