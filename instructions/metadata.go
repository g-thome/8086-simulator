package instructions

import (
	"fmt"
	"github.com/g-thome/8086-simulator/registers"
	"strconv"
)

type Immediate struct {
	Value    int32
	Relative bool
}

type EffectiveAddressBase = uint32

const (
	EFFECTIVE_ADDRESS_DIRECT EffectiveAddressBase = iota

	EFFECTIVE_ADDRESS_BX_SI
	EFFECTIVE_ADDRESS_BX_DI
	EFFECTIVE_ADDRESS_BP_SI
	EFFECTIVE_ADDRESS_BP_DI
	EFFECTIVE_ADDRESS_SI
	EFFECTIVE_ADDRESS_DI
	EFFECTIVE_ADDRESS_BP
	EFFECTIVE_ADDRESS_BX

	EFFECTIVE_ADDRESS_COUNT
)

type EffectiveAddressExpression struct {
	Segment      registers.RegisterIndex
	Base         EffectiveAddressBase
	Displacement int32
}

type InstructionOperand struct {
	Type      OperandType
	Address   EffectiveAddressExpression
	Register  registers.RegisterAccess
	Immediate Immediate
}

type Instruction struct {
	Address  uint32
	Size     uint32 // amount of bytes
	Op       OperationType
	Flags    uint32
	Operands [2]InstructionOperand
}

type InstructionFormat struct {
	Op   OperationType
	Bits []InstructionBits
}

type InstructionBitsUsage = uint8

const (
	BITS_LITERAL InstructionBitsUsage = iota
	BITS_MOD
	BITS_REG
	BITS_RM
	BITS_SR
	BITS_DISP
	BITS_DATA

	BITS_HAS_DISP
	BITS_DISP_ALWAYS_W
	BITS_HAS_DATA
	BITS_W_MAKES_DATA_W
	BITS_RM_REG_ALWAYS_W
	BITS_REL_JMP_DISP
	BITS_D
	BITS_S
	BITS_W
	BITS_V
	BITS_Z

	BITS_COUNT
)

type InstructionFlag = uint32

const (
	INST_LOCK    InstructionFlag = 1 << 0
	INST_REP                     = 1 << 1
	INST_SEGMENT                 = 1 << 2
	INST_WIDE                    = 1 << 3
)

type OperandType = uint8

const (
	OPERAND_NONE OperandType = iota
	OPERAND_REGISTER
	OPERAND_MEMORY
	OPERAND_IMMEDIATE
	OPERAND_RELATIVE_IMMEDIATE
)

type InstructionBits struct {
	Usage    InstructionBitsUsage
	BitCount uint8
	Shift    uint8
	Value    uint8
}

// for debugging purposes
func (f *InstructionFormat) PrintBits() {
	result := ""

	for _, bits := range f.Bits {
		result += strconv.FormatInt(int64(bits.Value), 2)
	}

	fmt.Printf("%s\n", result)
}
