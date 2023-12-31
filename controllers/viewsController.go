package controllers

import (
	"net/http"

	"github.com/MarcosIgnacioo/htmx-go/initializers"
	"github.com/MarcosIgnacioo/htmx-go/models"
	"github.com/gin-gonic/gin"
)

func ViewNotes(c *gin.Context)  {
    notes := SelectAllNotes()
    c.HTML(http.StatusOK, "index.html", notes)
}

func ViewNote(c *gin.Context)  {
    id := c.Param("id")

    var note *models.Note
    initializers.DB.First(&note, id)
    c.HTML(http.StatusOK, "note-details.html", note)
}
