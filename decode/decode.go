package decode

import (
	"errors"
	"github.com/g-thome/8086-simulator/instructions"
	"github.com/g-thome/8086-simulator/memory"
	"github.com/g-thome/8086-simulator/registers"
	"math/bits"
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

func toSignedInt(number uint32) int32 {
	bitCount := bits.Len(uint(number))

	if bitCount <= 8 {
		return int32(int8(number))
	} else if bitCount <= 16 {
		return int32(int16(number))
	} else if bitCount <= 32 {
		return int32(int16(number))
	} else {
		return int32(number)
	}
}

func GetRegOperand(intelRegIndex uint32, wide uint32) instructions.InstructionOperand {
	regTable := [8][2]registers.RegisterAccess{
		{{registers.REGISTER_A, 0, 1}, {registers.REGISTER_A, 0, 2}},
		{{registers.REGISTER_C, 0, 1}, {registers.REGISTER_C, 0, 2}},
		{{registers.REGISTER_D, 0, 1}, {registers.REGISTER_D, 0, 2}},
		{{registers.REGISTER_B, 0, 1}, {registers.REGISTER_B, 0, 2}},
		{{registers.REGISTER_A, 1, 1}, {registers.REGISTER_SP, 0, 2}},
		{{registers.REGISTER_C, 1, 1}, {registers.REGISTER_BP, 0, 2}},
		{{registers.REGISTER_D, 1, 1}, {registers.REGISTER_SI, 0, 2}},
		{{registers.REGISTER_B, 1, 1}, {registers.REGISTER_DI, 0, 2}},
	}

	return instructions.InstructionOperand{
		Type:     instructions.OPERAND_REGISTER,
		Register: regTable[intelRegIndex&0x7][wide],
	}
}

func parseDataValue(m *memory.Memory, access *memory.SegmentedAccess, exists bool, wide bool, signExtended bool) uint32 {
	var result uint32

	if exists {
		if wide {
			d0 := memory.ReadMemory(m, memory.GetAbsoluteAddressOf(access, 0))
			d1 := memory.ReadMemory(m, memory.GetAbsoluteAddressOf(access, 1))
			result = (uint32(d1) << 8) | uint32(d0)
			access.SegmentOffset += 2
		} else {
			result = uint32(memory.ReadMemory(m, memory.GetAbsoluteAddressOf(access, 0)))
			if signExtended {
				result = uint32(int8(result & 0xff))
			}

			access.SegmentOffset += 1
		}

	}

	return result
}

func TryDecode(ctx *DisasmContext, inst instructions.InstructionFormat, m *memory.Memory, at memory.SegmentedAccess) (instructions.Instruction, error) {
	result := instructions.Instruction{}
	hasBits := 0
	var bits [instructions.BITS_COUNT]uint32

	startingAddress := memory.GetAbsoluteAddressOf(&at, 0)

	var bitsPendingCount uint8
	var bitsPending byte

	for bitsIdx := 0; bitsIdx < len(inst.Bits); bitsIdx++ {
		testBits := inst.Bits[bitsIdx]
		if testBits.Usage == instructions.BITS_LITERAL && testBits.BitCount == 0 {
			break
		}

		readBits := testBits.Value

		if testBits.BitCount != 0 {
			if bitsPendingCount == 0 {
				bitsPendingCount = 8
				bitsPending = memory.ReadMemory(m, memory.GetAbsoluteAddressOf(&at, 0))
				at.SegmentOffset++
			}

			if testBits.BitCount > bitsPendingCount {
				panic("instruction too large")
			}

			bitsPendingCount -= testBits.BitCount
			readBits = bitsPending
			readBits >>= bitsPendingCount
			readBits &= ^(0xff << testBits.BitCount)
		}

		if testBits.Usage == instructions.BITS_LITERAL {
			if readBits != testBits.Value {
				return result, errors.New("Bits don't match encoding")
			}
		} else {
			bits[testBits.Usage] |= uint32(readBits << testBits.Shift)
			hasBits |= 1 << testBits.Usage
		}
	}

	mod := bits[instructions.BITS_MOD]
	rm := bits[instructions.BITS_RM]
	w := bits[instructions.BITS_W]
	s := bits[instructions.BITS_S]
	d := bits[instructions.BITS_D]

	hasDirectAddress := (mod == 0b00) && (rm == 0b110)
	hasDisplacement := bits[instructions.BITS_HAS_DISP] > 0 || mod == 0b10 || mod == 0b01 || hasDirectAddress
	displacementIsW := bits[instructions.BITS_DISP_ALWAYS_W] > 0 || mod == 0b10 || hasDirectAddress
	dataIsW := bits[instructions.BITS_W_MAKES_DATA_W] > 0 && s == 0 && w > 0

	bits[instructions.BITS_DISP] |= parseDataValue(m, &at, hasDisplacement, displacementIsW, !displacementIsW)
	bits[instructions.BITS_DATA] |= parseDataValue(m, &at, bits[instructions.BITS_HAS_DATA] > 0, dataIsW, s > 0)

	result.Op = inst.Op
	result.Flags = ctx.AdditionalFlags
	result.Address = startingAddress
	result.Size = memory.GetAbsoluteAddressOf(&at, 0) - startingAddress

	if w > 0 {
		result.Flags |= instructions.INST_WIDE
	}

	disp := bits[instructions.BITS_DISP]
	displacement := int16(disp)

	var regOperand *instructions.InstructionOperand
	var modOperand *instructions.InstructionOperand

	if d > 0 {
		regOperand = &result.Operands[0]
		modOperand = &result.Operands[1]
	} else {
		regOperand = &result.Operands[1]
		modOperand = &result.Operands[0]
	}

	if hasBits&(1<<instructions.BITS_SR) == 1 {
		regOperand.Type = instructions.OPERAND_REGISTER
		regOperand.Register.Index = registers.RegisterIndex(uint32(registers.REGISTER_ES) + (bits[instructions.BITS_SR] & 0x3))
		regOperand.Register.Count = 2
	}

	if (hasBits & (1 << instructions.BITS_REG)) > 0 {
		*regOperand = GetRegOperand(bits[instructions.BITS_REG], w)
	}

	if (hasBits & (1 << instructions.BITS_MOD)) > 0 {
		if mod == 0b11 {
			if w > 0 {
				*modOperand = GetRegOperand(rm, w)
			} else {
				*modOperand = GetRegOperand(rm, bits[instructions.BITS_RM_REG_ALWAYS_W])
			}
		} else {
			modOperand.Type = instructions.OPERAND_MEMORY
			modOperand.Address.Segment = ctx.DefaultSegment
			modOperand.Address.Displacement = int32(displacement)

			if (mod == 0b00) && (rm == 0b110) {
				modOperand.Address.Base = instructions.EFFECTIVE_ADDRESS_DIRECT
			} else {
				modOperand.Address.Base = rm + 1
			}
		}
	}

	lastOperand := &result.Operands[0]

	if lastOperand.Type > 0 {
		lastOperand = &result.Operands[1]
	}

	if bits[instructions.BITS_REL_JMP_DISP] > 0 {
		lastOperand.Type = instructions.OPERAND_RELATIVE_IMMEDIATE
		lastOperand.Immediate = instructions.Immediate{int32(displacement) + int32(result.Size), true}
	}

	if bits[instructions.BITS_HAS_DATA] > 0 {
		lastOperand.Type = instructions.OPERAND_IMMEDIATE
		lastOperand.Immediate = instructions.Immediate{toSignedInt(bits[instructions.BITS_DATA]), false}
	}

	if (hasBits & (1 << instructions.BITS_V)) > 0 {
		if bits[instructions.BITS_V] > 0 {
			lastOperand.Type = instructions.OPERAND_REGISTER
			lastOperand.Register.Index = registers.REGISTER_C
			lastOperand.Register.Offset = 0
			lastOperand.Register.Count = 1
		} else {
			lastOperand.Type = instructions.OPERAND_IMMEDIATE
			lastOperand.Immediate = instructions.Immediate{toSignedInt(bits[instructions.BITS_DATA]), false}
		}
	}

	return result, nil
}

func DecodeInstruction(ctx *DisasmContext, m *memory.Memory, at *memory.SegmentedAccess) (instructions.Instruction, error) {
	var instruction instructions.Instruction
	for _, f := range instructions.InstructionFormats {
		instruction, err := TryDecode(ctx, f, m, *at)
		if err == nil {
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
