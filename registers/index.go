package registers

type RegisterIndex uint32

const (
	REGISTER_NONE RegisterIndex = iota
	REGISTER_A
	REGISTER_B
	REGISTER_C
	REGISTER_D
	REGISTER_SP
	REGISTER_BP
	REGISTER_SI
	REGISTER_DI
	REGISTER_ES
	REGISTER_CS
	REGISTER_SS
	REGISTER_DS
	REGISTER_IP
	REGISTER_FLAGS
	REGISTER_COUNT
)

func (r RegisterIndex) String() string {
	switch r {
	case REGISTER_A:
		return "A"
	case REGISTER_B:
		return "B"
	case REGISTER_C:
		return "C"
	case REGISTER_D:
		return "D"
	case REGISTER_SP:
		return "SP"
	case REGISTER_BP:
		return "BP"
	case REGISTER_SI:
		return "SI"
	case REGISTER_DI:
		return "DI"
	case REGISTER_ES:
		return "ES"
	case REGISTER_CS:
		return "CS"
	case REGISTER_SS:
		return "SS"
	case REGISTER_DS:
		return "DS"
	case REGISTER_IP:
		return "IP"
	case REGISTER_FLAGS:
		return "FLAGS"
	default:
		return ""
	}
}

type RegisterAccess struct {
	Index  RegisterIndex
	Offset uint8
	Count  uint8
}
