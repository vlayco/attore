package main

import (
	"fmt"
	"time"

	"github.com/vlayco/attore/attore"
)

type Connect struct {
	Name string
}

type ServerActor struct {
	e            *attore.Engine
	gameStatePid string
}

func NewServerActor(e *attore.Engine, gameStatePid string) *ServerActor {
	return &ServerActor{
		e:            e,
		gameStatePid: gameStatePid,
	}
}

func (a *ServerActor) Receive(msg any) {
	switch v := msg.(type) {
	case *Connect:
		fmt.Println("new player connected to the server: ", v.Name)
		a.e.Send(a.gameStatePid, &PlayerInfo{v.Name})

	}
}

type PlayerInfo struct {
	Name string
}

type GameStateActor struct {
	player string
}

func NewGameStateActor() *GameStateActor {
	return &GameStateActor{}
}

func (a *GameStateActor) Receive(msg any) {
	switch v := msg.(type) {
	case *PlayerInfo:
		fmt.Println("setting player name: ", v.Name)
		a.player = v.Name
	}
}

func main() {
	e := attore.NewEngine()
	gameStatePid := e.Spawn(NewGameStateActor(), "GAME_STATE")
	serverPid := e.Spawn(NewServerActor(e, gameStatePid), "SERVER")

	e.Send(serverPid, &Connect{Name: "Vlad The Impaler"})

	time.Sleep(time.Second * 10)
}
