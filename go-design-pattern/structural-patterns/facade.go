package structuralpatterns

import "fmt"

// 定义多个子系统和组件

type CPU struct {
	speed float32
}

func (c *CPU) run() {
	fmt.Println("CPU running at", c.speed, "GHz")
}

type Memory struct {
	capacity int
}

func (m *Memory) load() {
	fmt.Println("Memory loaded with", m.capacity, "GB")
}

type Disk struct {
	size int
}

func (d *Disk) read() {
	fmt.Println("Disk reading data with size", d.size, "KB/s")
}

// 定义外观接口

type ComputerFacade struct {
	cpu    *CPU
	memory *Memory
	disk   *Disk
}

func NewComputerFacade() *ComputerFacade {
	cpu := &CPU{speed: 3.4}
	memory := &Memory{capacity: 16}
	disk := &Disk{size: 1024}

	return &ComputerFacade{cpu: cpu, memory: memory, disk: disk}
}

func (f *ComputerFacade) Run() {
	fmt.Println("Computer starting...")
	f.cpu.run()
	f.memory.load()
	f.disk.read()
	fmt.Println("Computer running...")
}
