// Code generated by "esc -o ./migrations/fs.go -pkg migrations ./migrations"; DO NOT EDIT.

package migrations

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

	"/migrations/fs.go": {
		name:    "fs.go",
		local:   "migrations/fs.go",
		size:    0,
		modtime: 1574378160,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/migrations/postgres/1561423897_create_feature_flags_table.down.sql": {
		name:    "1561423897_create_feature_flags_table.down.sql",
		local:   "migrations/postgres/1561423897_create_feature_flags_table.down.sql",
		size:    69,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/3IJ8g9Q8PRzcY1QUMrJz8+OT0tNLCktSo1Py0lML45PqozPS8xNVbLmAisMcXTycVVQ
QlGjZM0FCAAA//9Q/JbBRQAAAA==
`,
	},

	"/migrations/postgres/1561423897_create_feature_flags_table.up.sql": {
		name:    "1561423897_create_feature_flags_table.up.sql",
		local:   "migrations/postgres/1561423897_create_feature_flags_table.up.sql",
		size:    560,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/4SRwW7qMBBF9/mKK1YgPfZPYhWohawG0yaOBCvLIUNiNdg0dqry91VSWhGJCi9nfI7u
zKxSFksGGS8ThsmRdOhaUsdGV34STSMAMCUevoylPE4gthIiTxK8pHwTp3s8s/2/QWL1iR5JJNtJ4Ffy
DbauaVwXVKttqTzRvTBLvuZC3oADOZ/jaD4RauMRHLJNnCRcyJ/eUPe165oSBfVfqURxQa0/jK0QasKZ
2gPZoCuC9tAwNvyHsUPv5EpqRgl9aHWg6qJusD8T3gVLOhhvnFWNq8xB6bMZdnKdZ7UVmUzj3jO6k+p3
q4xXnTXvyAV/zRmmfXEWzRZRdD0xF09sh8a5NzXGi8tgwFaMvcgzLtYoQkt09S2irwAAAP//pg8A3zAC
AAA=
`,
	},

	"/migrations/postgres/1562536633_create_pilots_table.down.sql": {
		name:    "1562536633_create_pilots_table.down.sql",
		local:   "migrations/postgres/1562536633_create_pilots_table.down.sql",
		size:    84,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/3IJ8g9Q8PRzcY1QUMrJz88uLYgvyMzJLymOT6qMT0tNLCktSo1Py0lMj89MiU/MS4lP
rShJLcpLzInPTFGy5gJrD3F08nFVUILoU7LmAgQAAP//3ZtLD1QAAAA=
`,
	},

	"/migrations/postgres/1562536633_create_pilots_table.up.sql": {
		name:    "1562536633_create_pilots_table.up.sql",
		local:   "migrations/postgres/1562536633_create_pilots_table.up.sql",
		size:    381,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/4SQQWvDIBiG7/6Kj55a6D/oyXRSBPlkiYHePsxih8xpZw1s/34kKSPJZV59fN/H91wL
bgQYXikBu7sPqTx2bM8AAHwPq1PJSyNqyRWgNoCtUseJuzlbhuzoFuw7+X7kJJrxZs257+JytIGewUZc
zZy84WJOIbi/9kprJTiuuAk8a2xMzceyyZyG6L/oLX12PtriU4QW5WsrYL9RPC5dDuxwYuw5hMQXcYWQ
0sdwp3kO6n5o855s7Gn5HY2zwQPaRuIFupKd+6f2xH4DAAD//0TGwKl9AQAA
`,
	},

	"/migrations/postgres/1562584482_create_tokens_table.down.sql": {
		name:    "1562584482_create_tokens_table.down.sql",
		local:   "migrations/postgres/1562584482_create_tokens_table.down.sql",
		size:    63,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/3IJ8g9Q8PRzcY1QUMrJz88uLYgvyc9OzSuOT6qEsOJLUitKlKy5wCpDHJ18XBWUIEqU
rLkAAQAA//97cxxbPwAAAA==
`,
	},

	"/migrations/postgres/1562584482_create_tokens_table.up.sql": {
		name:    "1562584482_create_tokens_table.up.sql",
		local:   "migrations/postgres/1562584482_create_tokens_table.up.sql",
		size:    334,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/3yQQUsDMRCF7/kVj55a8KLgqae0DiWwndXdWShewpYNGCob3SSo/166jYI9OJcZmO89
3sy2IS0E0ZuKsEjh5Ma4UEsFAH5AqY3ZtdQYXQHgWsBdVd3MTHzp72/vzpPQQX74v0z4GN1ksx/+YXyM
2Q22TxCzp1b0/lGer5ghT33yYZzzGJZrnxna1txKo8/ryzV2btZHm0f/jo7NU0dYXoKv1GqtVHmC4Qc6
4DWEU36zRXz8KvrkPhNqLqboWsM7HNPk3K/XWn0HAAD//+wxzmJOAQAA
`,
	},

	"/migrations/postgres/1563313726_create_index_lookup_pilots_by_ext_id.down.sql": {
		name:    "1563313726_create_index_lookup_pilots_by_ext_id.down.sql",
		local:   "migrations/postgres/1563313726_create_index_lookup_pilots_by_ext_id.down.sql",
		size:    37,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/3IJ8g9Q8PRzcY1QUMrJz88uLYgvyMzJLymOT6qMT60oic9MUbIGBAAA//8cIMYUJQAA
AA==
`,
	},

	"/migrations/postgres/1563313726_create_index_lookup_pilots_by_ext_id.up.sql": {
		name:    "1563313726_create_index_lookup_pilots_by_ext_id.up.sql",
		local:   "migrations/postgres/1563313726_create_index_lookup_pilots_by_ext_id.up.sql",
		size:    79,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/3IOcnUMcVXw9HNxjVBQysnPzy4tiC/IzMkvKY5PqoxPrSiJz0xRUvD3U1CCiCophAZ7
+rkrJJUUpaYqaCilVpSkFuUl5oCUaVoDAgAA//9FwpxlTwAAAA==
`,
	},

	"/migrations/postgres/1564874467_create_test_entities_table.down.sql": {
		name:    "1564874467_create_test_entities_table.down.sql",
		local:   "migrations/postgres/1564874467_create_test_entities_table.down.sql",
		size:    27,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/3IJ8g9QCHF08nFVUCpJLS6JT80rySzJTC1WsgYEAAD//18Bt9QbAAAA
`,
	},

	"/migrations/postgres/1564874467_create_test_entities_table.up.sql": {
		name:    "1564874467_create_test_entities_table.up.sql",
		local:   "migrations/postgres/1564874467_create_test_entities_table.up.sql",
		size:    59,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/3IOcnUMcVUIcXTycVVQKkktLolPzSvJLMlMLVbi0uBSUFBQyExRcPJ0D3YN8nT0UfDz
D1HwC/Xx4dK0BgQAAP//QyRBQTsAAAA=
`,
	},

	"/migrations/postgres/1568143862_alter_table_release_flags.down.sql": {
		name:    "1568143862_alter_table_release_flags.down.sql",
		local:   "migrations/postgres/1568143862_alter_table_release_flags.down.sql",
		size:    188,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/2TOsQrCMBRG4T1P8XNfI1NqLxJoEglB3UKEWwejhUYH395BKZbuh4/T8d56rVQfwwFH
yyfQKOX5miWPtVwbaaXMkDjC+p7PoDpNtzxLldK+Sb6886PchRQARPbGMVL4lStsSRc0mW5g0L/XNtBm
aBecs0l/AgAA///JK7KBvAAAAA==
`,
	},

	"/migrations/postgres/1568143862_alter_table_release_flags.up.sql": {
		name:    "1568143862_alter_table_release_flags.up.sql",
		local:   "migrations/postgres/1568143862_alter_table_release_flags.up.sql",
		size:    223,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/2zMuwrCMBSH8T1P8Sejr9AprUcJ5ALpQd1ChNTBaKHRwbcXUYq3/ft+La21a4RQhilA
uyXtIMs4HuOQ0+U65TiUdKhxf4vndMpSAEAgpyyB/SudcsmpPtO5nFFWrSHID+/HeSfq4+0CKSZsNG2/
X6gePRnqGAusgrf/dm+t5uYeAAD//6bKr1vfAAAA
`,
	},

	"/migrations/postgres/1570774790_create_release_allows_table.down.sql": {
		name:    "1570774790_create_release_allows_table.down.sql",
		local:   "migrations/postgres/1570774790_create_release_allows_table.down.sql",
		size:    95,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/3IJ8g9Q8PRzcY1QSMvMS4kvSs1JTSxOjU/LSUyPzyyIT0xJKYpPzMnJLy+OT6qECqdY
c4H1hTg6+bgq4NFiDQgAAP//1QXxzl8AAAA=
`,
	},

	"/migrations/postgres/1570774790_create_release_allows_table.up.sql": {
		name:    "1570774790_create_release_allows_table.up.sql",
		local:   "migrations/postgres/1570774790_create_release_allows_table.up.sql",
		size:    255,
		modtime: 1572749977,
		compressed: `
H4sIAAAAAAAC/3yOQQqDMBBF9znFXyr0Bq5iGyQ0HUuMUFchkliE0BYtlN6+qOnWWf558+cdteBGwPBS
CUwhBjcHO0R3t+PLOu8n62J8fmaWMQAYPdYpZdUILbkC1QbUKoWrlheuO5xFd1jZrcUvrCSzJH9226cH
kCQMywvGkoukk7hhGB/e7gjZ/ptiv7bVtKePtpFUoX9PISBLd3nBfgEAAP//y7Dajv8AAAA=
`,
	},

	"/migrations": {
		name:  "migrations",
		local: `./migrations`,
		isDir: true,
	},

	"/migrations/postgres": {
		name:  "postgres",
		local: `migrations/postgres`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"./migrations": {
		_escData["/migrations/fs.go"],
		_escData["/migrations/postgres"],
	},

	"migrations/postgres": {
		_escData["/migrations/postgres/1561423897_create_feature_flags_table.down.sql"],
		_escData["/migrations/postgres/1561423897_create_feature_flags_table.up.sql"],
		_escData["/migrations/postgres/1562536633_create_pilots_table.down.sql"],
		_escData["/migrations/postgres/1562536633_create_pilots_table.up.sql"],
		_escData["/migrations/postgres/1562584482_create_tokens_table.down.sql"],
		_escData["/migrations/postgres/1562584482_create_tokens_table.up.sql"],
		_escData["/migrations/postgres/1563313726_create_index_lookup_pilots_by_ext_id.down.sql"],
		_escData["/migrations/postgres/1563313726_create_index_lookup_pilots_by_ext_id.up.sql"],
		_escData["/migrations/postgres/1564874467_create_test_entities_table.down.sql"],
		_escData["/migrations/postgres/1564874467_create_test_entities_table.up.sql"],
		_escData["/migrations/postgres/1568143862_alter_table_release_flags.down.sql"],
		_escData["/migrations/postgres/1568143862_alter_table_release_flags.up.sql"],
		_escData["/migrations/postgres/1570774790_create_release_allows_table.down.sql"],
		_escData["/migrations/postgres/1570774790_create_release_allows_table.up.sql"],
	},
}
