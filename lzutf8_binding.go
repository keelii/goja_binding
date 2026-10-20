package goja_binding

import (
	_ "embed"
	"log"
	"sync"

	"github.com/dop251/goja"
)

//go:embed js/lzutf8.min.js
var lzutf8 []byte

var exportFunctions = `
function compressText(str) { return LZUTF8.compress(str, {outputEncoding: "Base64"}) }
function decompressText(str) { return LZUTF8.decompress(str, {inputEncoding: "Base64"}) }
`

var (
	compressText   func(string) string
	decompressText func(string) string
	runtimeMu      sync.Mutex
)

func init() {
	vm := goja.New()
	_, err := vm.RunString(string(lzutf8) + exportFunctions)
	if err != nil {
		log.Fatalln("lzutf8 error:", err)
	}
	err = vm.ExportTo(vm.Get("compressText"), &compressText)
	if err != nil {
		log.Fatalln("compressText error:", err)
	}

	err = vm.ExportTo(vm.Get("decompressText"), &decompressText)
	if err != nil {
		log.Fatalln("decompressText error:", err)
	}
}

func CompressText(code string) string {
	runtimeMu.Lock()
	defer runtimeMu.Unlock()

	return compressText(code)
}
func DecompressText(code string) string {
	runtimeMu.Lock()
	defer runtimeMu.Unlock()

	return decompressText(code)
}
