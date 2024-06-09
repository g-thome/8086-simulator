package text

import (
	"github.com/g-thome/8086-simulator/instructions"
	"github.com/g-thome/8086-simulator/registers"
	"strconv"
)

func Mnemonic(op instructions.OperationType) string {
	return OperationTypeToMnemonic[op]
}

func RegName(reg registers.RegisterAccess) string {
	var Names = [][3]string{
		{"", "", ""},
		{"al", "ah", "ax"},
		{"bl", "bh", "bx"},
		{"cl", "ch", "cx"},
		{"dl", "dh", "dx"},
		{"sp", "sp", "sp"},
		{"bp", "bp", "bp"},
		{"si", "si", "si"},
		{"di", "di", "di"},
		{"es", "es", "es"},
		{"cs", "cs", "cs"},
		{"ss", "ss", "ss"},
		{"ds", "ds", "ds"},
		{"ip", "ip", "ip"},
		{"flags", "flags", "flags"},
	}

	result := Names[reg.Index][2]
	if reg.Count != 2 {
		result = Names[reg.Index][reg.Offset&1]
	}

	return result
}

func EffectiveAddressBase(address instructions.EffectiveAddressBase) string {
	return EffectiveAddressBaseToText[address]
}

func IsPrintable(inst instructions.Instruction) bool {
	return inst.Op != instructions.OpLock &&
		inst.Op != instructions.OpRep &&
		inst.Op != instructions.OpSegment
}

func PrintInstruction(inst instructions.Instruction) string {
	result := ""

	flags := inst.Flags
	w := flags & instructions.INST_WIDE

	if (flags & instructions.INST_LOCK) > 0 {
		if inst.Op == instructions.OpXchg {
			tmp := inst.Operands[0]
			inst.Operands[0] = inst.Operands[1]
			inst.Operands[1] = tmp
		}

		result += "lock "
	}

	mnemonicSuffix := ""

	if (flags & instructions.INST_REP) > 0 {
		result += "rep "
		if w > 0 {
			mnemonicSuffix = "w"
		} else {
			mnemonicSuffix = "b"
		}
	}

	result += Mnemonic(inst.Op) + mnemonicSuffix + " "

	separator := ""

	for _, o := range inst.Operands {
		if o.Type != instructions.OPERAND_NONE {
			result += separator
			separator = ", "

			switch o.Type {
			case instructions.OPERAND_NONE:
				break
			case instructions.OPERAND_REGISTER:
				result += RegName(o.Register)
			case instructions.OPERAND_MEMORY:
				if inst.Operands[0].Type != instructions.OPERAND_REGISTER {
					if w > 0 {
						result += "word"
					} else {
						result += "byte"
					}
				}

				if (flags & instructions.INST_SEGMENT) > 0 {
					result += RegName(registers.RegisterAccess{o.Address.Segment, 0, 2})
				}

				result += "[" + EffectiveAddressBase(o.Address.Base)
				if o.Address.Displacement != 0 {
					result += strconv.Itoa(int(o.Address.Displacement))
				}

				result += "]"
			case instructions.OPERAND_IMMEDIATE:
				if o.Immediate.Relative {
					result += strconv.Itoa(int(o.Immediate.Value + int32(inst.Size)))
				} else {
					result += strconv.Itoa(int(o.Immediate.Value))
				}
			}
		}
	}

	return result
}
