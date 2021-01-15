package txtarfs

import (
	"io/fs"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var txtArchive = []byte(`-- foo/foo.go --
package foo
-- main.go --
package main`)

func TestTxtarfs(t *testing.T) {
	rootFS := Parse(txtArchive)

	gotBody, err := fs.ReadFile(rootFS, "foo/foo.go")
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff([]byte("package foo\n"), gotBody); diff != "" {
		t.Fatalf("foo/foo.go mismatch %s", diff)
	}
}
