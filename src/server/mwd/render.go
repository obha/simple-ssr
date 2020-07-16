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

		c.Set("error", Must(err))

		_, err = vm.RunString(string(file))

		c.Set("error", Must(err))

		module := vm.Get("server")

		renderJS := module.ToObject(vm).Get("render")

		exclude, ok := goja.AssertFunction(renderJS)

		if ok {
			html, err := exclude(nil, vm.ToValue(c.Request.RequestURI))

			c.Set("error", Must(err))

			htmlString := html.Export().(string)

			c.Set("HTML", htmlString)

			c.Next()
		}

		c.Set("error", "render methode cannot be executed")

	}
}

func Must(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
