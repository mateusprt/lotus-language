package vm

import "testing"

func TestWriteChunk(t *testing.T) {
	var chunk Chunk
	WriteChunk(&chunk, OP_CONSTANT, 1)

	if len(chunk.Code) != 1 {
		t.Errorf("Expected chunk.Code length to be 1, got %d", len(chunk.Code))
	}
	if chunk.Code[0] != OP_CONSTANT {
		t.Errorf("Expected first instruction to be OP_CONSTANT, got %d", chunk.Code[0])
	}

	if len(chunk.lines) != 1 || chunk.lines[0] != 1 {
		t.Errorf("expected lines[0] to be 1, got %d", chunk.lines[0])
	}
}

func TestAddConstant(t *testing.T) {
	var chunk Chunk
	idx := AddConstant(&chunk, 3.14)

	if idx != 0 {
		t.Errorf("expected index 0, got %d", idx)
	}
	if len(chunk.Constants) != 1 {
		t.Errorf("expected Constants length 1, got %d", len(chunk.Constants))
	}
	if chunk.Constants[0] != 3.14 {
		t.Errorf("expected Constants[0] to be 3.14, got %v", chunk.Constants[0])
	}
}
