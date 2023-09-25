package goja_binding

import (
	"github.com/dop251/goja"
	"log"
)

func JSONParse(str string) interface{} {
	return nil
}

var runtimeFunction1 = `
function JsonParse(str) {
	return JSON.parse(str)
}
function JsonStringify(obj, replacer, space) {
	return JSON.stringify(obj, replacer, space)
}
`

var JsonParse func(string) interface{}
var JsonStringify func(interface{}, interface{}, string) string

func init() {
	vm := goja.New()
	_, err := vm.RunString(runtimeFunction1)
	if err != nil {
		log.Fatalln("error:", err)
	}

	err = vm.ExportTo(vm.Get("JsonParse"), &JsonParse)
	if err != nil {
		panic(err)
	}

	err = vm.ExportTo(vm.Get("JsonStringify"), &JsonStringify)
	if err != nil {
		panic(err)
	}
}
