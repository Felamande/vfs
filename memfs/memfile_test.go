package memfs

import (
	"testing"

	"github.com/Felamande/vfs/v2"
)

func TestFileInterface(t *testing.T) {
	_ = vfs.File(NewMemFile("", nil, nil))
}
