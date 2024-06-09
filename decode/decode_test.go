package decode

import (
	"github.com/g-thome/8086-simulator/instructions"
	"github.com/g-thome/8086-simulator/memory"
	"github.com/g-thome/8086-simulator/registers"
	"reflect"
	"testing"
)

func TestDecodeMovRegReg(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := DefaultDisAsmContext()
	m := memory.Memory{}
	memory.LoadMemoryFromFile("../fixtures/bin/mov_cx_bx", &m)

	expected := instructions.Instruction{
		Address: 0,
		Size:    2,
		Op:      instructions.OpMov,
		Flags:   8,
		Operands: [2]instructions.InstructionOperand{
			{
				Type: instructions.OPERAND_REGISTER,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_C,
					Offset: 0,
					Count:  2,
				},
				Immediate: instructions.Immediate{},
			},
			{
				Type: instructions.OPERAND_REGISTER,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_B,
					Offset: 0,
					Count:  2,
				},
				Immediate: instructions.Immediate{},
			},
		},
	}

	received, err := DecodeInstruction(&ctx, &m, &at)

	if err != nil {
		t.Fatalf(`Error decoding instruction %v`, err)
	}

	if !reflect.DeepEqual(received, expected) {
		t.Fatalf("Expected \n%+v, \ngot \n%+v", expected, received)
	}
}

func TestDecodeMovRegUnsignedImmediate(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := DefaultDisAsmContext()
	m := memory.Memory{}
	memory.LoadMemoryFromFile("../fixtures/bin/mov_cx_3", &m)

	expected := instructions.Instruction{
		Address: 0,
		Size:    3,
		Op:      instructions.OpMov,
		Flags:   8,
		Operands: [2]instructions.InstructionOperand{
			{
				Type: instructions.OPERAND_REGISTER,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_C,
					Offset: 0,
					Count:  2,
				},
				Immediate: instructions.Immediate{},
			},
			{
				Type: instructions.OPERAND_IMMEDIATE,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_NONE,
					Offset: 0,
					Count:  0,
				},
				Immediate: instructions.Immediate{3, false},
			},
		},
	}

	received, err := DecodeInstruction(&ctx, &m, &at)

	if err != nil {
		t.Fatalf(`Error decoding instruction %v`, err)
	}

	if !reflect.DeepEqual(received, expected) {
		t.Fatalf("Expected \n%+v, \ngot \n%+v", expected, received)
	}
}

func TestDecodeMovRegSignedImmediate(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := DefaultDisAsmContext()
	m := memory.Memory{}
	memory.LoadMemoryFromFile("../fixtures/bin/mov_bx_-12", &m)

	expected := instructions.Instruction{
		Address: 0,
		Size:    3,
		Op:      instructions.OpMov,
		Flags:   8,
		Operands: [2]instructions.InstructionOperand{
			{
				Type: instructions.OPERAND_REGISTER,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_B,
					Offset: 0,
					Count:  2,
				},
				Immediate: instructions.Immediate{},
			},
			{
				Type: instructions.OPERAND_IMMEDIATE,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_NONE,
					Offset: 0,
					Count:  0,
				},
				Immediate: instructions.Immediate{-12, false},
			},
		},
	}

	received, err := DecodeInstruction(&ctx, &m, &at)

	if err != nil {
		t.Fatalf(`Error decoding instruction %v`, err)
	}

	if !reflect.DeepEqual(received, expected) {
		t.Fatalf("Expected \n%+v, \ngot \n%+v", expected, received)
	}
}

func TestDecodeMovRegHighBitsSignedImmediate(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := DefaultDisAsmContext()
	m := memory.Memory{}
	memory.LoadMemoryFromFile("../fixtures/bin/mov_ch_-12", &m)

	expected := instructions.Instruction{
		Address: 0,
		Size:    2,
		Op:      instructions.OpMov,
		Flags:   0,
		Operands: [2]instructions.InstructionOperand{
			{
				Type: instructions.OPERAND_REGISTER,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_C,
					Offset: 1,
					Count:  1,
				},
				Immediate: instructions.Immediate{},
			},
			{
				Type: instructions.OPERAND_IMMEDIATE,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_NONE,
					Offset: 0,
					Count:  0,
				},
				Immediate: instructions.Immediate{-12, false},
			},
		},
	}

	received, err := DecodeInstruction(&ctx, &m, &at)

	if err != nil {
		t.Fatalf(`Error decoding instruction %v`, err)
	}

	if !reflect.DeepEqual(received, expected) {
		t.Fatalf("Expected \n%+v, \ngot \n%+v", expected, received)
	}
}

func TestDecodeMovRegLowBitsImmediate(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := DefaultDisAsmContext()
	m := memory.Memory{}
	memory.LoadMemoryFromFile("../fixtures/bin/mov_cl_12", &m)

	expected := instructions.Instruction{
		Address: 0,
		Size:    2,
		Op:      instructions.OpMov,
		Flags:   0,
		Operands: [2]instructions.InstructionOperand{
			{
				Type: instructions.OPERAND_REGISTER,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_C,
					Offset: 0,
					Count:  1,
				},
				Immediate: instructions.Immediate{},
			},
			{
				Type: instructions.OPERAND_IMMEDIATE,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_NONE,
					Offset: 0,
					Count:  0,
				},
				Immediate: instructions.Immediate{12, false},
			},
		},
	}

	received, err := DecodeInstruction(&ctx, &m, &at)

	if err != nil {
		t.Fatalf(`Error decoding instruction %v`, err)
	}

	if !reflect.DeepEqual(received, expected) {
		t.Fatalf("Expected \n%+v, \ngot \n%+v", expected, received)
	}
}

func TestDecodeMovAccMemory(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := DefaultDisAsmContext()
	m := memory.Memory{}
	memory.LoadMemoryFromFile("../fixtures/bin/mov_ax_[16]", &m)

	expected := instructions.Instruction{
		Address: 0,
		Size:    3,
		Op:      instructions.OpMov,
		Flags:   8,
		Operands: [2]instructions.InstructionOperand{
			{
				Type: instructions.OPERAND_REGISTER,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_A,
					Offset: 0,
					Count:  2,
				},
				Immediate: instructions.Immediate{},
			},
			{
				Type: instructions.OPERAND_MEMORY,
				Address: instructions.EffectiveAddressExpression{
					Segment:      12,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 16,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_NONE,
					Offset: 0,
					Count:  0,
				},
				Immediate: instructions.Immediate{},
			},
		},
	}

	received, err := DecodeInstruction(&ctx, &m, &at)

	if err != nil {
		t.Fatalf(`Error decoding instruction %v`, err)
	}

	if !reflect.DeepEqual(received, expected) {
		t.Fatalf("Expected \n%+v, \ngot \n%+v", expected, received)
	}
}

func TestDecodeMovMemoryReg(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := DefaultDisAsmContext()
	m := memory.Memory{}
	memory.LoadMemoryFromFile("../fixtures/bin/mov_cx_[104]", &m)

	expected := instructions.Instruction{
		Address: 0,
		Size:    4,
		Op:      instructions.OpMov,
		Flags:   8,
		Operands: [2]instructions.InstructionOperand{
			{
				Type: instructions.OPERAND_REGISTER,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_C,
					Offset: 0,
					Count:  2,
				},
				Immediate: instructions.Immediate{},
			},
			{
				Type: instructions.OPERAND_MEMORY,
				Address: instructions.EffectiveAddressExpression{
					Segment:      12,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 104,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_NONE,
					Offset: 0,
					Count:  0,
				},
				Immediate: instructions.Immediate{},
			},
		},
	}

	received, err := DecodeInstruction(&ctx, &m, &at)

	if err != nil {
		t.Fatalf(`Error decoding instruction %v`, err)
	}

	if !reflect.DeepEqual(received, expected) {
		t.Fatalf("Expected \n%+v, \ngot \n%+v", expected, received)
	}
}

func TestDecodeMovFromAddressCalculationToReg(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := DefaultDisAsmContext()
	m := memory.Memory{}
	memory.LoadMemoryFromFile("../fixtures/bin/mov_al_[bx + si]", &m)

	expected := instructions.Instruction{
		Address: 0,
		Size:    2,
		Op:      instructions.OpMov,
		Flags:   0,
		Operands: [2]instructions.InstructionOperand{
			{
				Type: instructions.OPERAND_REGISTER,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_A,
					Offset: 0,
					Count:  1,
				},
				Immediate: instructions.Immediate{},
			},
			{
				Type: instructions.OPERAND_MEMORY,
				Address: instructions.EffectiveAddressExpression{
					Segment:      12,
					Base:         instructions.EFFECTIVE_ADDRESS_BX_SI,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_NONE,
					Offset: 0,
					Count:  0,
				},
				Immediate: instructions.Immediate{},
			},
		},
	}

	received, err := DecodeInstruction(&ctx, &m, &at)

	if err != nil {
		t.Fatalf(`Error decoding instruction %v`, err)
	}

	if !reflect.DeepEqual(received, expected) {
		t.Fatalf("Expected \n%+v, \ngot \n%+v", expected, received)
	}
}

func TestDecodeMovFromAddressCalculationWithDisplacementToReg(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := DefaultDisAsmContext()
	m := memory.Memory{}
	memory.LoadMemoryFromFile("../fixtures/bin/mov_ah_[bx + si + 4]", &m)

	expected := instructions.Instruction{
		Address: 0,
		Size:    3,
		Op:      instructions.OpMov,
		Flags:   0,
		Operands: [2]instructions.InstructionOperand{
			{
				Type: instructions.OPERAND_REGISTER,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_A,
					Offset: 1,
					Count:  1,
				},
				Immediate: instructions.Immediate{},
			},
			{
				Type: instructions.OPERAND_MEMORY,
				Address: instructions.EffectiveAddressExpression{
					Segment:      12,
					Base:         instructions.EFFECTIVE_ADDRESS_BX_SI,
					Displacement: 4,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_NONE,
					Offset: 0,
					Count:  0,
				},
				Immediate: instructions.Immediate{},
			},
		},
	}

	received, err := DecodeInstruction(&ctx, &m, &at)

	if err != nil {
		t.Fatalf(`Error decoding instruction %v`, err)
	}

	if !reflect.DeepEqual(received, expected) {
		t.Fatalf("Expected \n%+v, \ngot \n%+v", expected, received)
	}
}

func TestDecodeMovFromAddressCalculationWithLargeDisplacementToReg(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := DefaultDisAsmContext()
	m := memory.Memory{}
	memory.LoadMemoryFromFile("../fixtures/bin/mov_al_[bx + si + 4999]", &m)

	expected := instructions.Instruction{
		Address: 0,
		Size:    4,
		Op:      instructions.OpMov,
		Flags:   0,
		Operands: [2]instructions.InstructionOperand{
			{
				Type: instructions.OPERAND_REGISTER,
				Address: instructions.EffectiveAddressExpression{
					Segment:      0,
					Base:         instructions.EFFECTIVE_ADDRESS_DIRECT,
					Displacement: 0,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_A,
					Offset: 0,
					Count:  1,
				},
				Immediate: instructions.Immediate{},
			},
			{
				Type: instructions.OPERAND_MEMORY,
				Address: instructions.EffectiveAddressExpression{
					Segment:      12,
					Base:         instructions.EFFECTIVE_ADDRESS_BX_SI,
					Displacement: 4999,
				},
				Register: registers.RegisterAccess{
					Index:  registers.REGISTER_NONE,
					Offset: 0,
					Count:  0,
				},
				Immediate: instructions.Immediate{},
			},
		},
	}

	received, err := DecodeInstruction(&ctx, &m, &at)

	if err != nil {
		t.Fatalf(`Error decoding instruction %v`, err)
	}

	if !reflect.DeepEqual(received, expected) {
		t.Fatalf("Expected \n%+v, \ngot \n%+v", expected, received)
	}
}
