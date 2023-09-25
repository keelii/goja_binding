package goja_binding

import (
	"dario.cat/mergo"
	_ "embed"
	"encoding/json"
	"github.com/dop251/goja"
	"log"
)

//go:embed js/prettier.js
var prettierJs string

var polyfill = `
function RangeError() {}
`
var runtimeFunction = `
function format(str, opt) {
	return prettier.format(str, JSON.parse(opt))
}
`

type PrettierFormatOptions struct {
	UseTabs       bool   `json:"useTabs"`
	TabWidth      int    `json:"tabWidth"`
	PrintWidth    int    `json:"printWidth"`
	SingleQuote   bool   `json:"singleQuote"`
	TrailingComma string `json:"trailingComma"`
	Semi          bool   `json:"semi"`
}

var format func(code string, optJsonString string) string

var defaults = PrettierFormatOptions{
	UseTabs:       false,
	TabWidth:      4,
	PrintWidth:    80,
	SingleQuote:   false,
	TrailingComma: "none",
	Semi:          false,
}

func ToJsonString[T any](value T) string {
	data, err := json.Marshal(value)
	if err != nil {
		log.Println("ToJsonString error:", err)
		return ""
	}
	return string(data)
}

func PrettierFormat(code string, opts PrettierFormatOptions) (string, error) {
	if err := mergo.Map(&opts, defaults); err != nil {
		log.Fatalln("mergo defaults error:", err)
		return code, err
	}
	return format(code, ToJsonString(opts)), nil
}

func init() {
	vm := goja.New()
	_, err := vm.RunString(polyfill + prettierJs + runtimeFunction)
	if err != nil {
		log.Fatalln("prettier error:", err)
	}

	err = vm.ExportTo(vm.Get("format"), &format)
	if err != nil {
		panic(err)
	}
}
