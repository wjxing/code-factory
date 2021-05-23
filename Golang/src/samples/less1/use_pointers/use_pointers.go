package main

import (
	"fmt"
	"reflect"
)

type String struct {
	Value string
}

type ModuleBase struct {
	props []interface{}
}

func unpack1(props ...interface{}) {
	for _, v := range props {
		field := v.(*String)
		field.Value = "[wjxing]"
	}
}

func unpack2(props ...interface{}) {
	for _, v := range props {
		fieldType := reflect.TypeOf(v)
		field := reflect.ValueOf(v).Elem().Field(0)
		fmt.Println("Type :", fieldType, ", value :", field)
		field.SetString("[Wjxing]")
	}
}

func main() {
	module := &ModuleBase{}
	str1 := &String{"[]"}
	module.props = append(module.props, str1)
	unpack1(module.props...)
	for _, v := range module.props {
		fmt.Println("In stage 1 : Value", v)
	}
	unpack2(module.props...)
	for _, v := range module.props {
		fmt.Println("In stage 2 : Value", v)
	}
}
