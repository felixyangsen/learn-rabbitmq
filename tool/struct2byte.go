package tool

import (
	"bytes"
	"encoding/gob"
)

func Struct2Byte(p map[string]interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}
