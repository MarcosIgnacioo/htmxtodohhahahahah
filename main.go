package main

import (
	"github.com/MarcosIgnacioo/htmx-go/controllers"
	"github.com/MarcosIgnacioo/htmx-go/initializers"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func init()  {
    initializers.LoadEnvVariables()
}

func BlockTemplates() multitemplate.Renderer {
  r := multitemplate.NewRenderer()
  r.AddFromFiles("index","./public/views/index.html", "./public/views/test.tmpl")
  return r
}
func main()  {
    r := gin.Default()
    r.LoadHTMLGlob("public/views/*")

    type test struct {
        Title string
    }

    r.GET("/notes/:id", controllers.ViewNote)
    r.GET("/", controllers.ViewNotes)
    // r.GET("/", controllers.ViewNotes)
    r.GET("htmx/notes", controllers.GetNotes)
    r.GET("htmx/notes/:id", controllers.GetNote)

    r.PUT("htmx/notes/:id", controllers.UpdateNote)
    r.DELETE("htmx/notes/:id", controllers.DeleteNote)
    r.POST("htmx/note", controllers.PostNote)

    r.Run()
}
