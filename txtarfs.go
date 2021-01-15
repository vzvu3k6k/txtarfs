package txtarfs

import (
	"bytes"
	"io/fs"
	"time"

	"github.com/rogpeppe/go-internal/txtar"
)

type Txtarfs struct {
	files map[string]([]byte)
}

func (tfs Txtarfs) Open(name string) (fs.File, error) {
	content := tfs.files[name]
	return txtarFile{
		name:    name,
		content: bytes.NewBuffer(content),
		size:    int64(len(content)),
	}, nil
}

func Parse(data []byte) Txtarfs {
	a := txtar.Parse(data)
	tfs := Txtarfs{files: make(map[string][]byte)}
	for _, f := range a.Files {
		tfs.files[f.Name] = f.Data
	}
	return tfs
}

type txtarFile struct {
	name    string
	content *bytes.Buffer
	size    int64
}

func (tf txtarFile) Read(b []byte) (int, error) {
	return tf.content.Read(b)
}

func (tf txtarFile) Stat() (fs.FileInfo, error) {
	return txtarFileInfo{name: tf.name, size: tf.size}, nil
}

func (tf txtarFile) Close() error { return nil }

type txtarFileInfo struct {
	name string
	size int64
}

func (i txtarFileInfo) Name() string       { return i.name }
func (i txtarFileInfo) Size() int64        { return i.size }
func (i txtarFileInfo) Mode() fs.FileMode  { return fs.ModePerm }
func (i txtarFileInfo) ModTime() time.Time { return time.Unix(0, 0) }
func (i txtarFileInfo) IsDir() bool        { return false }
func (i txtarFileInfo) Sys() interface{}   { return func() {} }
