package Interfaces

func GenerateLabyrinth(length, width int) {
	lab := Labyrinth{
		nil,
		make([]*Room, length*width),
		0,
		nil,
		nil,
		make(chan bool),
		make(chan []SkillInfo),
		make(chan string),
		make(chan Event),
	}
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			lab.rooms[i*length+j] = GenerateRoom(lab.fightBgToUi, lab.fightUiToBg, lab.fightConfirmChan)
		}
	}
}

func GenerateRoom(bgToUi chan []SkillInfo, uiToBg chan string, confirm chan bool) *Room {
	r := new(Room)
	r.Init([]*Enemy{}, bgToUi, uiToBg, confirm)
	return r
}
