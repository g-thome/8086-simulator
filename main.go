package main

import (
	"fmt"
	"os"

	"github.com/g-thome/8086-simulator/decode"
	"github.com/g-thome/8086-simulator/memory"
	"github.com/g-thome/8086-simulator/text"
)

func disAsm8086(m *memory.Memory, disAsmByteCount uint32, disAsmStart memory.SegmentedAccess) {
	at := disAsmStart

	ctx := decode.DefaultDisAsmContext()

	count := disAsmByteCount

	for count > 0 {
		instruction, err := decode.DecodeInstruction(&ctx, m, &at)
		if err != nil {
			panic(err)
		}

		if count < instruction.Size {
			panic("Instruction extends outside disassembly region")
		}

		count -= instruction.Size

		decode.UpdateContext(&ctx, instruction)

		if text.IsPrintable(instruction) {
			fmt.Println(text.PrintInstruction(instruction))
		}
	}
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "" {
		fmt.Println("USAGE: ", os.Args[0], " <name of the binary file>")
		return
	}

	fileName := os.Args[1]

	ram := memory.Memory{}

	bytesRead := memory.LoadMemoryFromFile(fileName, &ram)

	disAsm8086(&ram, uint32(bytesRead), memory.SegmentedAccess{0, 0})
}
