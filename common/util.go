package common

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os/user"
	"path"
	"path/filepath"
	"strings"
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

// absolute path로 변경하기
// TODO: rename func
// TODO: Refactoring
func RelativeToAbsolutePath(rpath string) (string, error) {
	if rpath == "" {
		return rpath, nil
	}

	absolutePath := ""

	// 1. ./ ../ 경우
	if strings.Contains(rpath, "./") {
		abs, err := filepath.Abs(rpath)
		if err != nil {
			return rpath, err
		}
		return abs, nil
	}

	// 2. ~/ 홈폴더 경우
	if strings.Contains(rpath, "~") {
		i := strings.Index(rpath, "~") // 처음 나온 ~만 반환

		if i > -1 {
			pathRemain := rpath[i+1:]
			// user home 얻기
			usr, err := user.Current()
			if err != nil {
				return rpath, err
			}
			return path.Join(usr.HomeDir, pathRemain), nil

		} else {
			return rpath, nil
		}
	}

	if string(rpath[0]) == "/" {
		return rpath, nil
	}

	if string(rpath[0]) != "." && string(rpath[0]) != "/" {
		currentPath, err := filepath.Abs(".")
		if err != nil {
			return rpath, err
		}

		return path.Join(currentPath, rpath), nil
	}

	return absolutePath, nil

}
