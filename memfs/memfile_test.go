package memfs

import (
	"testing"

	"github.com/Felamande/vfs"
)

func TestFileInterface(t *testing.T) {
	_ = vfs.File(NewMemFile("", nil, nil))
}
