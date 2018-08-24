package emulator

type Cpu struct {
	a byte
	f byte
	b byte
	c byte
	d byte
	e byte
	h byte
	l byte
	sp uint16
	pc uint16
}

type Instruction struct {
	Name string
	Cycles int
	Run func(cpu *Cpu, ram *Ram)
}

var instructions = map[int]Instruction{
	1 : {"LOD", 3, func(cpu *Cpu, ram *Ram) {}},
}

func (cpu *Cpu) emulateCycle(ram *Ram)  {
	opCode := ram.load(cpu.pc)
	instructions[int(opCode)].Run(cpu, ram)
}

func load(high byte, low byte, value uint16)  {

}