//+build wireinject

package main

import "github.com/google/wire"

func InitMission(a string) *Mission {
	wire.Build(NewMonster, NewMission, NewPlayer)
	return &Mission{}
}

func InitMission3(a string) *Mission {
	panic(wire.Build(NewMonster, NewMission, NewPlayer))

}
