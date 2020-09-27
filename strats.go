package main

import (
	"github.com/YuriyNasretdinov/gophdragbench/strats/alik"
	"github.com/YuriyNasretdinov/gophdragbench/strats/yourock"
	"github.com/quasilyte/gophers-and-dragons/game"
)

type strat struct {
	name string
	cb   func(game.State) game.CardType
}

var strats = []strat{
	{"yourock/hero", yourock.Hero},
	{"alik/WiningTactic", alik.WiningTactic},
}