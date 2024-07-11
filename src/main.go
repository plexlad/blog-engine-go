// Binary for web server
package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Template rendering (doesn't really need its own library)
type Template struct {
  templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
  e := echo.New()  // New Echo server
  // Templating engine
  t := &Template{
    templates: template.Must(template.ParseGlob("templates/*")),
  }
  e.Renderer = t

  e.Use(middleware.Logger()) // Logging and recovering
  e.Use(middleware.Recover())

  // Root //
  var empty interface{}
  e.GET("/", func(c echo.Context) error {
    return c.Render(http.StatusOK, "base", empty)
  })

  // Sign-in functionality //
  // Gets the login website
  e.GET("/signin", func(c echo.Context) error {
    return c.String(http.StatusOK, "Sign in")
  })

  // Where the user login info is sent to
  e.POST("/signin", func(c echo.Context) error { 
    return c.String(http.StatusOK, "Sign in post")
  })

  // Private articles //
  // Group is used for permission middleware
  privateGroup := e.Group("/private")
  // TODO create permission middleware
  //privateGroup.Use()
  privateGroup.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Permission granted!")
  })

  e.Logger.Fatal(e.Start(":42069"))
}
