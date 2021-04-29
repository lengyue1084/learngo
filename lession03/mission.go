package main

import "fmt"

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) *Mission {
	return &Mission{
		p, m,
	}
}

func (m *Mission) start() {
	fmt.Printf("勇士：%s战胜了 怪兽:%s", m.Player.Name, m.Monster.Name)
}
