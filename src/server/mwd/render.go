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

		file, err := ioutil.ReadFile("public/server.js")

		Must(c, err)

		_, err = vm.RunString(string(file))

		Must(c, err)

		module := vm.Get("server")

		renderJS := module.ToObject(vm).Get("render")

		exclude, ok := goja.AssertFunction(renderJS)

		if ok {
			rendered, err := exclude(nil, vm.ToValue(c.Request.RequestURI))

			Must(c, err)

			htmlString := rendered.Export().(string)

			c.Set("HTML", htmlString)

			c.Next()
		}

		c.Set("error", "render methode cannot be executed")

	}
}

func Must(c *gin.Context, err error) {
	if err != nil {
		c.Set("error", err.Error())
		c.Set("HTML", "")
		c.Next()
		return
	}

}
