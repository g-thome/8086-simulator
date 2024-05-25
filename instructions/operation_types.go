package instructions

type OperationType uint32

const (
	OpNone OperationType = iota
	OpMov
	OpPush
	OpPop
	OpXchg
	OpIn
	OpOut
	OpXlat
	OpLea
	OpLds
	OpLes
	OpLahf
	OpSahf
	OpPushf
	OpPopf
	OpAdd
	OpAdc
	OpInc
	OpAaa
	OpDaa
	OpSub
	OpSbb
	OpDec
	OpNeg
	OpCmp
	OpAas
	OpDas
	OpMul
	OpImul
	OpAam
	OpDiv
	OpIdiv
	OpAad
	OpCbw
	OpCwd
	OpNot
	OpShl
	OpShr
	OpSar
	OpRol
	OpRor
	OpRcl
	OpRcr
	OpAnd
	OpTest
	OpOr
	OpXor
	OpRep
	OpMovs
	OpCmps
	OpScas
	OpLods
	OpStos
	OpCall
	OpJmp
	OpRet
	OpRetf
	OpJe
	OpJl
	OpJle
	OpJb
	OpJbe
	OpJp
	OpJo
	OpJs
	OpJne
	OpJnl
	OpJg
	OpJnb
	OpJa
	OpJnp
	OpJno
	OpJns
	OpLoop
	OpLoopz
	OpLoopnz
	OpJcxz
	OpInt
	OpInt3
	OpInto
	OpIret
	OpClc
	OpCmc
	OpStc
	OpCld
	OpStd
	OpCli
	OpSti
	OpHlt
	OpWait
	OpEsc
	OpLock
	OpSegment
	OpCount
)
