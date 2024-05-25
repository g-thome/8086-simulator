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
