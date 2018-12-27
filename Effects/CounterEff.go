package Effects

import . "DunCrawl/Interfaces"

type CounterEff struct {
	id EffectID
	cd int
}

// как должен работать стан на плеера?
// Пока основная идея -- просто пропустить ход. Но это очень разрушительно
// Тогда у вражеских скиллов на стан должна быть маленькая скорость и не 100% шанс
// И сами враги со станом редкие и либо в остальном слабые, либо минибоссы и боссы
// Другие варианы:
// - не использовать самый быстрый скилл (логично, но может быть не очень честно)
// - не использовать скилл на первого врага (не логично и тоже не очень честно)
// - не использовать хилобафы на следующий ход (честно, но не очень логично)
// Надо подумать
// Апдейт: я тупой. Следующий скилл после получения стана. Лол.
func (s *CounterEff) Init(values ...interface{}) Effect {
	s.id = CounterAtk
	if len(values) == 0 {
		s.cd = 4
	} else {
		s.cd = values[0].(int)
	}
	return Effect(s)
}

func (s *CounterEff) GetID() EffectID {
	return s.id
}

func (s *CounterEff) GetAmount() int {
	return 0
}

func (s *CounterEff) GetInfo() string {
	return "Counter"
}

func (s *CounterEff) DecreaseCD() {
	s.cd--
}

func (s *CounterEff) GetCD() int {
	return s.cd
}
