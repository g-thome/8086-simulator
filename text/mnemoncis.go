package text

import "github.com/g-thome/8086-simulator/instructions"

var OperationTypeToMnemonic = map[instructions.OperationType]string{
	instructions.OpNone:    "none",
	instructions.OpMov:     "mov",
	instructions.OpPush:    "push",
	instructions.OpPop:     "pop",
	instructions.OpXchg:    "xchg",
	instructions.OpIn:      "in",
	instructions.OpOut:     "out",
	instructions.OpXlat:    "xlat",
	instructions.OpLea:     "lea",
	instructions.OpLds:     "lds",
	instructions.OpLes:     "les",
	instructions.OpLahf:    "lahf",
	instructions.OpSahf:    "sahf",
	instructions.OpPushf:   "pushf",
	instructions.OpPopf:    "popf",
	instructions.OpAdd:     "add",
	instructions.OpAdc:     "adc",
	instructions.OpInc:     "inc",
	instructions.OpAaa:     "aaa",
	instructions.OpDaa:     "daa",
	instructions.OpSub:     "sub",
	instructions.OpSbb:     "sbb",
	instructions.OpDec:     "dec",
	instructions.OpNeg:     "neg",
	instructions.OpCmp:     "cmp",
	instructions.OpAas:     "aas",
	instructions.OpDas:     "das",
	instructions.OpMul:     "mul",
	instructions.OpImul:    "imul",
	instructions.OpAam:     "aam",
	instructions.OpDiv:     "div",
	instructions.OpIdiv:    "idiv",
	instructions.OpAad:     "aad",
	instructions.OpCbw:     "cbw",
	instructions.OpCwd:     "cwd",
	instructions.OpNot:     "not",
	instructions.OpShl:     "shl",
	instructions.OpShr:     "shr",
	instructions.OpSar:     "sar",
	instructions.OpRol:     "rol",
	instructions.OpRor:     "ror",
	instructions.OpRcl:     "rcl",
	instructions.OpRcr:     "rcr",
	instructions.OpAnd:     "and",
	instructions.OpTest:    "test",
	instructions.OpOr:      "or",
	instructions.OpXor:     "xor",
	instructions.OpRep:     "rep",
	instructions.OpMovs:    "movs",
	instructions.OpCmps:    "cmps",
	instructions.OpScas:    "scas",
	instructions.OpLods:    "lods",
	instructions.OpStos:    "stos",
	instructions.OpCall:    "call",
	instructions.OpJmp:     "jmp",
	instructions.OpRet:     "ret",
	instructions.OpRetf:    "retf",
	instructions.OpJe:      "je",
	instructions.OpJl:      "jl",
	instructions.OpJle:     "jle",
	instructions.OpJb:      "jb",
	instructions.OpJbe:     "jbe",
	instructions.OpJp:      "jp",
	instructions.OpJo:      "jo",
	instructions.OpJs:      "js",
	instructions.OpJne:     "jne",
	instructions.OpJnl:     "jnl",
	instructions.OpJg:      "jg",
	instructions.OpJnb:     "jnb",
	instructions.OpJa:      "ja",
	instructions.OpJnp:     "jnp",
	instructions.OpJno:     "jno",
	instructions.OpJns:     "jns",
	instructions.OpLoop:    "loop",
	instructions.OpLoopz:   "loopz",
	instructions.OpLoopnz:  "loopnz",
	instructions.OpJcxz:    "jcxz",
	instructions.OpInt:     "int",
	instructions.OpInt3:    "int3",
	instructions.OpInto:    "into",
	instructions.OpIret:    "iret",
	instructions.OpClc:     "clc",
	instructions.OpCmc:     "cmc",
	instructions.OpStc:     "stc",
	instructions.OpCld:     "cld",
	instructions.OpStd:     "std",
	instructions.OpCli:     "cli",
	instructions.OpSti:     "sti",
	instructions.OpHlt:     "hlt",
	instructions.OpWait:    "wait",
	instructions.OpEsc:     "esc",
	instructions.OpLock:    "lock",
	instructions.OpSegment: "segment",
	instructions.OpCount:   "count",
}
