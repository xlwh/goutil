package ioutil

import (
	"testing"
)

func TestByteToString(t *testing.T) {
	s := "No zuo no die, why still try?"
    b := []byte(s)
    ret := byteToString(b)
    t.Log(s)
    t.Log(b)
    t.Log("byte to String result:")
    t.Log(ret)
    t.Log([]byte(ret))
}
