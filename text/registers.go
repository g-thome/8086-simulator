package text

import "github.com/g-thome/8086-simulator/registers"

var RegisterIndexToName = map[registers.RegisterIndex]string{
	registers.REGISTER_A:     "a",
	registers.REGISTER_B:     "b",
	registers.REGISTER_C:     "c",
	registers.REGISTER_D:     "d",
	registers.REGISTER_SP:    "sp",
	registers.REGISTER_BP:    "bp",
	registers.REGISTER_SI:    "si",
	registers.REGISTER_DI:    "di",
	registers.REGISTER_ES:    "es",
	registers.REGISTER_CS:    "cs",
	registers.REGISTER_SS:    "ss",
	registers.REGISTER_DS:    "ds",
	registers.REGISTER_IP:    "ip",
	registers.REGISTER_FLAGS: "flags",
}
