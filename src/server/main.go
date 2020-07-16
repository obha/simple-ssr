package server

import (
	"html/template"
	"net/http"
	mdw "simple-final/src/server/mwd"

	"github.com/gin-gonic/gin"
)

type App struct {
	Engin *gin.Engine
}

func NewApp() *App {

	r := gin.Default()

	r.Use(mdw.RenderMiddleware())

	r.Static("/", "public")

	r.LoadHTMLFiles("public/index.html", "public/500.html")

	r.NoRoute(func(c *gin.Context) {
		htmlString := c.MustGet("HTML").(string)
		errorString := c.MustGet("error").(string)

		if errorString != "" {
			c.HTML(500, "500.html", gin.H{
				"__ERROR__": errorString,
			})
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"__TITLE__": "Influlook",
			"__HTML__":  template.HTML(htmlString),
		})
	})

	return &App{Engin: r}

}

func (app *App) Run() {
	app.Engin.Run(":8080")
}
