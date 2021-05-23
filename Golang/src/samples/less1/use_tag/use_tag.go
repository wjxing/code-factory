package main

import (
	"fmt"
	"reflect"
)

type aType struct {
	Names string
}

type AType interface {
	names()
}

func (a *aType) names() {
	fmt.Println("Names = ", a.Names)
}

//var _ AType = aType{}
var _ AType = &aType{}

func main() {
	type User struct {
		UserID int `android:"user_id,USER_ID"`
	}
	u := &User{UserID: 1}
	t := reflect.TypeOf(u)
	field := t.Elem().Field(0)
	fmt.Println("Tag : ", field.Tag.Get("android"))
	fmt.Println("Tag : ", field.Tag)

	a := &aType{Names: "wjxing"}
	a.names()
}
