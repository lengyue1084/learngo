package main

type Monster struct {
	Name string
}

func NewMonster(name string) Monster {
	return Monster{
		Name: name,
	}
}
