package stdlib

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var stdlibPath = getPath()

// Supports checks if we support this stdlib import
func Supports(importPath string) (alias string, err error) {
	if stdlibPath == "" {
		return "", errors.New("stdlib path not found")
	}

	if val, ok := stdlib[importPath]; ok {
		if !val {
			return "", fmt.Errorf("the '%s' package is not supported in joy yet", importPath)
		}

		return path.Join(stdlibPath, importPath), nil
	}

	return "", nil
}

// get the new path of the standard library
func getPath() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}

	gosrc := path.Join(os.Getenv("GOPATH"), "src")
	rel, e := filepath.Rel(gosrc, path.Dir(file))
	if e != nil {
		return ""
	}

	return rel
}

// Map containing the standard library
var stdlib = map[string]bool{
	"archive/tar":          false,
	"archive/zip":          false,
	"bufio":                false,
	"builtin":              false,
	"bytes":                false,
	"compress/bzip2":       false,
	"compress/flate":       false,
	"compress/gzip":        false,
	"compress/lzw":         false,
	"compress/zlib":        false,
	"container/heap":       false,
	"container/list":       false,
	"container/ring":       false,
	"context":              false,
	"crypto":               false,
	"crypto/aes":           false,
	"crypto/cipher":        false,
	"crypto/des":           false,
	"crypto/dsa":           false,
	"crypto/ecdsa":         false,
	"crypto/elliptic":      false,
	"crypto/hmac":          false,
	"crypto/md5":           false,
	"crypto/rand":          false,
	"crypto/rc4":           false,
	"crypto/rsa":           false,
	"crypto/sha1":          false,
	"crypto/sha256":        false,
	"crypto/sha512":        false,
	"crypto/subtle":        false,
	"crypto/tls":           false,
	"crypto/x509":          false,
	"crypto/x509/pkix":     false,
	"database/sql":         false,
	"database/sql/driver":  false,
	"debug/dwarf":          false,
	"debug/elf":            false,
	"debug/gosym":          false,
	"debug/macho":          false,
	"debug/pe":             false,
	"debug/plan9obj":       false,
	"encoding":             false,
	"encoding/ascii85":     false,
	"encoding/asn1":        false,
	"encoding/base32":      false,
	"encoding/base64":      false,
	"encoding/binary":      false,
	"encoding/csv":         false,
	"encoding/gob":         false,
	"encoding/hex":         false,
	"encoding/json":        true,
	"encoding/pem":         false,
	"encoding/xml":         false,
	"errors":               true,
	"expvar":               false,
	"flag":                 false,
	"fmt":                  true,
	"go/ast":               false,
	"go/build":             false,
	"go/constant":          false,
	"go/doc":               false,
	"go/format":            false,
	"go/importer":          false,
	"go/parser":            false,
	"go/printer":           false,
	"go/scanner":           false,
	"go/token":             false,
	"go/types":             false,
	"hash":                 false,
	"hash/adler32":         false,
	"hash/crc32":           false,
	"hash/crc64":           false,
	"hash/fnv":             false,
	"html":                 false,
	"html/template":        false,
	"image":                false,
	"image/color":          false,
	"image/color/palette":  false,
	"image/draw":           false,
	"image/gif":            false,
	"image/jpeg":           false,
	"image/png":            false,
	"index/suffixarray":    false,
	"io":                   false,
	"io/ioutil":            false,
	"log":                  false,
	"log/syslog":           false,
	"math":                 false,
	"math/big":             false,
	"math/bits":            false,
	"math/cmplx":           false,
	"math/rand":            false,
	"mime":                 false,
	"mime/multipart":       false,
	"mime/quotedprintable": false,
	"net":                 false,
	"net/http":            false,
	"net/http/cgi":        false,
	"net/http/cookiejar":  false,
	"net/http/fcgi":       false,
	"net/http/httptest":   false,
	"net/http/httptrace":  false,
	"net/http/httputil":   false,
	"net/http/pprof":      false,
	"net/mail":            false,
	"net/rpc":             false,
	"net/rpc/jsonrpc":     false,
	"net/smtp":            false,
	"net/textproto":       false,
	"net/url":             false,
	"os":                  false,
	"os/exec":             false,
	"os/signal":           false,
	"os/user":             false,
	"path":                false,
	"path/filepath":       false,
	"plugin":              false,
	"reflect":             false,
	"regexp":              false,
	"regexp/syntax":       false,
	"runtime":             false,
	"runtime/cgo":         false,
	"runtime/debug":       false,
	"runtime/msan":        false,
	"runtime/pprof":       false,
	"runtime/race":        false,
	"runtime/trace":       false,
	"sort":                false,
	"strconv":             true,
	"strings":             true,
	"sync":                false,
	"sync/atomic":         false,
	"syscall":             false,
	"testing":             false,
	"testing/iotest":      false,
	"testing/quick":       false,
	"text/scanner":        false,
	"text/tabwriter":      false,
	"text/template":       false,
	"text/template/parse": false,
	"time":                true,
	"unicode":             false,
	"unicode/utf16":       false,
	"unicode/utf8":        false,
	"unsafe":              false,
}
