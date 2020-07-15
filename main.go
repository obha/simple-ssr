package main

import (
	"errors"
	"html/template"
	"io/ioutil"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

type App struct {
	Engin *gin.Engine
}

func NewApp() (*App, error) {

	vm := goja.New()

	file, err := ioutil.ReadFile("dist/bundle.js")

	Must(err)

	_, err = vm.RunString(string(file))

	Must(err)

	module := vm.Get("server")
	renderJS := module.ToObject(vm).Get("render")

	exclude, ok := goja.AssertFunction(renderJS)

	if ok {
		html, err := exclude(nil)

		Must(err)

		r := gin.Default()

		r.Static("/public", "./public")

		var templates = template.Must(template.ParseFiles("public/index.html"))

		r.GET("/", func(c *gin.Context) {
			templates.ExecuteTemplate(
				c.Writer,
				"index.html",
				map[string]interface{}{
					"HTML": template.HTML(html.Export().(string)),
				})
		})

		return &App{Engin: r}, nil
	}
	return nil, errors.New("cannot create app instance")

}

func (app *App) Run() {
	app.Engin.Run(":8080")
}

func main() {
	app, err := NewApp()
	Must(err)
	app.Run()
}
