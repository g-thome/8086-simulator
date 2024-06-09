package simulator

import (
	"fmt"
	"github.com/g-thome/8086-simulator/instructions"
	"github.com/g-thome/8086-simulator/registers"
)

type Register struct {
	Value uint32
}

type Simulation struct {
	Registers [registers.REGISTER_COUNT]Register
}

func (s *Simulation) Run(instruction instructions.Instruction) {
	if instruction.Op != instructions.OpMov {
		panic("Unsupported operation type")
	}

	source := instruction.Operands[1]
	destination := instruction.Operands[0]
	reg := &s.Registers[destination.Register.Index]
	reg.Value = uint32(source.Immediate.Value)
}

func (s *Simulation) PrintRegisters() {
	result := ""

	for i, reg := range s.Registers {
		if i == 0 {
			continue
		}
		result += fmt.Sprintf("%v: %v\n", registers.RegisterIndex(i), int(reg.Value))
	}

	fmt.Println(result)
}

func NewSimulation() Simulation {
	sim := Simulation{}

	fmt.Println(sim)
	for i := 0; i < int(registers.REGISTER_COUNT); i++ {
		sim.Registers[registers.RegisterIndex(i)] = Register{0}
	}

	return sim
}
