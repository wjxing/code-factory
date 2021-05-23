package main

import "fmt"

type ConfigInter interface {
	GeneratorBuild() bool
}

func isTrueConfig1(config ConfigInter) {
	fmt.Println(config)
}

func isTrueConfig2(config interface{}) {
	c, ok := config.(ConfigInter)
	fmt.Println(c)
	fmt.Println(ok)
}

type ConfigFake struct {
}

type ConfigImpl struct {
	name string
}

func (c ConfigImpl) GeneratorBuild() bool {
	return true
}

type Tebby struct {
}

func (c *Tebby) Meow() {
	fmt.Println("Tebby Meow")
}

func main() {
	config1 := ConfigFake{}
	isTrueConfig2(config1)
	config2 := ConfigImpl{name: "wjxing"}
	isTrueConfig1(config2)

	tebby := Tebby{}
	tebby.Meow()
}
