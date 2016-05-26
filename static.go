package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		return ioutil.ReadAll(f)
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/static/test.html": {
		local:   "static/test.html",
		size:    2544,
		modtime: 1463950210,
		compressed: `
H4sIAAAJbogA/6xW3Y7jNBS+LhLvcAgLTbRD3LLcTKctErsrfi52RzAgcYXcxNN4JrGN7UynQvtkXPBI
vALHdn7aJmVXo41G0+ac853v/Lr+9+9/lp+9evvy5vfr11DYqlx/+snSfcJjVQqzigpr1YKQ3W6X7l6k
Um/J/PLykjw6m8gbM5rjJ+CztNyWbP1GalvAtSwZ3DBj4Zpu2ZIEXWNoMs2VBaOzVUQIvaOP6VbKbcmo
4ibNZOVlpOQbQ+7+rJnek6/TWfqieUkrLtI7E62XJHhygZAmkuVG5vuWqJifCQcVwaQxVOvXwjINRlYM
cmrpYklUlxd7tFQzCjxfRVyo2kaQyRLL880sAi13+G0+i9Y3BTeAf96JdWTOE+be4Mcp97LWUAuOqcE9
2x8Sey7PiooI7F6xVeS8RfBAyxpfPFUPjsDYfYnyHc9tsYD5bPbFVQTklPmGVxgerdQ4m23Vx5wf5rug
phh36zRP8fgb0/x2D5qZurQLHB9FhXcYJH4OULTuSQfUWyaYppa19JvaWim6Mn7fqOEHHyIZxv7gYzgD
bwLswZMD5IaJrKiovj8D/q7Xe6hDq2NkJoWlHEPsKpZzo0q6Xwgp2FW07nwcVqeDj5bpeBX7ppA7+kCD
NGrK4J5n8W0tMsulgDiBv1p8r55+3tZ4mqRZybP7UwScPA4TQkOEo46n0+RqxC5V0qCSsPKWfItTvprC
c4/G7wjFOsbJxRDXcvhWvMeuD9Wt7Gi47fNAtRtFWMFPv7x9kyqqDQuokdgP4+i2qokFnaTdIr4P7Hbn
AOdG7Rzk3QVMXTkHxXznBMO+hdH+4K613TAUh/JsP1Ay/bLLuLM4rUGwc8l5E5xYmbNff/7xpawUjraw
8UnyycfrdJi9/+31cEQb0JNKP5lMhtXvtvRcAxA1cTOHZdMWp06wHbzCNYuTdMusGyCfqzObuNa4yZy1
gmchXhQdcbU5tWa5vEYkWrXcceJXPPgY7F/n5w/X8dbJZNrL3T70ipHtCgqfGCupMixH9ngsN/gqpJ4A
cb8Qs6sOHLJ9voJ5L2sSbs4TjBjc8OC6KvfrzvAkzRfgRi0OaNLSJ0nvJJQj7iVHffWC5jP8x9ouwqmO
TVRaqtgd0HRTsjy6sLpmo8faYUO6Ex4dmELu4rCqJ8EMJmrwcnAjcnei5i6E951wt/svAAD//zhqA1nw
CQAA
`,
	},

	"/": {
		isDir: true,
		local: "/",
	},

	"/static": {
		isDir: true,
		local: "/static",
	},
}
