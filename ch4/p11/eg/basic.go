package main

import (
	"github.com/glenn-brown/aima/ch4/p11"
	"github.com/glenn-brown/aima/ch4/p11/agent"
	"github.com/glenn-brown/vu"
	"time"
)

func main() {
	world := p11.NewWorld(20)
	agent := agent.New()
	window, err := vu.NewWindow(vu.Flat(vu.Frame(world)))
	if err != nil {
		panic(err)
	}
	for {
		world.Step(agent)
		window.Render()
		time.Sleep(300 * time.Second)
	}
}
