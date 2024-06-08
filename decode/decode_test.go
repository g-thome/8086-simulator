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
				UnsignedImmediate: 0,
				SignedImmediate:   0,
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
				UnsignedImmediate: 0,
				SignedImmediate:   0,
			},
		},
	}

	received, err := DecodeInstruction(&ctx, &m, &at)

	if err != nil {
		t.Fatalf(`Error decoding instruction %v`, err)
	}

	if received.Address != expected.Address {
		t.Fatalf("Wrong instruction address. Expected %v, got %v", expected.Address, received.Address)
	}

	if received.Size != expected.Size {
		t.Fatalf("Wrong instruction size. Expected %v, got %v", expected.Size, received.Size)
	}

	if received.Op != expected.Op {
		t.Fatalf("Wrong OperationType. Expected %v, got %v", expected.Op, received.Op)
	}

	if received.Flags != expected.Flags {
		t.Fatalf("Wrong flags. Expected %v, got %v", expected.Flags, received.Flags)
	}

	for i, _ := range received.Operands {
		expectedOper := expected.Operands[i]
		receivedOper := received.Operands[i]
		if receivedOper.Type != expectedOper.Type {
			t.Fatalf("Wrong operand type for operand %d. Expected %v, got %v", i+1, expectedOper.Type, receivedOper.Type)
		}

		if !reflect.DeepEqual(receivedOper.Address, expectedOper.Address) {
			t.Fatalf("Operand Address mismatch for operand %d. Expected %v, got %v", i+1, expectedOper.Address, receivedOper.Address)
		}

		if !reflect.DeepEqual(receivedOper.Register, expectedOper.Register) {
			t.Fatalf("Operand Register mismatch for operand %d. Expected %v, got %v", i+1, expectedOper.Register, receivedOper.Register)
		}

		if receivedOper.UnsignedImmediate != expectedOper.UnsignedImmediate {
			t.Fatalf("Operand unsigned immediate mismatch for operand %d. Expected %v, got %v", i+1, expectedOper.UnsignedImmediate, receivedOper.UnsignedImmediate)
		}

		if receivedOper.SignedImmediate != expectedOper.SignedImmediate {
			t.Fatalf("Operand signed immediate mismatch for operand %d. Expected %v, got %v", i+1, expectedOper.SignedImmediate, receivedOper.SignedImmediate)
		}
	}
}

func TestDecodeMovRegImmediate(t *testing.T) {
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
				UnsignedImmediate: 0,
				SignedImmediate:   0,
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
				UnsignedImmediate: 3,
				SignedImmediate:   0,
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
				UnsignedImmediate: 0,
				SignedImmediate:   0,
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
				UnsignedImmediate: 0,
				SignedImmediate:   0,
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
				UnsignedImmediate: 0,
				SignedImmediate:   0,
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
				UnsignedImmediate: 0,
				SignedImmediate:   0,
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
