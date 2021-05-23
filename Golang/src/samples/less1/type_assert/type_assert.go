package main

import "fmt"

type Module interface {
	Names() string
}

type Defaultable interface {
	defaults(name string)
}

type DefaultableModule interface {
	Module
	Defaultable
}

type DefaultableModuleBase struct {
}

func (d *DefaultableModuleBase) defaults(name string) {
	fmt.Printf("Defaults with name : %s\n", name)
}

type ModuleBase struct {
	DefaultableModuleBase
}

func (b ModuleBase) Names() string {
	return "[wjxing]"
}

var _ Defaultable = (*DefaultableModuleBase)(nil)

func InitDefaultableModule(module DefaultableModule) {
	module.(Defaultable).defaults(module.(Module).Names())
}

type Flags struct {
	BaseFlags
	base BaseFlags
}

func (flags *Flags) getFlags() {
	fmt.Println("Get Flags")
}

type BaseFlags struct {
}

func (flags *BaseFlags) getFlags() {
	fmt.Println("Get Base Flags")
}

func main() {
	b := &ModuleBase{}
	InitDefaultableModule(b)

	f := &Flags{
		BaseFlags: BaseFlags{},
		base:      BaseFlags{},
	}
	f.getFlags()
}
