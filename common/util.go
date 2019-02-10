package common

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
)

func Serialize(object interface{}) ([]byte, error) {
	data, err := json.Marshal(object)
	if err != nil {
		panic(fmt.Sprintf("Error encoding : %s", err))
	}
	return data, nil
}

func Deserialize(serializedBytes []byte, object interface{}) error {
	if len(serializedBytes) == 0 {
		return nil
	}
	err := json.Unmarshal(serializedBytes, object)
	if err != nil {
		panic(fmt.Sprintf("Error decoding : %s", err))
	}
	return err
}

func IntToBytes(n interface{}) []byte {
	buff := new(bytes.Buffer)
	switch n.(type) {
	case int64:
		err := binary.Write(buff, binary.LittleEndian, n.(int64))
		if err != nil {
			log.Panic(err)
		}
	case int32:
		err := binary.Write(buff, binary.LittleEndian, n.(int32))
		if err != nil {
			log.Panic(err)
		}
	}

	return buff.Bytes()
}
