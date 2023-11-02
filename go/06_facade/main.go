package main

import "fmt"

type CPU struct{}

func (c *CPU) start() {
	fmt.Println("CPU starting")
}

func (c *CPU) execute() {
	fmt.Println("CPU executing")
}

type Memory struct{}

func (m *Memory) load() {
	fmt.Println("Memory loaded")
}

type Disk struct{}

func (d *Disk) read() {
	fmt.Println("Disk reading")
}

// facade
type ComputerFacade struct {
	cpu    CPU
	memory Memory
	disk   Disk
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:    CPU{},
		memory: Memory{},
		disk:   Disk{},
	}
}

func (cf *ComputerFacade) start() {
	cf.cpu.start()
	cf.memory.load()
	cf.disk.read()
	cf.cpu.execute()
}

func main() {
	cf := NewComputerFacade()

	cf.start()
}
