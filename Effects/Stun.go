package Effects

import . "../Interfaces"

type StunEffect struct {
	Id EffectID
	CD int
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
func (s *StunEffect) Init(values ...interface{}) Effect {
	s.Id = Stun
	if len(values) == 0 {
		s.CD = 1
	} else {
		s.CD = values[0].(int)
	}
	return Effect(s)
}

func (s *StunEffect) GetID() EffectID {
	return s.Id
}

func (s *StunEffect) GetAmount() int {
	return 0
}

func (s *StunEffect) GetInfo() string {
	return ""
}

func (s *StunEffect) DecreaseCD() {
	s.CD--
}

func (s *StunEffect) GetCD() int {
	return s.CD
}
