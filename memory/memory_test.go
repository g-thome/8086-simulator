package memory

import (
	"reflect"
	"testing"
)

func TestLoadMemoryFromFile(t *testing.T) {
	m := Memory{}

	bytesRead := LoadMemoryFromFile("../fixtures/single_register_mov", &m)

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
