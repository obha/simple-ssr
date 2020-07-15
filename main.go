package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dop251/goja"
)

func main() {
	vm := goja.New()

	file, err := ioutil.ReadFile("dist/bundle.js")

	if err != nil {
		panic(err)
	}
	v, err := vm.RunString(string(file))

	if err != nil {
		panic(err)
	}

	fmt.Print(v)

}
