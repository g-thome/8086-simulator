package memory

import (
	"os"
)

type Memory struct {
	Bytes [1024 * 1024]byte
}

type SegmentedAccess struct {
	SegmentBase   uint16
	SegmentOffset uint32
}

const MEMORY_ACCESS_MASK = 0xfffff

func LoadMemoryFromFile(fileName string, memory *Memory) int {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	numberOfBytesRead, err := f.Read(memory.Bytes[:])
	if err != nil {
		panic(err)
	}

	return numberOfBytesRead
}

func GetAbsoluteAddressOf(access *SegmentedAccess, additionalOffset uint32) uint32 {
	base := access.SegmentBase
	offset := access.SegmentOffset
	return ((uint32(base) << 4) + uint32(offset+additionalOffset)&MEMORY_ACCESS_MASK)
}

func ReadMemory(m *Memory, absoluteAddress uint32) byte {
	if int(absoluteAddress) > len(m.Bytes) {
		panic("Tried accessing memory out of bounds")
	}

	return m.Bytes[absoluteAddress]
}
