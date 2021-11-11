package tool

import (
	"bytes"
	"encoding/gob"
)

func Byte2Struct(s []byte) map[string]interface{} {
	p := map[string]interface{}{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil {
		panic(err)
	}
	return p
}
