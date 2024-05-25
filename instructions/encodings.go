package instructions

import (
	"log"
	"strconv"
)

var (
	D = InstructionBits{BITS_D, 1, 0, 0}
	S = InstructionBits{BITS_S, 1, 0, 0}
	W = InstructionBits{BITS_W, 1, 0, 0}
	V = InstructionBits{BITS_V, 1, 0, 0}
	Z = InstructionBits{BITS_Z, 1, 0, 0}

	XXX = InstructionBits{BITS_DATA, 3, 0, 0}
	YYY = InstructionBits{BITS_DATA, 3, 3, 0}
	RM  = InstructionBits{BITS_RM, 3, 0, 0}
	MOD = InstructionBits{BITS_MOD, 2, 0, 0}
	REG = InstructionBits{BITS_REG, 3, 0, 0}
	SR  = InstructionBits{BITS_SR, 2, 0, 0}

	DISP      = InstructionBits{BITS_HAS_DISP, 0, 0, 1}
	ADDR      = []InstructionBits{{BITS_HAS_DISP, 0, 0, 1}}
	DATA      = InstructionBits{BITS_HAS_DATA, 0, 0, 1}
	DATA_IF_W = InstructionBits{BITS_W_MAKES_DATA_W, 0, 0, 1}
)

func ImpW(val uint8) InstructionBits {
	return InstructionBits{BITS_W, 0, 0, val}
}

func ImpREG(val uint8) InstructionBits {
	return InstructionBits{BITS_REG, 0, 0, val}
}

func ImpMOD(val uint8) InstructionBits {
	return InstructionBits{BITS_MOD, 0, 0, val}
}

func ImpRM(val uint8) InstructionBits {
	return InstructionBits{BITS_RM, 0, 0, val}
}

func ImpD(val uint8) InstructionBits {
	return InstructionBits{BITS_D, 0, 0, val}
}

func ImpS(val uint8) InstructionBits {
	return InstructionBits{BITS_S, 0, 0, val}
}

var encodings []InstructionEncoding

func b(bits string) InstructionBits {
	i, err := strconv.Atoi(bits)
	if err != nil {
		log.Fatalf("Malformed instructions table. %s is not a valid integer", bits)
	}
	return InstructionBits{BITS_LITERAL, uint8(len(bits)), uint8(i), 0}
}

func flags(f InstructionBitsUsage) InstructionBits {
	return InstructionBits{f, 0, 0, 1}
}

var InstructionFormats = []InstructionFormat{
	{OpMov, []InstructionBits{b("100010"), D, W, MOD, REG, RM}},
	{OpMov, []InstructionBits{b("1100011"), W, MOD, b("000"), RM, DATA, DATA_IF_W, ImpD(0)}},
	{OpMov, []InstructionBits{b("1011"), W, REG, DATA, DATA_IF_W, ImpD(1)}},
	{OpMov, []InstructionBits{b("1010000"), W, ADDR[0], ImpREG(0), ImpMOD(0), ImpRM(0b110), ImpD(1)}},
	{OpMov, []InstructionBits{b("1010001"), W, ADDR[0], ImpREG(0), ImpMOD(0), ImpRM(0b110), ImpD(0)}},
	{OpMov, []InstructionBits{b("100011"), D, b("0"), MOD, b("0"), SR, RM}}, // NOTE: This collapses 2 entries in the 8086 table by adding an explicit D bit

	{OpPush, []InstructionBits{b("11111111"), MOD, b("110"), RM, ImpW(1)}},
	{OpPush, []InstructionBits{b("01010"), REG, ImpW(1)}},
	{OpPush, []InstructionBits{b("000"), SR, b("110"), ImpW(1)}},

	{OpPop, []InstructionBits{b("10001111"), MOD, b("000"), RM, ImpW(1)}},
	{OpPop, []InstructionBits{b("01011"), REG, ImpW(1)}},
	{OpPop, []InstructionBits{b("000"), SR, b("111"), ImpW(1)}},

	{OpXchg, []InstructionBits{b("1000011"), W, MOD, REG, RM, ImpD(1)}},
	{OpXchg, []InstructionBits{b("10010"), REG, ImpMOD(0b11), ImpW(1), ImpRM(0)}},

	{OpIn, []InstructionBits{b("1110010"), W, DATA, ImpREG(0), ImpD(1)}},
	{OpIn, []InstructionBits{b("1110110"), W, ImpREG(0), ImpD(1), ImpMOD(0b11), ImpRM(2), flags(BITS_RM_REG_ALWAYS_W)}},
	{OpOut, []InstructionBits{b("1110011"), W, DATA, ImpREG(0), ImpD(0)}},
	{OpOut, []InstructionBits{b("1110111"), W, ImpREG(0), ImpD(0), ImpMOD(0b11), ImpRM(2), flags(BITS_RM_REG_ALWAYS_W)}},

	{OpXlat, []InstructionBits{b("11010111")}},
	{OpLea, []InstructionBits{b("10001101"), MOD, REG, RM, ImpD(1), ImpW(1)}},
	{OpLds, []InstructionBits{b("11000101"), MOD, REG, RM, ImpD(1), ImpW(1)}},
	{OpLes, []InstructionBits{b("11000100"), MOD, REG, RM, ImpD(1), ImpW(1)}},
	{OpLahf, []InstructionBits{b("10011111")}},
	{OpSahf, []InstructionBits{b("10011110")}},
	{OpPushf, []InstructionBits{b("10011100")}},
	{OpPopf, []InstructionBits{b("10011101")}},

	{OpAdd, []InstructionBits{b("000000"), D, W, MOD, REG, RM}},
	{OpAdd, []InstructionBits{b("100000"), S, W, MOD, b("000"), RM, DATA, DATA_IF_W}},
	{OpAdd, []InstructionBits{b("0000010"), W, DATA, DATA_IF_W, ImpREG(0), ImpD(1)}},

	{OpAdc, []InstructionBits{b("000100"), D, W, MOD, REG, RM}},
	{OpAdc, []InstructionBits{b("100000"), S, W, MOD, b("010"), RM, DATA, DATA_IF_W}},
	{OpAdc, []InstructionBits{b("0001010"), W, DATA, DATA_IF_W, ImpREG(0), ImpD(1)}},

	{OpInc, []InstructionBits{b("1111111"), W, MOD, b("000"), RM}},
	{OpInc, []InstructionBits{b("01000"), REG, ImpW(1)}},

	{OpAaa, []InstructionBits{b("00110111")}},
	{OpDaa, []InstructionBits{b("00100111")}},

	{OpSub, []InstructionBits{b("001010"), D, W, MOD, REG, RM}},
	{OpSub, []InstructionBits{b("100000"), S, W, MOD, b("101"), RM, DATA, DATA_IF_W}},
	{OpSub, []InstructionBits{b("0010110"), W, DATA, DATA_IF_W, ImpREG(0), ImpD(1)}},

	{OpSbb, []InstructionBits{b("000110"), D, W, MOD, REG, RM}},
	{OpSbb, []InstructionBits{b("100000"), S, W, MOD, b("011"), RM, DATA, DATA_IF_W}},
	{OpSbb, []InstructionBits{b("0001110"), W, DATA, DATA_IF_W, ImpREG(0), ImpD(1)}},

	{OpDec, []InstructionBits{b("1111111"), W, MOD, b("001"), RM}},
	{OpDec, []InstructionBits{b("01001"), REG, ImpW(1)}},

	{OpNeg, []InstructionBits{b("1111011"), W, MOD, b("011"), RM}},

	{OpCmp, []InstructionBits{b("001110"), D, W, MOD, REG, RM}},
	{OpCmp, []InstructionBits{b("100000"), S, W, MOD, b("111"), RM, DATA, DATA_IF_W}},
	{OpCmp, []InstructionBits{b("0011110"), W, DATA, DATA_IF_W, ImpREG(0), ImpD(1)}}, // TODO: The manual table suggests this data is only 8-bit, but wouldn't it be 16 as well?

	{OpAas, []InstructionBits{b("00111111")}},
	{OpDas, []InstructionBits{b("00101111")}},
	{OpMul, []InstructionBits{b("1111011"), W, MOD, b("100"), RM, ImpS(0)}},
	{OpImul, []InstructionBits{b("1111011"), W, MOD, b("101"), RM, ImpS(1)}},
	{OpAam, []InstructionBits{b("11010100"), b("00001010")}}, // TODO: The manual says this has a DISP... but how could it? What for??
	{OpDiv, []InstructionBits{b("1111011"), W, MOD, b("110"), RM, ImpS(0)}},
	{OpIdiv, []InstructionBits{b("1111011"), W, MOD, b("111"), RM, ImpS(1)}},
	{OpAad, []InstructionBits{b("11010101"), b("00001010")}},
	{OpCbw, []InstructionBits{b("10011000")}},
	{OpCwd, []InstructionBits{b("10011001")}},

	{OpNot, []InstructionBits{b("1111011"), W, MOD, b("010"), RM}},
	{OpShl, []InstructionBits{b("110100"), V, W, MOD, b("100"), RM}},
	{OpShr, []InstructionBits{b("110100"), V, W, MOD, b("101"), RM}},
	{OpSar, []InstructionBits{b("110100"), V, W, MOD, b("111"), RM}},
	{OpRol, []InstructionBits{b("110100"), V, W, MOD, b("000"), RM}},
	{OpRor, []InstructionBits{b("110100"), V, W, MOD, b("001"), RM}},
	{OpRcl, []InstructionBits{b("110100"), V, W, MOD, b("010"), RM}},
	{OpRcr, []InstructionBits{b("110100"), V, W, MOD, b("011"), RM}},

	{OpAnd, []InstructionBits{b("001000"), D, W, MOD, REG, RM}},
	{OpAnd, []InstructionBits{b("1000000"), W, MOD, b("100"), RM, DATA, DATA_IF_W}},
	{OpAnd, []InstructionBits{b("0010010"), W, DATA, DATA_IF_W, ImpREG(0), ImpD(1)}},

	{OpTest, []InstructionBits{b("100001"), D, W, MOD, REG, RM}},
	{OpTest, []InstructionBits{b("1111011"), W, MOD, b("000"), RM, DATA, DATA_IF_W}},
	{OpTest, []InstructionBits{b("1010100"), W, DATA, DATA_IF_W, ImpREG(0), ImpD(1)}}, // TODO: The manual table suggests this data is only 8-bit, but it seems like it could be 16 too?

	{OpOr, []InstructionBits{b("000010"), D, W, MOD, REG, RM}},
	{OpOr, []InstructionBits{b("1000000"), W, MOD, b("001"), RM, DATA, DATA_IF_W}},
	{OpOr, []InstructionBits{b("0000110"), W, DATA, DATA_IF_W, ImpREG(0), ImpD(1)}},

	{OpXor, []InstructionBits{b("001100"), D, W, MOD, REG, RM}},
	{OpXor, []InstructionBits{b("1000000"), W, MOD, b("110"), RM, DATA, DATA_IF_W}}, // TODO: The manual has conflicting information about this encoding, but I believe this is the correct binary pattern.
	{OpXor, []InstructionBits{b("0011010"), W, DATA, DATA_IF_W, ImpREG(0), ImpD(1)}},

	{OpRep, []InstructionBits{b("1111001"), Z}},
	{OpMovs, []InstructionBits{b("1010010"), W}},
	{OpCmps, []InstructionBits{b("1010011"), W}},
	{OpScas, []InstructionBits{b("1010111"), W}},
	{OpLods, []InstructionBits{b("1010110"), W}},
	{OpStos, []InstructionBits{b("1010101"), W}},

	{OpCall, []InstructionBits{b("11101000"), ADDR[0]}},
	{OpCall, []InstructionBits{b("11111111"), MOD, b("010"), RM, ImpW(1)}},
	{OpCall, []InstructionBits{b("10011010"), ADDR[0], DATA, ImpW(1)}},
	{OpCall, []InstructionBits{b("11111111"), MOD, b("011"), RM, ImpW(1)}},

	{OpJmp, []InstructionBits{b("11101001"), ADDR[0]}},
	{OpJmp, []InstructionBits{b("11101011"), DISP}},
	{OpJmp, []InstructionBits{b("11111111"), MOD, b("100"), RM, ImpW(1)}},
	{OpJmp, []InstructionBits{b("11101010"), ADDR[0], DATA, ImpW(1)}},
	{OpJmp, []InstructionBits{b("11111111"), MOD, b("101"), RM, ImpW(1)}},

	{OpRet, []InstructionBits{b("11000011")}},
	{OpRet, []InstructionBits{b("11000010"), DATA, DATA_IF_W, ImpW(1)}},
	{OpRet, []InstructionBits{b("11001011")}},
	{OpRet, []InstructionBits{b("11001010"), DATA, DATA_IF_W, ImpW(1)}},

	{OpJe, []InstructionBits{b("01110100"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJl, []InstructionBits{b("01111100"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJle, []InstructionBits{b("01111110"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJb, []InstructionBits{b("01110010"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJbe, []InstructionBits{b("01110110"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJp, []InstructionBits{b("01111010"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJo, []InstructionBits{b("01110000"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJs, []InstructionBits{b("01111000"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJne, []InstructionBits{b("01110101"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJnl, []InstructionBits{b("01111101"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJg, []InstructionBits{b("01111111"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJnb, []InstructionBits{b("01110011"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJa, []InstructionBits{b("01110111"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJnp, []InstructionBits{b("01111011"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJno, []InstructionBits{b("01110001"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJns, []InstructionBits{b("01111001"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpLoop, []InstructionBits{b("11100010"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpLoopz, []InstructionBits{b("11100001"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpLoopnz, []InstructionBits{b("11100000"), DISP, flags(BITS_REL_JMP_DISP)}},
	{OpJcxz, []InstructionBits{b("11100011"), DISP, flags(BITS_REL_JMP_DISP)}},

	{OpInt, []InstructionBits{b("11001101"), DATA}},
	{OpInt3, []InstructionBits{b("11001100")}}, // TODO: The manual does not suggest that this intrinsic has an "int3" mnemonic, but NASM thinks so

	{OpInto, []InstructionBits{b("11001110")}},
	{OpIret, []InstructionBits{b("11001111")}},

	{OpClc, []InstructionBits{b("11111000")}},
	{OpCmc, []InstructionBits{b("11110101")}},
	{OpStc, []InstructionBits{b("11111001")}},
	{OpCld, []InstructionBits{b("11111100")}},
	{OpStd, []InstructionBits{b("11111101")}},
	{OpCli, []InstructionBits{b("11111010")}},
	{OpSti, []InstructionBits{b("11111011")}},
	{OpHlt, []InstructionBits{b("11110100")}},
	{OpWait, []InstructionBits{b("10011011")}},
	{OpEsc, []InstructionBits{b("11011"), XXX, MOD, YYY, RM}},
	{OpLock, []InstructionBits{b("11110000")}},
	{OpSegment, []InstructionBits{b("001"), SR, b("110")}},
}
