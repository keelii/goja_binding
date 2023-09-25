package goja_binding

import "testing"

func TestCompressText(t *testing.T) {
	// compress
	if ret := CompressText("你好"); ret != "5L2g5aW9" {
		t.Error("CompressText error", ret)
	}
	// decompress
	if ret := DecompressText("5L2g5aW9"); ret != "你好" {
		t.Error("DecompressText error", ret)
	}

}
