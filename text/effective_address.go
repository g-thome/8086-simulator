package text

import "github.com/g-thome/8086-simulator/instructions"

var EffectiveAddressBaseToText = map[instructions.EffectiveAddressBase]string{
	instructions.EFFECTIVE_ADDRESS_DIRECT: "direct",

	instructions.EFFECTIVE_ADDRESS_BX_SI: "bx+si",
	instructions.EFFECTIVE_ADDRESS_BX_DI: "bx+di",
	instructions.EFFECTIVE_ADDRESS_BP_SI: "bp+si",
	instructions.EFFECTIVE_ADDRESS_BP_DI: "bp+di",
	instructions.EFFECTIVE_ADDRESS_SI:    "si",
	instructions.EFFECTIVE_ADDRESS_DI:    "di",
	instructions.EFFECTIVE_ADDRESS_BP:    "bp",
	instructions.EFFECTIVE_ADDRESS_BX:    "bx",

	instructions.EFFECTIVE_ADDRESS_COUNT: "count",
}
