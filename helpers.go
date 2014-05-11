package main

import (
	"github.com/blang/gosqm"
)

func exportGroup(g *gosqm.Group) bool {
	if g.Side == "LOGIC" {
		return false
	}
	for _, u := range g.Units {
		// if u.Player != "" {
		// 	return false
		// }
		if u.Side == "LOGIC" {
			return false
		}
	}
	return true
}
