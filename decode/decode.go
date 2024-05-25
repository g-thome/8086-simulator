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
	hasBits := 0
	var bits uint32
	valid := true

	startingAddress := memory.GetAbsoluteAddressOf(at.SegmentBase, at.SegmentOffset, 0)

	bitsPendingCount := 0
	bitsPending := 0

	for bitsIdx := 0; valid && bitsIdx < len(inst.Bits); bitsIdx++ {
		testBits := inst.Bits[bitsIdx]
		if testBits.Usage == instructions.BITS_LITERAL && testBits.BitCount == 0 {
			break
		}

		readBits := testBits.Value

		if testBits.BitCount != 0 {
			if bitsPendingCount == 0 {
				bitsPendingCount = 8
				bitsPending = memory.ReadMemory(m, memory.GetAbsoluteAddressOf(at.SegmentBase, at.SegmentOffset, 0))
				at.SegmentOffset++
			}

			if testBits.BitCount <= bitsPendingCount {
				panic("instruction too large")
			}

			bitsPendingCount -= testBits.BitCount
			readBits = bitsPending
			readBits >>= bitsPendingCount
			readBits &= ~(0xff << testBits.BitCount)
		}

		if testBits.Usage == instructions.BITS_LITERAL {
			valid = valid && (readBits == testBits.Value)
		} else {
			bits[testBits.Usage] |= readBits << testBits.Shift
			hasBits |= 1 << testBits.Usage
		}
	}

	if valid {
		mod = bits[instructions.BITS_MOD]
		rm = bits[instructions.BITS_RM]
		w = bits[instructions.W]
		s = bits[instructions.S]
		d = bits[instructions.D]

		hasDirectAccess = (mod == 0b00) && (rm == 0b110)
		hasDisplacement = bits[instructions.BITS_HAS_DISP] || mod == 0b10 || mod == 0b01 || hasDirectAddress
		displacementIsW = bits[instructions.BITS_DISP_ALWAYS_W] || mod == 0b10 || hasDirectAddress
		dataIsW = bits[instructions.BITS_W_MAKES_DATA_W] && !s && w

		bits[instructions.BITS_DISP] |= ParseDataValue(m, at, hasDisplacement, displacementIsW, !displacementIsW)
		bits[instructions.BITS_DATA] |= ParseDataValue(m, at, bits[instructions.BITS_HAS_DATA], dataIsW, s)

		result.Op = inst.Op
		result.Flags = ctx.AdditionalFlags
		result.Address = startingAddress
		result.Size = memory.GetAbsoluteAddressOf(at.SegmentBase, at.SegmentOffset, 0) - startingAddress

		if w {
			result.Flags |= INST_WIDE
		}

		disp := bits[instructions.BITS_DISP]
		displacement := int16(disp)

		var regOperand *instructions.InstructionOperand
		var modOperand *instructions.InstructionOperand

		if d {
			regOperand = &result.Operands[0]
			modOperand = &result.Operands[1]
		} else {
			regOperand = &result.Operands[1]
			modOperand = &result.Operands[0]
		}

		if hasBits&(1<<instructions.BITS_SR) == 1 {
			regOperand.Type = OPERAND_REGISTER
			regOperand.Register.Index = registers.REGISTER_ES + (bits[instructions.BITS_SR] & 0x3)
			regOperand.Register.Count = 2
		}

		if hasBits & (1 << instructions.BITS_REG) {
			regOperand = GetRegOperand(bits[instructions.BITS_REG], w)
		}

		if hasBits & (1 << instructions.BITS_MOD) {
			if mod == 0b11 {
				modOperand = GetRegOperand(rm, w || bits[instructions.BITS_RM_REG_ALWAYS_W])
			} else {
				modOperand.Type = OPERAND_MEMORY
				modOperand.Address.Segment = ctx.DefaultSegment
				modOperand.Address.Displacement = displacement

				if (mod == 0b00) && (rm == 0b110) {
					modOperand.Address.Base = EFFECTIVE_ADDRESS_DIRECT
				} else {
					modOperand.Address.Base = rm + 1
				}
			}
		}

		lastOperand = &result.Operands[0]

		if lastOperand.Type != nil {
			lastOperand = &result.Operands[1]
		}

		if bits[instructions.BITS_REL_JMP_DISP] {
			lastOperand.Type = OPERAND_RELATIVE_IMMEDIATE
			lastOperand.SignedImmediate = displacement + result.Size
		}

		if bits[instructions.BITS_HAS_DATA] {
			lastOperand.Type = OPERAND_IMMEDIATE
			lastOperand.UnsignedImmediate = bits[instructions.BITS_DATA]
		}

		if hasBits & (1 << instructions.BITS_V) {
			if bits[instructions.BITS_V] {
				lastOperand.Type = OPERAND_REGISTER
				lastOperand.Register.Index = instructions.REGISTER_C
				lastOperand.Register.Offset = 0
				lastOperand.Register.Count = 1
			} else {
				lastOperand.Type = OPERAND_IMMEDIATE
				lastOperand.SignedImmediate = 1
			}
		}
	}

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
