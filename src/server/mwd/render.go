package mdw

import (
	"io/ioutil"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
)

//RenderMiddleware - render compiled js via vm
func RenderMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		vm := goja.New()

		file, err := ioutil.ReadFile("dist/server.js")

		Must(err)

		_, err = vm.RunString(string(file))

		Must(err)

		module := vm.Get("server")

		renderJS := module.ToObject(vm).Get("render")

		exclude, _ := goja.AssertFunction(renderJS)

		html, err := exclude(nil, vm.ToValue(c.Request.RequestURI))

		Must(err)

		htmlString := html.Export().(string)

		c.Set("HTML", htmlString)
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
