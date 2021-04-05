package main

import (
	"com/labs/pg/gadget"
	"fmt"
)

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	/*var player Player
	player = gadget.TapePlayer{}
	mixTape := []string{"Jessie's Girl","Whip It","9 to 5"}
	playList(player,mixTape)
	player = gadget.TapeRecorder{}
	playList(player,mixTape)
	play(Robot("Botco Ambler"))
	recorder:=player.(gadget.TapeRecorder)
	recorder.Record()*/
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	record, ok := player.(gadget.TapeRecorder)
	if ok {
		record.Record()
	}
}

type Player interface {
	Play(string)
	Stop()
}

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk")
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}
func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
	robot := n.(Robot)
	robot.Walk()
}
