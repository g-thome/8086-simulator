package simulator

import (
	"github.com/g-thome/8086-simulator/decode"
	"github.com/g-thome/8086-simulator/memory"
	"github.com/g-thome/8086-simulator/registers"
	"testing"
)

func TestMovImmediateToRegister(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := decode.DefaultDisAsmContext()
	m := memory.Memory{}
	bytesCount := memory.LoadMemoryFromFile("../fixtures/bin/mov_cx_3", &m)
	s := NewSimulation()

	expected := [registers.REGISTER_COUNT]Register{
		{0}, {0}, {0}, {3}, {0}, {0}, {0}, {0}, {0}, {0},
	}

	count := bytesCount
	for count > 0 {
		instruction, _ := decode.DecodeInstruction(&ctx, &m, &at)
		count -= int(instruction.Size)

		decode.UpdateContext(&ctx, instruction)
		s.Run(instruction)
	}

	if s.Registers != expected {
		t.Fatalf("Wrong end state in registers. Expected \n%v\n, got \n%v", expected, s.Registers)
	}
}

func TestMovRegisterToRegister(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := decode.DefaultDisAsmContext()
	m := memory.Memory{}
	bytesCount := memory.LoadMemoryFromFile("../fixtures/bin/mov_reg_reg", &m)
	s := NewSimulation()

	expected := [registers.REGISTER_COUNT]Register{
		{0}, {0}, {13}, {13}, {0}, {0}, {0}, {0}, {0}, {0},
	}

	count := bytesCount
	for count > 0 {
		instruction, _ := decode.DecodeInstruction(&ctx, &m, &at)
		count -= int(instruction.Size)

		decode.UpdateContext(&ctx, instruction)
		s.Run(instruction)
	}

	if s.Registers != expected {
		t.Fatalf("Wrong end state in registers. Expected \n%v\n, got \n%v", expected, s.Registers)
	}
}

func TestListing0043(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := decode.DefaultDisAsmContext()
	m := memory.Memory{}
	bytesCount := memory.LoadMemoryFromFile("../fixtures/bin/listings/listing_0043_immediate_movs", &m)
	s := NewSimulation()

	expected := [registers.REGISTER_COUNT]Register{
		{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {0},
	}

	count := bytesCount
	for count > 0 {
		instruction, _ := decode.DecodeInstruction(&ctx, &m, &at)
		count -= int(instruction.Size)

		decode.UpdateContext(&ctx, instruction)
		s.Run(instruction)
	}

	if s.Registers != expected {
		t.Fatalf("Wrong end state in registers. Expected \n%v\n, got \n%v", expected, s.Registers)
	}
}

func TestListing0044(t *testing.T) {
	at := memory.SegmentedAccess{0, 0}
	ctx := decode.DefaultDisAsmContext()
	m := memory.Memory{}
	bytesCount := memory.LoadMemoryFromFile("../fixtures/bin/listings/listing_0044_register_movs", &m)
	s := NewSimulation()

	expected := [registers.REGISTER_COUNT]Register{
		{0}, {4}, {3}, {2}, {1}, {1}, {2}, {3}, {4}, {0},
	}

	count := bytesCount
	for count > 0 {
		instruction, _ := decode.DecodeInstruction(&ctx, &m, &at)
		count -= int(instruction.Size)

		decode.UpdateContext(&ctx, instruction)
		s.Run(instruction)
	}

	if s.Registers != expected {
		t.Fatalf("Wrong end state in registers. Expected \n%v\n, got \n%v", expected, s.Registers)
	}
}
