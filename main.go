package main

import (
	"github.com/MarcosIgnacioo/htmx-go/controllers"
	"github.com/MarcosIgnacioo/htmx-go/initializers"
	"github.com/gin-gonic/gin"
)

func init()  {
    initializers.LoadEnvVariables()
}

func main()  {
    r := gin.Default()

    r.GET("/", )
    r.GET("/notes", controllers.GetNotes)
    r.GET("/notes/:id", controllers.GetNote)

    r.PUT("/notes/:id", controllers.UpdateNote)
    r.DELETE("/notes/:id", controllers.DeleteNote)
    r.POST("/note", controllers.PostNote)

    r.Run()
}
