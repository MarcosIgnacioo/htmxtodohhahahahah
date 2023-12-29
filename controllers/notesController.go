package controllers

import (
	"github.com/MarcosIgnacioo/htmx-go/initializers"
	"github.com/MarcosIgnacioo/htmx-go/models"
	"github.com/gin-gonic/gin"
)

func init()  {
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func PostNote (c *gin.Context) {

    var body struct {
        Title string `json:"title"` 
        Content string `json:"content"`
        IsDone bool `json:"isDone"`
    }

    c.Bind(&body)
    
    note := models.Note{ Title: body.Title, Content: body.Content, IsDone: body.IsDone }

    result := initializers.DB.Create(&note) // Pasamos el pointer porque queremos que modifique o en si pueda hacer lo que sea con ese y gaste menos memoria
    // value interface{} Es el any pero chido gamer porque namas accepta structs d seguro

    if result.Error != nil {
        c.Status(400)
        return
    }

    c.JSON(200, gin.H{
        "message": note,
    })

}

func GetNotes(c *gin.Context) {

    var notes [] models.Note
    initializers.DB.Find(&notes)

    c.JSON(200, gin.H{
        "message": notes,
    })

}

func GetNote(c *gin.Context) {

    var note  models.Note
    id := c.Param("id")
    initializers.DB.First(&note, id)

    c.JSON(200, gin.H{
        "message": note,
    })

}

func UpdateNote(c *gin.Context)  {

    var body struct {
        Title string `json:"title"` 
        Content string `json:"content"`
        IsDone bool `json:"isDone"`
    }

    c.Bind(&body)


    var note  models.Note
    id := c.Param("id")
    initializers.DB.First(&note, id)
    initializers.DB.Model(&note).Updates(models.Note {
        Title: body.Title,
        Content: body.Content,
        IsDone: body.IsDone,
    })

    initializers.DB.Save(note)

    c.JSON(200, gin.H{
        "message": note,
    })

}

func DeleteNote(c *gin.Context)  {
    
    id := c.Param("id")

    initializers.DB.Delete(&models.Note{}, id)

    c.JSON(200, gin.H{
        "message": 200,
    })

}
