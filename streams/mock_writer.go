package streams

import (
	"errors"
	"io"
)

type MockWriter struct {
	io.Writer

	writtenOperations [][]byte
}

// Write writes to the mock writer and stores it for later use.
func (writer *MockWriter) Write(bytes []byte) {
	writer.writtenOperations = append(writer.writtenOperations, bytes)
}

// Pop removes the first operation written and returns it back to the user.
func (writer *MockWriter) Pop() ([]byte, error) {
	if len(writer.writtenOperations) == 0 {
		return nil, errors.New("no more operations to pop")
	}

	operation := writer.writtenOperations[0]
	writer.writtenOperations = writer.writtenOperations[:1]

	return operation, nil
}

// AssertEmpty makes sure that after testing, there's no more messages in the buffer.
func (writer *MockWriter) AssertEmpty() error {
	if len(writer.writtenOperations) != 0 {
		return errors.New("mock buffer was not empty")
	}
	return nil
}