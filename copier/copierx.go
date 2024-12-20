package copier

import (
	"encoding/json"
	"fmt"
)

func Struct2Json(from interface{}) (err error) {
	str, err := json.Marshal(from)
	if err != nil {
		return
		//log.Warn("Struct2Json, can not convert: %v", err)
	}
	from = string(str)
	// return string(str), err
	return
}

func Json2Struct(to interface{}, from string) (err error) {
	err = json.Unmarshal([]byte(from), to)
	if err != nil {
		return
		//log.Warn("Json2Struct, can not convert: %v", err)
	}
	return
}

func Struct2StructByJson(to interface{}, from interface{}) {
	Struct2Json(from)
	Json2Struct(to, from.(string))
}

func Json2Map(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	return m, nil
}

// Copy 从一个结构体复制到另一个结构体
func CopyX(from, to interface{}) error {
	b, err := json.Marshal(from)
	if err != nil {
		return err
		// return errors.Wrap(err, "marshal from data err")
	}

	err = json.Unmarshal(b, to)
	if err != nil {
		return err
		//return errors.Wrap(err, "unmarshal to data err")
	}

	return nil
}
