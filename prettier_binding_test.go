package goja_binding

import (
	"testing"
)

func TestPrettierFormat(t *testing.T) {
	// js
	if ret, _ := PrettierFormat("var a=1", PrettierFormatOptions{}); ret != "var a = 1\n" {
		t.Error("PrettierFormat error")
	}
	// ts
	if ret, _ := PrettierFormat("var a: number = 1", PrettierFormatOptions{}); ret != "var a: number = 1\n" {
		t.Error("PrettierFormat error", ret)
	}
	// jsx
	if ret, _ := PrettierFormat("var a=<b>1</b>", PrettierFormatOptions{}); ret != "var a = <b>1</b>\n" {
		t.Error("PrettierFormat error", ret)
	}
}
