// Code generated by "esc -o ./views/fs.go -ignore fs.go -pkg views -prefix views ./views"; DO NOT EDIT.

package views

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
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
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
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
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
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

	"/create.html": {
		name:    "create.html",
		local:   "views/create.html",
		size:    2091,
		modtime: 1559804833,
		compressed: `
H4sIAAAAAAAC/8xVTWvcMBC951cMurQ92JvkvDaElkAhlJCQHzBrzdoCfSGN8sGy/73I9iaO2aRxKSV7
2dHOm/dGb7TSbidpqyyBMKis2O9PAADWUt1DozHGSjTOMlkWdZ/ps935LFl0hBJULBqyTEHUux2Uv9AQ
7PfrVXc+qZ0w+xSoaCfEc4AuNu6x0KGFHprGr+JsVtPXbV0wgA0rZyuxagIhkwBD3DlZCe8ii1fCPf45
KlCr1pIUIJGx2Ch+wCDJFg/ITadsW4ljqoOyIi0jcX1yNH9029m44HTRBpf8G8TPxRo3pGHrQiW2hJwC
ifpyCCC7vF71iD+wKOsTg5IvJGDRUCW2Gtsyh+JdgvHDT54qwfTIH4J7jQ11Tkt66R6yYi9eluU7m1+v
pLr/H656CvnkYkuivnFau8Rw/fzbUnsnbFOHw0BcRg7I1D6VE9zHjQ9oM69RthKnAgw+VuLs9PRDDPeo
E+WyT2B5JJKiviWS4LbAHUFAK50BHylJd1gpO+QG75ZOohcZjXs1g8yetW9R8wLzbTIbCn9x7vN5H7UP
O8u9lWUJX53Pdxbqb59hLJIaFZWzF17d3VyJ+se4hovrn3B3c7V0AjO+d/8PB+yVa1Vz4dWCwaSgl0+F
OxX7q+hLhCZFdgYOLYDOPRw2/Y9mtEnMzo4dx7QxavYkjYBJXPigDIYnUX/vHzS41NiuV0PyDa316uVB
OpJzwcxe3KHvI8sx3O3Iyv3+dwAAAP//emMeZisIAAA=
`,
	},

	"/edit.html": {
		name:    "edit.html",
		local:   "views/edit.html",
		size:    4226,
		modtime: 1559908989,
		compressed: `
H4sIAAAAAAAC/8RW3WrrOBC+z1MMutndA457znVsKKSFQFlKSx9AsSa2WFkyklxSjN99kWLnOK7jnzbs
+iZ2NPpGmm/mm6kqhgcuEUhOuSR1vQIA2DD+DomgxkQkUdKitCT2K341+9VbDDKkDLgJEpQWNYmrCtaP
gqbrv2mOUNebMPvVAejAF6XGIO2g9w1EsFfHQOgUvGnZ/AQ/e3v8voPSOdDEciUjEiLjlkCONlMsIoUy
lly49dbnt4AKnkpkA7gnbI6CGbTxanB98GIuPlqJINWqLK4AnzcLukcBB6UjUqB2oaQpkvhFCaFKC8/n
/zahN52A47IoLXB2gTa6pXkkzTEiB0efPvleG6upxfRjvRDKfhQYEU1ligRyLiNyRyCnx4j8vLubhfBO
RYkROSdUE431a3ui32GBuh4J8SZk/P2/4M6gy6FXRAbqADZD0FQylUNhsGSq/eLytHa6z1JKvZMvkukO
4I73SoVdwKIs8z3qb5H20vE8zpZ7qir8sTEFlZ+KNsjRGJpiwKXgEkm8CZ1d/CNsBOx/TQCGCTdcyfuC
v708kXjbfMP98w7eXp6Wct3D+24Jt3BPKuXJfcEXpECpxSzrqgJ+uFaw2wv/zzuo68kiH9hDqgolG+G7
8xSCJpgpwVBHxGbcgAvNHwaS0liVQxsSEA6/5Qn+VIXrJFT89R1dORHZYcS9kiaiGWcMJflUMU3fHPP7
CZezKdTddgJzX1qrZINiyn3Oey2zMei8B4XmOdUfJH4rGLW4CU//XwnJJvzdRwfWlM57o0AvvM13f2l8
Yqkq8D0I1s9cKGuglzanuaGvMoMDxvgQcEFK4XyNszJGyAAWHu1u28JZPNoLsIejRS2pWAjqcmcE9RGp
LTU+eqsT8GpKAVE61clR2vXULa/d1CPM6HFOIAfdrWYpWpJh8s9eHSfNG0l7aM4Fde23ImswkLWCFENr
dKUAGvVf3b4Cy6kKvFZ9/cqrKkDJ2jKZX2hfGdx7Q7vnP0w0Uos3Gd4na/arfX8440n8cH6f7PSdLt9B
WM3q7t1CmbWtl/FXBGKknd0kUv7Yu62b2eYHp920JDKczYyIl7zVkhHCdxHYbcFm1IJEZAasgj1Cq1vh
XigvCl+I8o0b+/ck5ZW+30JQBtp2I5f/BgAA///7bwKfghAAAA==
`,
	},

	"/index.html": {
		name:    "index.html",
		local:   "views/index.html",
		size:    817,
		modtime: 1559859371,
		compressed: `
H4sIAAAAAAAC/4SS3WrDMAyF7/sUwrDLxNBr12MwCoMxxvYEaqwmAtcpjlLojN99JO3omv7dBX0nOkeS
U3K05kCgNshB5TwDADCOd1B57LqFqtogFETZkYy0mU9g0RA64K6oKAhFZZeE0keCpcca3rkTo5u5nZ1a
CK48/XXZ9pGKQ+X0WTRt5J82CPp/3oefB7tpLZ4XjkL7gRsyWprr9JPiEBnrO5qXSrgN3aXA6KnpoBmj
TbKtWrc/V6YEEUNNUMJx5w+nOQBnU4JymApyNlrcfeFX633bS/ktEYXqfXmaGHJOCbaRg4B6Ug/bXQUj
7P1tOAo8W4PQRFovlCbH8sxuMcR7e4WclR1KRqM12vMdH33L6Hrwy/ukBBTcdONGTw5k9PgCj2c02vHO
zlKi4HL+DQAA//8AupN5MQMAAA==
`,
	},

	"/layout.html": {
		name:    "layout.html",
		local:   "views/layout.html",
		size:    1870,
		modtime: 1559804403,
		compressed: `
H4sIAAAAAAAC/7yVzW7bOBDH73mKMXNdWru3HCwBi6BGC7SHNscgB0Ycy5NQpMAZ2TAMv3tBSXbkLyBo
gp48nI///Dgize3W4oI8gmpMhWq3uwEAmE1sKGXTICyldkXvSyY446tcoVd7Jxpb3HR2txYSh8UcjbQR
Ye5MBQ8YV1TiLOtjo+QaxUC5NJFRctXKQt8Num9hb2rMlUUuIzVCwSsogxf0kqv74CUGB5vQRtj3/IUO
DSPDmmQJC2cqnsB8/jC5Ir0iXDchykh3TVaWucWErbvFP0CehIzTXBqH+X/Tf9V4J478K0R0uWLZOOQl
oihYRlzkKjPMKJw1bcSS+19dk5+WzGOm92ik+kXwos0aOdQ4aLyJTLR+pAU4Qfj2Be6eij9jrCJZ1hG5
CZ5phTo4q+ki9uQRvaXFk9bFGUYlA0VyfBbKJQatjzg+bR4JwJlNaIUzJou6Rt8Ow/hbg7jI8K4JzLL+
enb2c7Cb4mYIWFoB2Vz1uqcnCH6gb0FCVTmEI14zsN0mBtVJJOs7+VcFpTPMvUOnjY34DspfTf3cxgoj
UBn8kXiXw43xxSzrft7aZmZMuIfvGE7qU2zg6O/Zec6wkdMsnWZFvtqP/1YV8/nDcetDeevO6x2xXGg1
fOfzfBKsVXEYaaYuSaYppn9RTiCzzNEHGpQRjeDVNvdd+ON9bq9v5P+f1+VnWetOPmZmaTU+Bd3ywjkw
5E8Gv90K1o0zgtCHYTo8badKR3b/yADHMlcvnLU0fWGVjmPnP1T0V2m4Yd0Dud2it7vd7wAAAP//AcHd
0k4HAAA=
`,
	},

	"/login.html": {
		name:    "login.html",
		local:   "views/login.html",
		size:    299,
		modtime: 1559745144,
		compressed: `
H4sIAAAAAAAC/1SPQYrDMAxF9zmF0L7JBeJcYuYCbqy2YmzJ2ApDMb77kLYwyU58/SeeWgt0YyHA5Fmw
9wEAYL5pSbBGX6vDvBW67AGCX41VHE5R7ywIieyhwWHWari80DfOFEMl+49eMUveDMQncmj6Q4Jgz0wO
s6/1V0tAyNGv9NAYqDj83jvjOOIynA9dNzOVD1y3a2LDk+6ncJgvuXDy5YnLF98FWObpvThoT2fvedq/
XobWSELvfwEAAP//jOQDGCsBAAA=
`,
	},

	"/views": {
		name:  "views",
		local: `./views`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"./views": {
		_escData["/create.html"],
		_escData["/edit.html"],
		_escData["/index.html"],
		_escData["/layout.html"],
		_escData["/login.html"],
	},
}