package goja_binding

import (
	"testing"
)

func TestJSONParse(t *testing.T) {
	// parse
	if val, exists := JsonParse(`{"a":1}`).(map[string]interface{})["a"]; exists && val.(int64) != 1 {
		t.Error("JsonParse error", val)
	}
	// stringify
	if ret := JsonStringify(map[string]interface{}{"a": 1}, nil, ""); ret != `{"a":1}` {
		t.Error("JsonStringify error", ret)
	}
	if ret := JsonStringify(map[string]interface{}{"a": 1}, nil, " "); ret != "{\n \"a\": 1\n}" {
		t.Error("JsonStringify error", ret)
	}
}
