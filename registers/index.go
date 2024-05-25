package registers

type RegisterIndex = uint32

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

type RegisterAccess struct {
	Index  RegisterIndex
	Offset uint8
	Count  uint8
}
