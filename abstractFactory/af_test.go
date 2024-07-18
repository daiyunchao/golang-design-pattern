package abstractFactory

import (
	"testing"
)

func TestCloseFile(t *testing.T) {
	opt := NewFileA()
	OpenFile(opt)
	CloseFile(opt)
	opt = NewFileB()
	OpenFile(opt)
	CloseFile(opt)
}
