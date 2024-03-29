package Effects

import . "DunCrawl/Interfaces"

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
func NewStun(cd int) Effect {
	return newBasicEffect(Stun, cd, "Stun", 0)
}
