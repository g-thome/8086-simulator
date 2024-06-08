package memory

import (
	"reflect"
	"testing"
)

func TestLoadMemoryFromFile(t *testing.T) {
	m := Memory{}

	bytesRead := LoadMemoryFromFile("../fixtures/bin/mov_cx_bx", &m)

	expectedBytesRead := 2
	if bytesRead != expectedBytesRead {
		t.Fatalf(`bytesRead mismatch. Expected %d got %d`, expectedBytesRead, bytesRead)
	}

	expectedBytes := []byte{137, 217}
	actualBytes := m.Bytes[:2]
	if !reflect.DeepEqual(expectedBytes, actualBytes) {
		t.Fatalf(`bytes mismatch. Expected %v got %v`, expectedBytes, actualBytes)
	}
}

func TestReadMemory(t *testing.T) {
	m := Memory{}

	var expected byte = 0b00010001
	m.Bytes[4] = expected

	received := ReadMemory(&m, 4)

	if expected != received {
		t.Fatalf(`bytes mismatch. Expected %v got %v`, expected, received)
	}
}
