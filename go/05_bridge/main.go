package main

import "fmt"

// Device defines the Device interface
type Device interface {
	TurnOn()
	TurnOff()
}

// TV is a concrete Device implementation
type TV struct{}

func (t TV) TurnOn() {
	fmt.Println("TV is turned on")
}

func (t TV) TurnOff() {
	fmt.Println("TV is turned off")
}

// DVDPlayer is a concrete Device implementation
type DVDPlayer struct{}

func (d DVDPlayer) TurnOn() {
	fmt.Println("DVD player is turned on")
}

func (d DVDPlayer) TurnOff() {
	fmt.Println("DVD player is turned off")
}

// Remote defines the Remote interface
type Remote interface {
	PowerOn()
	PowerOff()
}

// TVRemote is a concrete Remote implementation
type TVRemote struct {
	device Device
}

func (tr TVRemote) PowerOn() {
	tr.device.TurnOn()
}

func (tr TVRemote) PowerOff() {
	tr.device.TurnOff()
}

// DVDRemote is a concrete Remote implementation
type DVDRemote struct {
	device Device
}

func (dr DVDRemote) PowerOn() {
	dr.device.TurnOn()
}

func (dr DVDRemote) PowerOff() {
	dr.device.TurnOff()
}

func main() {
	tv := TV{}
	dvd := DVDPlayer{}

	tvRemote := TVRemote{device: tv}
	dvdRemote := DVDRemote{device: dvd}

	tvRemote.PowerOn()
	tvRemote.PowerOff()

	dvdRemote.PowerOn()
	dvdRemote.PowerOff()
}
