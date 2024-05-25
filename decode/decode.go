package decode

import (
	"errors"
	"github.com/g-thome/8086-simulator/instructions"
	"github.com/g-thome/8086-simulator/memory"
	"github.com/g-thome/8086-simulator/registers"
)

type DisasmContext struct {
	DefaultSegment  registers.RegisterIndex
	AdditionalFlags uint32
}

func DefaultDisAsmContext() DisasmContext {
	dc := DisasmContext{}
	dc.DefaultSegment = registers.REGISTER_DS

	return dc
}

func TryDecode(ctx *DisasmContext, inst instructions.InstructionFormat, m *memory.Memory, at *memory.SegmentedAccess) (instructions.Instruction, error) {
	result := instructions.Instruction{}
	return result, nil
}

func DecodeInstruction(ctx *DisasmContext, m *memory.Memory, at *memory.SegmentedAccess) (instructions.Instruction, error) {
	var instruction instructions.Instruction
	for _, f := range instructions.InstructionFormats {
		instruction, err := TryDecode(ctx, f, m, at)
		if err != nil {
			at.SegmentOffset += instruction.Size
			return instruction, nil
		}
	}

	return instruction, errors.New("Instruction doesn't match known formats")
}

func UpdateContext(ctx *DisasmContext, inst instructions.Instruction) {
	ctx.AdditionalFlags = 0
	ctx.DefaultSegment = registers.REGISTER_DS
}
