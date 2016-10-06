package goutil

import (
	"unsafe"
)

//将byte转换为字符串
func byteToString(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}

